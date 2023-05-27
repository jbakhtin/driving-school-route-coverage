package repository

import (
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/query"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(client *postgres.Postgres) (*UserRepository, error) {
	return &UserRepository{
		client,
	}, nil
}

func (ur *UserRepository) CreateUser(user repositories.UserRegistration) (*models.User, error) {
	var stored models.User
	err := ur.QueryRow(query.CreateUser, &user.Name, &user.Lastname, &user.Login, &user.Email, &user.Password).
		Scan(&stored.ID,
			&stored.Name,
			&stored.Lastname,
			&stored.Login,
			&stored.Email,
			&stored.Password,
			&stored.CreatedAt,
			&stored.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &stored, nil
}

func (ur *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User
	err := ur.QueryRow(query.GetUserByID, id).
		Scan(&user.ID,
			&user.Name,
			&user.Lastname,
			&user.Login,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	return nil, nil
}

func (ur *UserRepository) UpdateUser() (bool, error) {
	return false, nil
}

func (ur *UserRepository) DeleteUser() (bool, error) {
	return false, nil
}
