package services

import (
	"context"

	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/twpayne/go-geom/encoding/geojson"
)

type RouteCreationDTO struct {
	Name        string            `json:"name" validate:"required"`
	Description *string           `json:"description" validate:"required"`
	Line        *geojson.Geometry `json:"geometry" validate:"required,linestring"`
}

type UpdateRoute struct {
	Name        string            `json:"name" validate:"required"`
	Description *string           `json:"description" validate:"required"`
	Geometry    *geojson.Geometry `json:"geometry" validate:"required,linestring"`
}

type RouteService interface {
	CreateRoute(ctx context.Context, routeCreationDTO RouteCreationDTO) (*models.Route, error)
	GetRouteByID(ctx context.Context, routeID string) (*models.Route, error)
	GetRoutes(ctx context.Context) (*[]models.Route, error)
	UpdateRouteByID(ctx context.Context, routeID string, updateRoute UpdateRoute) (*models.Route, error)
	DeleteRouteByID(ctx context.Context, routeID string) error
}
