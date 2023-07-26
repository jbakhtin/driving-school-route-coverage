package services

import (
	"context"
)

type UserLoginRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserLoginResponse struct {
	Token string `json:"token,omitempty"`
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

type AuthService interface {
	RegisterUser(ctx context.Context, request UserRegistrationRequest) (*UserRegistrationResponse, error)
	LoginUser(ctx context.Context, request UserLoginRequest) (*UserLoginResponse, error)
}
