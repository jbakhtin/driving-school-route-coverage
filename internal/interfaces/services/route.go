package services

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
)

type RouteService interface {
	CreateRoute(ctx context.Context, routeCreationDTO services.RouteCreationDTO) (*models.Route, error)
	GetRouteByID(ctx context.Context, routeID string) (*models.Route, error)
}
