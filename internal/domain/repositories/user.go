package repositories

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
)

type UserRegistration struct {
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Login    string `json:"login,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user UserRegistration) (*models.User, error)
	UpdateUser(ctx context.Context) (bool, error)
	DeleteUser(ctx context.Context) (bool, error)
}
