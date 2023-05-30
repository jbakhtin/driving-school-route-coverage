package application

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/handlers"
	appMiddleware "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	postgresRepo "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/repository"
	"net/http"
)

type Server struct {
	*http.Server
	config *config.Config
}

func New(cfg config.Config) (*Server, error) {
	server := http.Server{
		Addr: cfg.ServerAddress,
	}

	return &Server{
		&server,
		&cfg,
	}, nil
}

// notFound возврат 404 ошибки в общем стиле в json формате
func notFound(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	return apperror.NotFound
}

func (s *Server) Start(ctx context.Context) error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.NotFound(apperror.Handler(notFound))

	// TODO: объединить компоновщиком
	postgresClient, err := postgres.New(*s.config)
	if err != nil {
		return err
	}
	repo, err := postgresRepo.NewUserRepository(postgresClient)
	if err != nil {
		return err
	}
	userService, err := services.NewAuthService(repo)
	if err != nil {
		return err
	}
	handlersList, err := handlers.NewAuth(*s.config, userService)
	if err != nil {
		return err
	}

	r.Route("/", func(r chi.Router) {

		// TODO: вынести список роутов в отдельный файл
		r.With(appMiddleware.ValidateRegistrationParams).Post("/register", apperror.Handler(handlersList.Register))
		r.With(appMiddleware.ValidateLoginParams).Post("/login", apperror.Handler(handlersList.LogIn))

		r.Group(func(r chi.Router) {
			r.Use(appMiddleware.CheckAuth)

			//r.Route("/areas", func(r chi.Router) {
			//	r.Post("/", handlersList.CreateArea())
			//	r.Get("/", handlersList.GetAreas())
			//
			//	r.Route("/{id}", func(r chi.Router) {
			//		r.Get("/", handlersList.GetAreaById())
			//		r.Put("/", handlersList.UpdateArea())
			//		r.Delete("/", handlersList.DeleteArea())
			//
			//		r.Get("/points", handlersList.GetAreaPoints())
			//	})
			//})

			//r.Route("/routes", func(r chi.Router) {
			//	r.Post("/", handlersList.CreateRoute())
			//	r.Get("/", handlersList.GetRoutes())
			//
			//	r.Route("/{id}", func(r chi.Router) {
			//		r.Get("/", handlersList.GetRouteById())
			//		r.Put("/", handlersList.UpdateRoute())
			//		r.Delete("/", handlersList.DeleteRoute())
			//
			//		r.Get("/points", handlersList.GetRoutePoints()) // Получить точки по конкретному маршруту
			//	})
			//})
			//
			//r.Route("/marks", func(r chi.Router) {
			//	r.Post("/", handlersList.CreateMark())
			//	r.Get("/", handlersList.GetMarks())
			//
			//	r.Route("/{id}", func(r chi.Router) {
			//		r.Get("/", handlersList.GetMarkById())
			//		r.Put("/", handlersList.UpdateMark())
			//		r.Delete("/", handlersList.DeleteMark())
			//
			//		r.Get("/point", handlersList.GetMarkPoints()) // Получить точку по конкретной марке
			//	})
			//})
		})
	})

	err = http.ListenAndServe(s.Addr, r)
	if err != nil {
		return err
	}

	return nil
}
