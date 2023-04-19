package mark

import (
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/storages/sqlstorage/entities/mark"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	storage *mark.Storage
	logger *zap.Logger
}

func NewHandler(cfg config.Config) (*Handler, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	storage, err := mark.NewStorage(cfg)
	if err != nil {
		return nil, err
	}

	return &Handler{
		storage: storage,
		logger: logger,
	}, nil
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) GetPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) GetPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}