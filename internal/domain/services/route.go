package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
)

type LineString struct {
	Type string `json:"type" validate:"required,eq='LineString'"`
	Coordinates [][]float64 `json:"coordinates" validate:"required,linestring"`
}

type RouteCreationDTO struct {
	Name string `json:"name" validate:"required"`
	Line LineString
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

func (us *RouteService) CreateRoute(ctx context.Context, routeCreationDto RouteCreationDTO) (*RouteCreatedDTO, error) {
	return nil, errors.New("can not create route")
}

