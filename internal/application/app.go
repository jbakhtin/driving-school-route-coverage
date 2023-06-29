package application

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/composites/api"
	appMiddleware "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/middleware"
	"net/http"
)

type Server struct {
	*http.Server
	config *config.Config
}

func New(cfg config.Config) (*Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.NotFound(notFound())

	authComposite, _ := api.NewAuthComposite(cfg)
	authComposite.Register(r)

	r.Group(func(r chi.Router) {
		r.Use(appMiddleware.CheckAuth)

		r.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("test")
		})
	})

	server := http.Server{
		Addr: cfg.ServerAddress,
		Handler: r,
	}

	return &Server{
		&server,
		&cfg,
	}, nil
}

// notFound возврат 404 ошибки в общем стиле в json формате
func notFound() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		return apperror.NotFound
	}

	return apperror.Handler(fn)
}

func (s *Server) Start() error {
	var err error

	go func() {
		err = s.ListenAndServe()
	}()

	return err
}
