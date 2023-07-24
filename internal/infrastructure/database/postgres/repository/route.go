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

	err := ur.QueryRowContext(ctx, query.CreateRoute, &createRoute.UserID, &createRoute.Name, &createRoute.LineString).
		Scan(
			&stored.ID,
			&stored.UserID,
			&stored.Name,
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

	userID := ctx.Value("user_id")
	err := ur.QueryRowContext(ctx, query.GetRouteByID, routeID, userID).
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

	userID := ctx.Value("user_id")
	err := ur.QueryRowContext(ctx, query.UpdateRouteByID, &routeID, userID, updateRoute.Name, &updateRoute.LineString).
		Scan(&route.ID,
			&route.UserID,
			&route.Name,
			&route.LineString,
			&route.CreatedAt,
			&route.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (ur *RouteRepository) DeleteRouteByID(ctx context.Context, routeID string) error {
	userID := ctx.Value("user_id")
	err := ur.QueryRowContext(ctx, query.DeleteRouteByID, &routeID, userID).Err()
	if err != nil {
		return err
	}

	return nil
}
