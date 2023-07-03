package services

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
)

type AuthService interface {
	RegisterUser(ctx context.Context, request services.UserRegistrationRequest) (*services.UserRegistrationResponse, error)
	LoginUser(ctx context.Context, request services.UserLoginRequest) (*services.UserLoginResponse, error)
}
