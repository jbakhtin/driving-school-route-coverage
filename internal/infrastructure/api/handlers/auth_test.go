package handlers

import (
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	authService "github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/api/middleware"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/repository"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestAuthHandler_LogIn(t *testing.T) {
	cfg, err := config.GetConfig()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

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
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	type args struct {
		RequestParams map[string]string
	}

	type want struct {
		StatusCode int
		ContentType string
		Body map[string]any
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
					"login": "test_login",
					"password": "test_password",
				},
			},
			want{
				200,
				"application/json",
				map[string]any{
					"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0M30.1-0ZrVKyo27KhzeiKCCItCK4gr2tGXDqjOhsCaQRkFs",
				},
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
				"application/json",
				map[string]any{
					"message": "User doesn't exist with this login",
					"code": "001",
				},
			},
		},
		{
			"Parameters not passed",
			args{
				map[string]string{
					"login": "leperiton11",
				},
			},
			want{
				400,
				"application/json",
				map[string]any{
					"message": "Bad request params",
					"code": "004",
					"errors": map[string]string{
						"password": "Password parameter is required",
					},
				},
			},
		},
		{
			"Parameters not passed",
			args{
				map[string]string{
					"password": "zetati16",
				},
			},
			want{
				400,
				"application/json",
				map[string]any{
					"message": "Bad request params",
					"code": "004",
					"errors": map[string]string{
						"login": "Login parameter is required",
					},
				},
			},
		},
		{
			"Parameters not passed",
			args{
			},
			want{
				400,
				"application/json",
				map[string]any{
					"message": "Bad request params",
					"code": "004",
					"errors": map[string]string{
						"login": "Login parameter is required",
						"password": "Password parameter is required",
					},
				},
			},
		},
	}
	h := &AuthHandler{
		service: authService,
		logger:  logger,
		config:  cfg,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := middleware.ValidateLoginParams(h.LogIn())

			rec := httptest.NewRecorder()
			defer rec.Result().Body.Close()

			buf, _ := json.Marshal(tt.args.RequestParams)
			body := strings.NewReader(string(buf))

			req, _ := http.NewRequest("POST", "/login", body)
			defer req.Body.Close()

			handler.ServeHTTP(rec, req)

			response := rec.Result()
			defer response.Body.Close()
			if response.StatusCode != tt.want.StatusCode {
				t.Errorf("The status code does not match the expected one. Want %v, received %v.", tt.want.StatusCode, response.StatusCode)
			}

			if rec.Header().Get("Content-Type") != tt.want.ContentType {
				t.Errorf("The content type does not match the expected one. Want %v, received %v.", tt.want.ContentType, rec.Header().Get("Content-Type"))
			}

			responseMap := make(map[string]any)
			json.Unmarshal(rec.Body.Bytes(), &responseMap)

			test, _ := json.Marshal(responseMap)

			test2, _ := json.Marshal(tt.want.Body)

			if !reflect.DeepEqual(test, test2) {
				t.Errorf("The body does not match the expected one. Want %v, received %v.", test2, test)
			}
		})
	}
}

func TestAuthHandler_Register(t *testing.T) {
	cfg, err := config.GetConfig()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

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
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	type args struct {
		RequestParams map[string]string
	}

	type want struct {
		StatusCode int
		ContentType string
		Body map[string]any
	}

	h := &AuthHandler{
		service: authService,
		logger:  logger,
		config:  cfg,
	}

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
					"email": "test_user@yandex.ru",
					"login": "test_user",
					"password": "zetati16",
					"password_confirmation": "zetati16",
				},
			},
			want{
				200,
				"application/json",
				map[string]any{
					"message": "User created",
				},
			},
		},
		{
			"Parameters not passed",
			args{
				map[string]string{
					"lastname": "Бахтин",
					"name": "Юрий",
					"email": "leperiton11@yandex.ru",
					"password": "zetati16",
					"password_confirmation": "zetati16",
				},
			},
			want{
				400,
				"application/json",
				map[string]any{
					"message": "Bad request params",
					"code": "003",
					"errors": map[string]string{
						"login": "Login parameter is required",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := middleware.ValidateRegistrationParams(h.Register())

			rec := httptest.NewRecorder()

			buf, _ := json.Marshal(tt.args.RequestParams)
			body := strings.NewReader(string(buf))

			req, _ := http.NewRequest("POST", "/register", body)
			defer req.Response.Body.Close()

			handler.ServeHTTP(rec, req)

			response := rec.Result()
			defer response.Body.Close()
			if response.StatusCode != tt.want.StatusCode {
				t.Errorf("The status code does not match the expected one. Want %v, received %v.", tt.want.StatusCode, response.StatusCode)
			}

			if rec.Header().Get("Content-Type") != tt.want.ContentType {
				t.Errorf("The content type does not match the expected one. Want %v, received %v.", tt.want.ContentType, rec.Header().Get("Content-Type"))
			}

			responseMap := make(map[string]any)
			json.Unmarshal(rec.Body.Bytes(), &responseMap)

			test, _ := json.Marshal(responseMap)

			test2, _ := json.Marshal(tt.want.Body)

			if !reflect.DeepEqual(test, test2) {
				t.Errorf("The body does not match the expected one. Want %v, received %v.", test2, test)
			}
		})
	}
}