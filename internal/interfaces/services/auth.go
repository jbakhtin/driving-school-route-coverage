package services

import (
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
)

type AuthService interface {
	RegisterUser(request services.UserRegistrationRequest) (*models.User, error)
	LoginUser(request services.UserLoginRequest) (*services.UserLoginResponse, error)
}
