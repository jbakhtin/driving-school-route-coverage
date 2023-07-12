package middleware

import (
	"errors"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(next http.Handler) http.Handler {
	test := func(w http.ResponseWriter, r *http.Request) error {
		config, err := config.GetConfig()
		if err != nil {
			return err
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			return apperror.New(errors.New("parameters not passed"), "Unauthorized", apperror.SystemErrorCode, "", nil)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.AppKey), nil
		})
		if err != nil || !token.Valid {
			return apperror.New(err, "Unauthorized", "000", "", nil)
		}
		next.ServeHTTP(w, r)
		return nil
	}

	return apperror.Handler(test)
}
