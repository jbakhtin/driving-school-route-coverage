package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем JWT-токен из заголовка запроса
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Проверяем токен на валидность
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Здесь нужно вернуть секретный ключ, который использовался для подписи токена
			return []byte("test"), nil //TODO: прокинуть конфиг
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Если токен верный, разрешаем доступ к запрашиваемому ресурсу
		next.ServeHTTP(w, r)
	})
}
