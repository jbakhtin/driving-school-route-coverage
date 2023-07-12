package services

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
)

type RouteService interface {
	CreateRoute(ctx context.Context, routeCreationDTO services.RouteCreationDTO) (*services.RouteCreatedDTO, error)
}
