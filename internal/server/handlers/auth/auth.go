package auth

import (
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/storages/sqlstorage/entities/user"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	storage *user.Storage
	logger *zap.Logger
	config *config.Config
}

func NewHandler(cfg config.Config) (*Handler, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	storage, err := user.NewStorage(cfg)
	if err != nil {
		return nil, err
	}

	return &Handler{
		storage: storage,
		logger: logger,
		config: &cfg,
	}, nil
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) LogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}