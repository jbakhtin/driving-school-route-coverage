package repositories

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
)

type CreateRoute struct {
	Name string
	LineString []byte `json:"line,omitempty" type:"geometry(LineString,4326)"`
}

type UpdateRoute struct {
	Name string `json:"Name,omitempty"`
	LineString []byte `json:"line,omitempty" type:"geometry(LineString,4326)"`
}

type RouteRepository interface {
	CreateRoute(ctx context.Context, createRoute CreateRoute) (*models.Route, error)
	GetRouteByID(ctx context.Context, routeID string) (*models.Route, error)
	UpdateRouteByID(ctx context.Context, routeID string, updateRoute UpdateRoute) (*models.Route, error)
}
