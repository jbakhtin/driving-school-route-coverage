package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
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
