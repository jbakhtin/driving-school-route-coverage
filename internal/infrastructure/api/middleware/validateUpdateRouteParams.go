package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	ifaceservice "github.com/jbakhtin/driving-school-route-coverage/internal/interfaces/services"
	"io"
	"net/http"
)

func ValidateUpdateRouteParams(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		validate := validator.New()
		err := validate.RegisterValidation("linestring", LineString)
		if err != nil {
			return err
		}

		req := r.Clone(r.Context())

		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body.Close()
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		errsList := map[string]string{}

		request := ifaceservice.UpdateRoute{}
		err = json.Unmarshal(bodyBytes, &request)
		if err != nil {
			return err
		}

		err = validate.Struct(request)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)

			for _, vError := range validationErrors {
				errsList[vError.Field()] = vError.Error()
			}
		}

		if len(errsList) > 0 {
			return apperror.New(nil, "Bad request params", apperror.BadRequestParamsCode, "", errsList)
		}

		next.ServeHTTP(w, req)
		return nil
	}

	return apperror.Handler(fn)
}
