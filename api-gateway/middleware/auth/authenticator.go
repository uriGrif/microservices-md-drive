package auth

import (
	"api-gateway/middleware"
	"net/http"
)

type Authenticator interface {
	Authenticate(r *http.Request) error
}

func Middleware(a Authenticator) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			err := a.Authenticate(req)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(rw, req)
		})
	}
}

func GetDefaultAuthenticators() map[string]Authenticator {
	auths := make(map[string]Authenticator)
	auths["auth0"] = NewAuth0Authenticator()
	return auths
}
