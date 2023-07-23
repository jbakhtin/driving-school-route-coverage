package services

import (
	"context"
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"github.com/twpayne/go-geom/encoding/geojson"
)

type RouteCreationDTO struct {
	Name string `json:"name" validate:"required"`
	Line *geojson.Geometry `json:"geometry" validate:"required,linestring"`
}

type RouteCreatedDTO struct {
	Message string `json:"message,omitempty"`
}

type RouteService struct {
	config *config.Config
	repo   repositories.RouteRepository
}

func NewRouteService(cfg config.Config, repo repositories.RouteRepository) (*RouteService, error) {
	return &RouteService{
		config: &cfg,
		repo:   repo,
	}, nil
}

func (us *RouteService) CreateRoute(ctx context.Context, routeCreationDto ifaceservice.RouteCreationDTO) (*models.Route, error) {
	bytes, err := json.Marshal(routeCreationDto.Line)
	if err != nil {
		return nil, err
	}

	createUser := repositories.CreateRoute{
		Name:       routeCreationDto.Name,
		LineString: bytes,
	}

	route, err := us.repo.CreateRoute(ctx, createUser)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func (us *RouteService) GetRouteByID(ctx context.Context, routeID string) (*models.Route, error) {
	route, err := us.repo.GetRouteByID(ctx, routeID)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func (us *RouteService) UpdateRouteByID(ctx context.Context, routeID string, updateRoute ifaceservice.UpdateRoute) (*models.Route, error) {
	bytes, err := json.Marshal(updateRoute.Geometry)
	if err != nil {
		return nil, err
	}

	updateRouteData := repositories.UpdateRoute{
		Name: updateRoute.Name,
		LineString: bytes,
	}

	route, err := us.repo.UpdateRouteByID(ctx, routeID, updateRouteData)
	if err != nil {
		return nil, err
	}

	return route, nil
}