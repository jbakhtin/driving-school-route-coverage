package repository

import (
	"context"

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

func (ur *UserRepository) CreateUser(ctx context.Context, user repositories.UserRegistration) (*models.User, error) {
	var stored models.User
	err := ur.QueryRowContext(ctx, query.CreateUser, &user.Name, &user.Lastname, &user.Login, &user.Email, &user.Password).
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

func (ur *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := ur.QueryRowContext(ctx, query.GetUserByID, id).
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

func (ur *UserRepository) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	err := ur.QueryRowContext(ctx, query.GetUserByLogin, login).
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

func (ur *UserRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	return nil, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context) (bool, error) {
	return false, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context) (bool, error) {
	return false, nil
}
