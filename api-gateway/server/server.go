package server

import (
	"io"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		io.WriteString(w, "Hello World!")
	})

	registerRoutes(mux)

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: cors.AllowAll().Handler(mux), // TODO aca agregar otros middlewares
	}

	server.ListenAndServe()
}
