package pingsql

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

func (h *Handler) PingSQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.storage.Ping()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}