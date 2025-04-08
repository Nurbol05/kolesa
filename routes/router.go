package routes

import (
	"github.com/gorilla/mux"

	"kolesa/auth"
	"kolesa/middleware"
)

func SetupRoutes(
	authService *auth.AuthService,
) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("/register", auth.RegisterHandler(authService)).Methods("POST")
	r.HandleFunc("/login", auth.LoginHandler(authService)).Methods("POST")

	return r
}
