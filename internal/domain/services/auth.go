package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	userRepo   repositories.UserRepository
	sessionRepo   repositories.Session
}

func NewAuthService(userRepo repositories.UserRepository, sessionRepo repositories.Session) (*AuthService, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &AuthService{
		logger: logger,
		userRepo:   userRepo,
		sessionRepo:   sessionRepo,
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	userCreated, err := us.userRepo.CreateUser(user)
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

	var response UserLoginResponse

	if userLogin.Login != "login" || userLogin.Password != "password" {
		return nil, errors.New("login or password not valid")
	}

	response.Token = "token"

	var seesion models.Session
	seesion.UserId = 1
	seesion.SessionId = response.Token
	seesion.TTL = time.Hour * 24

	err := us.sessionRepo.SetSession(context.TODO(), seesion)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
