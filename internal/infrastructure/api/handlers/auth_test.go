package handlers

import (
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	authService "github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/repository"
	"github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestAuthHandler_LogIn(t *testing.T) {
	cfg := config.GetConfig()
	pgClient, err := postgres.New(*cfg)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	repo, err := repository.NewUserRepository(pgClient)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	authService, err := authService.NewAuthService(*cfg, repo)

	type args struct {
		RequestParams map[string]string
	}

	type want struct {
		StatusCode int
	}
	var tests = []struct {
		name string
		args args // TODO: параметры передаваемые в хендлер
		want want
	}{
		{
			"User Success Authorized",
			args{
				map[string]string{
					"login": "leperiton11",
					"password": "zetati16",
				},
			},
			want{
				200,
			},
		},
		{
			"User Not Found",
			args{
				map[string]string{
					"login": "test",
					"password": "test",
				},
			},
			want{
				400,
			},
		},
		{
			"Parameters not passed",
			args{
				map[string]string{
					"login": "test",
				},
			},
			want{
				400,
			},
		},
		{
			"Parameters not passed",
			args{
				map[string]string{
					"password": "test",
				},
			},
			want{
				400,
			},
		},
		{
			"Parameters not passed",
			args{
			},
			want{
				400,
			},
		},
	}
	h := &AuthHandler{
		service: authService,
		logger:  logger,
		config:  cfg,
	}
	handler := http.HandlerFunc(apperror.Handler(h.LogIn))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			buf, _ := json.Marshal(tt.args.RequestParams)
			body := strings.NewReader(string(buf))

			req, _ := http.NewRequest("POST", "/login", body)

			handler.ServeHTTP(rec, req)

			if rec.Result().StatusCode != tt.want.StatusCode {
				t.Errorf("The test is not passed, the test result does not satisfy the expected.")
			}
		})
	}
}

func TestAuthHandler_Register(t *testing.T) {
	cfg := config.GetConfig()
	pgClient, err := postgres.New(*cfg)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	repo, err := repository.NewUserRepository(pgClient)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	authService, err := authService.NewAuthService(*cfg, repo)

	type args struct {
		RequestParams map[string]string
	}

	type want struct {
		StatusCode int
	}

	h := &AuthHandler{
		service: authService,
		logger:  logger,
		config:  cfg,
	}
	handler := http.HandlerFunc(apperror.Handler(h.Register))

	tests := []struct {
		name    string
		args    args
		want want
	}{
		{
			"User successful registered",
			args{
				map[string]string{
					"lastname": "Бахтин",
					"name": "Юрий",
					"email": "leperiton11@yandex.ru",
					"login": "leperiton11",
					"password": "zetati16",
					"password_confirmation": "zetati16",
				},
			},
			want{
				200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			buf, _ := json.Marshal(tt.args.RequestParams)
			body := strings.NewReader(string(buf))

			req, _ := http.NewRequest("POST", "/register", body)

			handler.ServeHTTP(rec, req)

			if rec.Result().StatusCode != tt.want.StatusCode {
				t.Errorf("The test is not passed, the test result does not satisfy the expected.")
			}
		})
	}
}

func TestNewAuth(t *testing.T) {
	type args struct {
		cfg     config.Config
		service services.AuthService
	}
	tests := []struct {
		name    string
		args    args
		want    *AuthHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuth(tt.args.cfg, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuth() got = %v, want %v", got, tt.want)
			}
		})
	}
}
