package middlewares

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/redis"
	redisRepo "github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/redis/repository"
	"net/http"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := r.Clone(r.Context())

		response := NewErrorResponse()

		var cfg config.Config

		// TODO: создать клиент редиса (Или другой базы данных, но в моем случае редиса)
		redisClint, _ := redis.New(cfg)
		// TODO: инициализировать репозиторий сессий с клиентом редиса
		seession, _ := redisRepo.NewSessionRepository(redisClint)

		token := r.Header.Get("Authorization")

		sessionToken, err := seession.GetSession(context.TODO(), token)
		fmt.Println(sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response.Errors["name"] = "Name parameter is required"
		}

		next.ServeHTTP(w, req)
	})
}
