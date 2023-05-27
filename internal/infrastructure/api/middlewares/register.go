package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Errors map[string]string `json:"errors,omitempty"`
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		Errors: make(map[string]string),
	}
}

func ValidateRegisterRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := r.Clone(r.Context())

		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body.Close() //  must close
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		response := NewErrorResponse()
		request := services.UserRegistrationRequest{}
		_ = json.Unmarshal(bodyBytes, &request)

		if request.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["name"] = "Name parameter is required"
		}

		if request.Lastname == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["lastname"] = "Lastname parameter is required"
		}

		if request.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["email"] = "Email parameter is required"
		}

		if request.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["password"] = "Password parameter is required"
		}

		if request.PasswordConfirmation == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["password_confirmation"] = "Password confirmation parameter is required"
		}

		if request.PasswordConfirmation != request.Password {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["password_confirmation"] = "Password confirmation don't match with password"
		}

		if len(response.Errors) > 0 {
			responseBytes, _ := json.Marshal(response)
			w.Write(responseBytes)
			return
		}

		next.ServeHTTP(w, req)
	})
}
