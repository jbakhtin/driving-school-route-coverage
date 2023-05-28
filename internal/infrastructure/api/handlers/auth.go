package handlers

import (
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

type AuthHandler struct {
	service ifaceservice.AuthService
	logger  *zap.Logger
	config  *config.Config
}

func New(cfg config.Config, service ifaceservice.AuthService) (*AuthHandler, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &AuthHandler{
		service: service,
		logger:  logger,
		config:  &cfg,
	}, nil
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json2")

	body, _ := io.ReadAll(r.Body)
	var request services.UserRegistrationRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return err
	}

	_, err = h.service.RegisterUser(request)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates") {
			return apperror.UserAlreadyExists
		}

		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	body, _ := io.ReadAll(r.Body)
	var request services.UserLoginRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return err
	}

	response, err := h.service.LoginUser(request)
	if err != nil {
		return err
	}

	w.Write(response.Marshal())
	w.WriteHeader(http.StatusOK)
	return nil
}
