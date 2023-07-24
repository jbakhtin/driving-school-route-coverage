package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/utils/mailer"
	"go.uber.org/zap"
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

func (h *AuthHandler) Register(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		var request ifaceservice.UserRegistrationRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			return err
		}

		registerResponse, err := h.service.RegisterUser(ctx, request)
		if err != nil {
			return err
		}

		mail := mailer.NewMail(request.Email, "Successful Registered", "You are successful registered.")
		mailsQueue, err := mailer.GetMailsQueue()
		if err != nil {
			return err
		}
		mailsQueue <- *mail

		registerResponseJSON, err := json.Marshal(registerResponse)
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

func (h *AuthHandler) LogIn(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		var request ifaceservice.UserLoginRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			return err
		}

		response, err := h.service.LoginUser(ctx, request)
		if err != nil {
			return err
		}

		bytes, err := json.Marshal(response)
		if err != nil {
			return err
		}

		_, err = w.Write(bytes)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}

	return apperror.Handler(fn)
}
