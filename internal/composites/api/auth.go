package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/handlers"
	appMiddleware "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	postgresRepo "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/repository"
)

type AuthComposite struct {
	handler *handlers.AuthHandler
}

func NewAuthComposite(cfg config.Config) (*AuthComposite, error) {
	postgresClient, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}
	repository, err := postgresRepo.NewUserRepository(postgresClient)
	if err != nil {
		return nil, err
	}
	service, err := services.NewAuthService(cfg, repository)
	if err != nil {
		return nil, err
	}
	handler, err := handlers.NewAuth(cfg, service)
	if err != nil {
		return nil, err
	}

	return &AuthComposite{
		handler,
	}, nil
}

func (c *AuthComposite) Register(router chi.Router) {
	router.With(appMiddleware.ValidateRegistrationParams).Post("/register", c.handler.Register())
	router.With(appMiddleware.ValidateLoginParams).Post("/login", c.handler.LogIn())
}
