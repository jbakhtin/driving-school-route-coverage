package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"io"
	"net/http"

	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
)

func ValidateLoginParams(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		req := r.Clone(r.Context())

		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body.Close() //  must close
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		errsList := map[string]string{}
		request := services.UserLoginRequest{}
		_ = json.Unmarshal(bodyBytes, &request)

		if request.Login == "" {
			errsList["login"] = "Login parameter is required"
		}

		if request.Password == "" {
			errsList["password"] = "Password parameter is required"
		}

		if len(errsList) > 0 {
			return apperror.New(nil, "Bad request params", "004", "", errsList)
		}

		next.ServeHTTP(w, req)
		return nil
	}

	return apperror.Handler(fn)
}
