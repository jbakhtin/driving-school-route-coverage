package handlers

import (
	"context"
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/types"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
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
		ctx = context.WithValue(ctx, types.ContextKeyUserID, r.Context().Value("user_id"))

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		var routeCreationDTO ifaceservice.RouteCreationDTO
		err = json.Unmarshal(body, &routeCreationDTO)
		if err != nil {
			return err
		}

		route, err := h.service.CreateRoute(ctx, routeCreationDTO)
		if err != nil {
			return err
		}

		buffer, err := json.Marshal(route)
		if err != nil {
			return err
		}

		_, err = w.Write(buffer)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusCreated)
		return nil
	}

	return apperror.Handler(fn)
}

func (h *RouteHandler) Get(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		return nil
	}

	return apperror.Handler(fn)
}

func (h *RouteHandler) Show(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		ctx = context.WithValue(ctx, types.ContextKeyUserID, r.Context().Value("user_id"))

		routeID := chi.URLParam(r, "routeID")

		route, err := h.service.GetRouteByID(ctx, routeID)
		if err != nil {
			return err
		}

		buffer, err := json.Marshal(route)
		if err != nil {
			return err
		}

		_, err = w.Write(buffer)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}

	return apperror.Handler(fn)
}

func (h *RouteHandler) Update(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		ctx = context.WithValue(ctx, types.ContextKeyUserID, r.Context().Value("user_id"))

		routeID := chi.URLParam(r, "routeID")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		var updateRoute ifaceservice.UpdateRoute
		err = json.Unmarshal(body, &updateRoute)
		if err != nil {
			return err
		}

		route, err := h.service.UpdateRouteByID(ctx, routeID, updateRoute)
		if err != nil {
			return err
		}

		buffer, err := json.Marshal(route)
		if err != nil {
			return err
		}

		_, err = w.Write(buffer)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}

	return apperror.Handler(fn)
}

func (h *RouteHandler) Delete(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		ctx = context.WithValue(ctx, types.ContextKeyUserID, r.Context().Value("user_id"))

		routeID := chi.URLParam(r, "routeID")

		err := h.service.DeleteRouteByID(ctx, routeID)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}

	return apperror.Handler(fn)
}
