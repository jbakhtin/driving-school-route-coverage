package repositories

import "github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"

type UserRegistration struct {
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Login    string `json:"login,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(user UserRegistration) (*models.User, error)
	UpdateUser() (bool, error)
	DeleteUser() (bool, error)
}
