package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	"go.uber.org/zap"
)

type UserLoginRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserLoginResponse struct {
	Token string `json:"token,omitempty"`
}

func (ulr *UserLoginResponse) Marshal() []byte {
	marshal, err := json.Marshal(ulr)
	if err != nil {
		return nil
	}

	return marshal
}

type UserRegistrationRequest struct {
	Name                 string `json:"name"`
	Lastname             string `json:"lastname"`
	Email                string `json:"email"`
	Login                string `json:"login"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type UserRegistrationResponse struct {
	Message string `json:"message,omitempty"`
}

func (e *UserRegistrationResponse) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

type AuthService struct {
	logger *zap.Logger
	config *config.Config
	repo   repositories.UserRepository
}

func NewAuthService(cfg config.Config, repo repositories.UserRepository) (*AuthService, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &AuthService{
		logger: logger,
		config: &cfg,
		repo:   repo,
	}, nil
}

func (us *AuthService) RegisterUser(request UserRegistrationRequest) (*UserRegistrationResponse, error) {
	user := repositories.UserRegistration{
		Name:     request.Name,
		Lastname: request.Lastname,
		Email:    request.Email,
		Login:    request.Login,
		Password: request.Password,
	}

	h := hmac.New(sha256.New, []byte(us.config.AppKey))
	h.Write([]byte(fmt.Sprintf("%s:%s", user.Login, user.Password)))
	dst := h.Sum(nil)

	user.Password = fmt.Sprintf("%x", dst)

	_, err := us.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	response := UserRegistrationResponse{
		Message: "User created",
	}

	return &response, nil
}

func (us *AuthService) LoginUser(request UserLoginRequest) (*UserLoginResponse, error) {
	userLogin := repositories.UserRegistration{
		Login:    request.Login,
		Password: request.Password,
	}

	// TODO: find user
	user, err := us.repo.GetUserByLogin(userLogin.Login)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, []byte(us.config.AppKey))
	h.Write([]byte(fmt.Sprintf("%s:%s", userLogin.Login, userLogin.Password)))
	hashedPassword := h.Sum(nil)

	if user.Password != fmt.Sprintf("%x", hashedPassword) {
		return nil, errors.New("password not valid")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	tokenString, err := token.SignedString([]byte(us.config.AppKey))
	if err != nil {
		return nil, err
	}

	var response UserLoginResponse

	response.Token = tokenString

	return &response, nil
}
