package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/handlers"
	appMiddleware "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	postgresRepo "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/repository"
)

type RouteComposite struct {
	handler *handlers.RouteHandler
}

func NewRouteComposite(cfg config.Config) (*RouteComposite, error) {
	postgresClient, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}

	repository, err := postgresRepo.NewRouteRepository(postgresClient)
	if err != nil {
		return nil, err
	}

	service, err := services.NewRouteService(cfg, repository)
	if err != nil {
		return nil, err
	}

	handler, err := handlers.NewRouteHandler(cfg, service)
	if err != nil {
		return nil, err
	}

	return &RouteComposite{
		handler,
	}, nil
}

func (c *RouteComposite) Register(ctx context.Context, router chi.Router) {
	router.Route("/routes", func(r chi.Router) {
		r.Use(appMiddleware.CheckAuth)

		r.Get("/", c.handler.Get(ctx))
		r.With(appMiddleware.ValidateCreateRouteParams).Post("/", c.handler.Create(ctx))

		r.Route("/{routeID}", func(r chi.Router) {
			r.Get("/", c.handler.Show(ctx))
			r.With(appMiddleware.ValidateUpdateRouteParams).Put("/", c.handler.Update(ctx))
			r.Delete("/", c.handler.Delete(ctx))
		})
	})
}
