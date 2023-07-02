package handlers

import (
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/utils/mailer"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type AuthHandler struct {
	service ifaceservice.AuthService
	logger  *zap.Logger
	config  *config.Config
}

func NewAuth(cfg config.Config, service ifaceservice.AuthService) (*AuthHandler, error) {
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

func (h *AuthHandler) Register() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")

		body, _ := io.ReadAll(r.Body)
		var request services.UserRegistrationRequest
		err := json.Unmarshal(body, &request)
		if err != nil {
			return err
		}

		registerResponse, err := h.service.RegisterUser(request)
		if err != nil {
			return err
		}

		mail := mailer.NewMail(request.Email, "Successful Registered", "You are successful registered.")
		mailsQueue, err := mailer.GetMailsQueue()
		if err != nil {
			return err
		}
		mailsQueue <- *mail

		registerResponseJSON, err := registerResponse.Marshal()
		if err != nil {
			return err
		}

		_, err = w.Write(registerResponseJSON)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusCreated)
		return nil
	}

	return apperror.Handler(fn)
}

func (h *AuthHandler) LogIn() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {

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

	return apperror.Handler(fn)
}
