package services

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"go.uber.org/zap"
)

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

func (us *AuthService) RegisterUser(ctx context.Context, request ifaceservice.UserRegistrationRequest) (*ifaceservice.UserRegistrationResponse, error) {
	userRegistration := repositories.UserRegistration{
		Name:     request.Name,
		Lastname: request.Lastname,
		Email:    request.Email,
		Login:    request.Login,
		Password: request.Password,
	}

	user, err := us.repo.GetUserByLogin(ctx, userRegistration.Login)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if user != nil  {
		return nil, apperror.UserAlreadyExists
	}

	h := hmac.New(sha256.New, []byte(us.config.AppKey))
	h.Write([]byte(fmt.Sprintf("%s:%s", userRegistration.Login, userRegistration.Password)))
	dst := h.Sum(nil)

	userRegistration.Password = fmt.Sprintf("%x", dst)

	_, err = us.repo.CreateUser(ctx, userRegistration)
	if err != nil {
		return nil, err
	}

	response := ifaceservice.UserRegistrationResponse{
		Message: "User created",
	}

	return &response, nil
}

func (us *AuthService) LoginUser(ctx context.Context, request ifaceservice.UserLoginRequest) (*ifaceservice.UserLoginResponse, error) {
	userLogin := repositories.UserRegistration{
		Login:    request.Login,
		Password: request.Password,
	}

	user, err := us.repo.GetUserByLogin(ctx, userLogin.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.UserNotFound
		}

		return nil, err
	}

	h := hmac.New(sha256.New, []byte(us.config.AppKey))
	h.Write([]byte(fmt.Sprintf("%s:%s", userLogin.Login, userLogin.Password)))
	hashedPassword := h.Sum(nil)

	if user.Password != fmt.Sprintf("%x", hashedPassword) {
		return nil, apperror.New(nil, "Invalid password", apperror.BadRequestParamsCode, "", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	tokenString, err := token.SignedString([]byte(us.config.AppKey))
	if err != nil {
		return nil, err
	}

	var response ifaceservice.UserLoginResponse

	response.Token = tokenString

	return &response, nil
}
