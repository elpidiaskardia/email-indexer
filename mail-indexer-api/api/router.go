package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// ConfigureMiddleware configures the middleware to handle CORS (Cross-Origin Resource Sharing).
func ConfigureMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			//Configuration CORS
			writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if request.Method == "OPTIONS" {
				writer.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(writer, request)
		})
	}
}

// ConfigureRouter sets up the router with middleware and routes.
func ConfigureRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(ConfigureMiddleware())

	router.Post("/api/getEmailsByParameter", GetEmailByParameter)
	router.Post("/api/getEmails", GetAllEmails)

	return router
}

// NewRouter returns an HTTP handler configured with the router.
func NewRouter() http.Handler {
	return ConfigureRouter()
}
