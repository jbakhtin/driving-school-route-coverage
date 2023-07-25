package repository

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/types"
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

	err := ur.QueryRowContext(ctx, query.CreateRoute, &createRoute.UserID, &createRoute.Name, &createRoute.Description, &createRoute.LineString).
		Scan(
			&stored.ID,
			&stored.UserID,
			&stored.Name,
			&stored.Description,
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

	userID := ctx.Value(types.ContextKeyUserID)
	err := ur.QueryRowContext(ctx, query.GetRouteByID, routeID, userID).
		Scan(&route.ID,
			&route.UserID,
			&route.Name,
			&route.Description,
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

	userID := ctx.Value(types.ContextKeyUserID)
	err := ur.QueryRowContext(ctx, query.UpdateRouteByID, &routeID, userID, updateRoute.Name, updateRoute.Description, &updateRoute.LineString).
		Scan(&route.ID,
			&route.UserID,
			&route.Name,
			&route.Description,
			&route.LineString,
			&route.CreatedAt,
			&route.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (ur *RouteRepository) DeleteRouteByID(ctx context.Context, routeID string) error {
	userID := ctx.Value(types.ContextKeyUserID)
	err := ur.QueryRowContext(ctx, query.DeleteRouteByID, &routeID, userID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (ur *RouteRepository) GetRoutes(ctx context.Context) (*[]models.Route, error) {
	var routes []models.Route

	Limit := 10

	userID := ctx.Value(types.ContextKeyUserID)
	rows, err := ur.QueryContext(ctx, query.GetRoutes, userID, Limit)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var route models.Route

		err := rows.Scan(&route.ID,
			&route.UserID,
			&route.Name,
			&route.Description,
			&route.CreatedAt,
			&route.UpdatedAt)

		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	return &routes, nil
}
