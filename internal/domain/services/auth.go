package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	"go.uber.org/zap"
)

// TODO: пусть сервисы имеют общую структуру, а хендлеры нет, так как к хендлеров нет специфических компонентов или интефейсов

type UserLoginRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserLoginResponse struct {
	Token string
}

func (e *UserLoginResponse) Marshal() []byte {
	marshal, err := json.Marshal(e)
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
	Message      string
	RegisteredAt string
}

type AuthService struct {
	logger *zap.Logger
	repo   repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) (*AuthService, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &AuthService{
		logger: logger,
		repo:   repo,
	}, nil
}

func (us *AuthService) RegisterUser(request UserRegistrationRequest) (*models.User, error) {
	user := repositories.UserRegistration{
		Name:     request.Name,
		Lastname: request.Lastname,
		Email:    request.Email,
		Login:    request.Login,
		Password: request.Password,
	}

	h := hmac.New(sha256.New, []byte("test_app_key"))
	h.Write([]byte(fmt.Sprintf("%s:%s", user.Login, user.Password)))
	dst := h.Sum(nil)

	user.Password = fmt.Sprintf("%x", dst)

	userCreated, err := us.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
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

	h := hmac.New(sha256.New, []byte("test_app_key"))
	h.Write([]byte(fmt.Sprintf("%s:%s", userLogin.Login, userLogin.Password)))
	hashedPassword := h.Sum(nil)

	if user.Password != fmt.Sprintf("%x", hashedPassword) {
		return nil, errors.New("password not valid")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// Добавляем информацию о пользователе в токен
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	// Подписываем токен с помощью секретного ключа
	tokenString, err := token.SignedString([]byte("test"))

	var response UserLoginResponse

	response.Token = tokenString

	return &response, nil
}
