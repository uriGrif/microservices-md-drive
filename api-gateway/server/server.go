package server

import (
	"api-gateway/middleware/auth"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Run() {
	mux := http.NewServeMux()

	// generate default authenticators
	authenticators := auth.GetDefaultAuthenticators()

	registerRoutes(mux, authenticators)

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: cors.AllowAll().Handler(mux), // TODO aca agregar otros middlewares
	}

	server.ListenAndServe()
}
