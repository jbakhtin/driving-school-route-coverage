package application

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/handlers"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/middlewares"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	postgresRepo "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/repository"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/redis"
	redisRepo "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/redis/repository"
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

func (s *Server) Start(ctx context.Context) error {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	client, err := postgres.New(*s.config)
	repo, err := postgresRepo.NewUserRepository(client)

	// TODO: создать клиент редиса (Или другой базы данных, но в моем случае редиса)
	redisClint, err := redis.New(*s.config)
	// TODO: инициализировать репозиторий сессий с клиентом редиса
	seession, err := redisRepo.NewSessionRepository(redisClint)

	userService, err := services.NewAuthService(repo, seession) // TODO: передать репозиторий сервис авторизации

	// TODO: создать миделвеер для проверки аутентификации пользваотеля
	// TODO: передать репозиторий сервис миделвеер

	handlersList, err := handlers.New(*s.config, userService)

	r.Route("/", func(r chi.Router) {
		r.With(middlewares.ValidateRegisterRequest).Post("/register", apperror.Middleware(handlersList.Register))
		r.Post("/login", apperror.Middleware(handlersList.LogIn))


		r.Group(func(r chi.Router) {
			r.Use(middlewares.CheckAuth)
			// TODO: добавить проверку авторизацию

			r.Post("/logout", handlersList.LogOut())

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

		r.Route("/ping/", func(r chi.Router) {
			//r.Get("/postgres", handlersList.PingSQL())
		})
	})

	err = http.ListenAndServe(s.Addr, r)
	if err != nil {
		return err
	}

	return nil
}
