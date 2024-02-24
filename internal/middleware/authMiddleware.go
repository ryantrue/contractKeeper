package middleware

import (
	"context"
	"github.com/ryantrue/contractKeeper.git/internal/auth"
	"net/http"
	"strings"
)

// AuthMiddleware проверяет наличие и валидность JWT токена в запросе
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Извлечение токена из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Проверка формата заголовка
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Верификация токена
		claims, err := auth.AuthenticateToken(bearerToken[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Вставка claims из токена в контекст запроса для использования в последующих обработчиках
		ctx := context.WithValue(r.Context(), "userClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
