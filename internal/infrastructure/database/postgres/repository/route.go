package repository

import (
	"context"
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

func (ur *RouteRepository) CreateRoute(ctx context.Context, createRoute repositories.CreateRoute) (*models.Route, error) {
	var stored models.Route

	err := ur.QueryRowContext(ctx, query.CreateRoute, &createRoute.LineString).
		Scan(
			&stored.ID,
			&stored.LineString,
			&stored.CreatedAt,
			&stored.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &stored, nil
}

func (ur *RouteRepository) GetRouteByID(ctx context.Context, routeID string) (*models.Route, error) {
	var route models.Route
	err := ur.QueryRowContext(ctx, query.GetRouteByID, routeID).
		Scan(&route.ID,
			&route.LineString,
			&route.CreatedAt,
			&route.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (ur *RouteRepository) UpdateRouteByID(ctx context.Context, routeID string, updateRoute repositories.UpdateRoute) (*models.Route, error) {
	var route models.Route
	err := ur.QueryRowContext(ctx, query.UpdateRouteByID, &routeID, &updateRoute.LineString).
		Scan(&route.ID,
			&route.LineString,
			&route.CreatedAt,
			&route.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &route, nil
}
