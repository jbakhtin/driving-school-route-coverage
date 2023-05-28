package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"io"
	"net/http"
)


func ValidateLoginParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := r.Clone(r.Context())

		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body.Close() //  must close
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		response := NewErrorResponse()
		request := services.UserLoginRequest{}
		_ = json.Unmarshal(bodyBytes, &request)

		if request.Login == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["login"] = "Login parameter is required"
		}

		if request.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Errors["password"] = "Password parameter is required"
		}

		if len(response.Errors) > 0 {
			responseBytes, _ := json.Marshal(response)
			w.Write(responseBytes)
			return
		}

		next.ServeHTTP(w, req)
	})
}
