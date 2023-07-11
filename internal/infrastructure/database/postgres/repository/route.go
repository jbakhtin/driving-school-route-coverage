package repository

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/query"
)

type RouteRepository struct {
	*postgres.Postgres
}

func NewRouteRepository(client *postgres.Postgres) (*RouteRepository, error) {
	return &RouteRepository{
		client,
	}, nil
}

func (ur *RouteRepository) CreateRoute(ctx context.Context, routeCreation repositories.RouteCreation) (*models.Route, error) {
	var stored models.Route

	fmt.Println(routeCreation)
	err := ur.QueryRowContext(ctx, query.CreateRoute, &routeCreation.LineString).
		Scan(&stored.Id, &stored.LineString, &stored.CreatedAt, &stored.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &stored, nil
}

func (ur *RouteRepository) GetRouteByID(ctx context.Context, id int) (*models.User, error) {
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