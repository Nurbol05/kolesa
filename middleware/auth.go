package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"kolesa/pkg/logger"
	"net/http"
	"os"
	"strings"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET")) // TODO: замените на секрет из config или переменной

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Infof("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware проверяет JWT токен и пропускает только авторизованных пользователей
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Разрешаем доступ без авторизации к /login и /register
		if path == "/login" || path == "/register" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		logger.Log.Infof("Authorized request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
