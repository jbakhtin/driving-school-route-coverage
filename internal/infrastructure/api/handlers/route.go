package handlers

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"net/http"
)

type RouteHandler struct {
	service ifaceservice.RouteService
	config  *config.Config
}

func NewRouteHandler(cfg config.Config, service ifaceservice.RouteService) (*RouteHandler, error) {

	return &RouteHandler{
		service: service,
		config:  &cfg,
	}, nil
}

func (h *RouteHandler) Create(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)
		return nil
	}

	return apperror.Handler(fn)
}
