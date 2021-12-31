package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func LoadCORS(r *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Request", "Content-Type", "Location", "Entity", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(headers, methods, origins)(r)

}