package repositories

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
)

type RouteCreation struct {
	LineString []byte `json:"line,omitempty;type:geometry(LineString,4326)"`
}

type RouteRepository interface {
	CreateRoute(ctx context.Context, routeCreation RouteCreation) (*models.Route, error)
	GetRouteByID(ctx context.Context, routeID string) (*models.Route, error)
}
