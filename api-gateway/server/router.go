package server

import (
	"api-gateway/middleware/auth"
	"api-gateway/services"
	"net/http"
)

func registerRoutes(mux *http.ServeMux, authenticators map[string]auth.Authenticator) {
	services := services.LoadServices()

	for _, s := range services {
		for _, r := range s.Routes {
			proxy := CreateProxy(s.LoadBalancer)
			var handler http.Handler = &proxy
			if r.Authenticator != nil {
				handler = auth.Middleware(authenticators[*r.Authenticator])(&proxy)
			}
			mux.Handle(r.Endpoint, handler)
		}
	}
}
