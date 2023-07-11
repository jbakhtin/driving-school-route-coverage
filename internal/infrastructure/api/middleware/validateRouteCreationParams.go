package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/services"
	"io"
	"net/http"
	"regexp"
)

func ValidateRouteCreationParams(next http.Handler) http.Handler {
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

		request := services.RouteCreationDTO{}
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

func LineString(fl validator.FieldLevel) bool {
	coordinates := fl.Field().Interface().([][]float64)
	// Проверяем, что в LineString есть хотя бы две точки
	if len(coordinates) < 2 {
		return false
	}

	// Проверяем формат каждой координаты (широта, долгота)
	coordRegex := regexp.MustCompile(`^-?\d+(\.\d+)?,-?\d+(\.\d+)?$`)
	for _, c := range coordinates {
		if !coordRegex.MatchString(fmt.Sprintf("%f,%f", c[0], c[1])) {
			return false
		}
	}
	return true
}

