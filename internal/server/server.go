package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/server/handlers/area"
	"github.com/jbakhtin/driving-school-route-coverage/internal/server/handlers/auth"
	"github.com/jbakhtin/driving-school-route-coverage/internal/server/handlers/mark"
	"github.com/jbakhtin/driving-school-route-coverage/internal/server/handlers/route"
	"net/http"
)

type Server struct {
	*http.Server
}

func New(cfg config.Config) (*Server, error) {
	server := http.Server{
		Addr: cfg.ServerAddress,
	}

	return &Server{
		&server,
	}, nil
}

func (s *Server) Run() error {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// handlers
	authHandler, err := auth.NewHandler()
	if err != nil {
		return err
	}

	areaHandler, err := area.NewHandler()
	if err != nil {
		return err
	}

	markHandler, err := mark.NewHandler()
	if err != nil {
		return err
	}

	routeHandler, err := route.NewHandler()
	if err != nil {
		return err
	}

	r.Route("/", func(r chi.Router) {
		r.Post("/register", authHandler.Register())
		r.Post("/login", authHandler.LogIn())
		r.Post("/logout", authHandler.LogOut())

		r.Route("/areas/", func(r chi.Router) {
			r.Post("/", areaHandler.Create())
			r.Get("/", areaHandler.GetAll())

			r.Get("/points", markHandler.Get()) // Получить все точки сгруппированные по областям

			r.Route("/{id}/", func(r chi.Router) {
				r.Get("/", areaHandler.Get())
				r.Put("/", areaHandler.Update())
				r.Delete("/", areaHandler.Delete())

				r.Get("/points", markHandler.Get())
			})
		})

		r.Route("/routes/", func(r chi.Router) {
			r.Post("/", routeHandler.Create())
			r.Get("/", routeHandler.GetAll())

			r.Get("/points", markHandler.Get()) // Получить точки маршрутов сгруппированных по маршрутам

			r.Route("/{id}/", func(r chi.Router) {
				r.Get("/", routeHandler.Get())
				r.Put("/", routeHandler.Update())
				r.Delete("/", routeHandler.Delete())

				r.Get("/points", markHandler.Get()) // Получить точки по конкретному маршруту
			})
		})

		r.Route("/marks/", func(r chi.Router) {
			r.Post("/", markHandler.Create())
			r.Get("/", markHandler.Get())

			r.Get("/points", markHandler.Get()) // Получить все точки сгруппированные по маркам

			r.Route("/{id}/", func(r chi.Router) {
				r.Get("/", markHandler.Get())
				r.Put("/", markHandler.Update())
				r.Delete("/", markHandler.Delete())

				r.Get("/point", markHandler.Get()) // Получить точку по конкретной марке
			})
		})


		r.Route("/ping/", func(r chi.Router) {
			r.Get("/postgres", markHandler.Create())
			r.Get("/mongodb", markHandler.Get())
		})
	})

	err = http.ListenAndServe(s.Addr, r)
	if err != nil {
		return err
	}

	return nil
}