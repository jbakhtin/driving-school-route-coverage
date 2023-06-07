package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"io"
	"net/http"

	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
)

func ValidateRegistrationParams(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		req := r.Clone(r.Context())

		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body.Close()
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		errsList := map[string]string{}
		request := services.UserRegistrationRequest{}
		_ = json.Unmarshal(bodyBytes, &request)

		if request.Name == "" {
			errsList["name"] = "Name parameter is required"
		}

		if request.Lastname == "" {
			errsList["lastname"] = "Lastname parameter is required"
		}

		if request.Email == "" {
			errsList["email"] = "Email parameter is required"
		}

		if request.Login == "" {
			errsList["login"] = "Login parameter is required"
		}

		if request.Password == "" {
			errsList["password"] = "Password parameter is required"
		}

		if request.PasswordConfirmation == "" {
			errsList["password_confirmation"] = "Password confirmation parameter is required"
		}

		if request.PasswordConfirmation != request.Password {
			errsList["password_confirmation"] = "Password confirmation don't match with password"
		}

		if len(errsList) > 0 {
			return apperror.New(nil, "Bad request params", apperror.BadRequestParamsCode, "", errsList)
		}

		next.ServeHTTP(w, req)
		return nil
	}

	return apperror.Handler(fn)
}
