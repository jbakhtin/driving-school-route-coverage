package apperror

import (
	"errors"
	"net/http"
)

type errorHandler func(w http.ResponseWriter, r *http.Request) error

func Handler(h errorHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var appErr *AppError
		err := h(writer, request)

		if err != nil {
			errors.As(err, &appErr)
			switch appErr {
			case NotFound:
				writer.WriteHeader(http.StatusNotFound)
			case UserAlreadyExists:
				writer.WriteHeader(http.StatusConflict)
			default:
				appErr = systemError(err)
				writer.WriteHeader(http.StatusBadRequest)
			}

			writer.Write(appErr.Marshal())
		}
	}
}
