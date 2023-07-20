package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
)

type LineString struct {
	Type string `json:"type" validate:"required,eq='LineString'"`
	Coordinates []geom.Coord `json:"coordinates" validate:"required,linestring"`
}

type RouteCreationDTO struct {
	Name string `json:"name" validate:"required"`
	Line *geojson.Geometry `json:"geometry" validate:"required,linestring"`
}

type RouteCreatedDTO struct {
	Message string `json:"message,omitempty"`
}

func (e *RouteCreatedDTO) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return marshal, nil
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

func (us *RouteService) CreateRoute(ctx context.Context, routeCreationDto RouteCreationDTO) (*models.Route, error) {
	bytes, _ := json.Marshal(routeCreationDto.Line)

	routeCreation := repositories.RouteCreation{
		bytes,
	}

	route, err := us.repo.CreateRoute(ctx, routeCreation)
	if err != nil {
		return nil, err
	}

	fmt.Println(route)

	return route, nil
}

func (us *RouteService) GetRouteByID(ctx context.Context, routeID string) (*models.Route, error) {
	route, err := us.repo.GetRouteByID(ctx, routeID)
	if err != nil {
		return nil, err
	}

	return route, nil
}

