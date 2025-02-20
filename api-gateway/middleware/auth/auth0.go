package auth

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type Auth0Authenticator struct {
	domain   string
	audience string
}

func (a *Auth0Authenticator) Authenticate(r *http.Request) error {

	// get token from request header
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return errors.New("missing authorization header")
	}
	token := strings.Split(reqToken, "Bearer ")[1]

	// set up validator
	issuerURL, err := url.Parse("https://" + a.domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{a.audience},
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	// validate
	claims, err := jwtValidator.ValidateToken(r.Context(), token)
	if err != nil {
		return errors.New("invalid token: " + err.Error())
	}

	validatedClaims, ok := claims.(*validator.ValidatedClaims)
	if !ok {
		return errors.New("failed to get claims")
	}

	// add userId for the target server
	r.Header.Add("X-Authenticated-User", validatedClaims.RegisteredClaims.Subject)

	return nil
}

func NewAuth0Authenticator() *Auth0Authenticator {
	return &Auth0Authenticator{
		domain:   os.Getenv("AUTH0_DOMAIN"),
		audience: os.Getenv("AUTH0_AUDIENCE"),
	}
}
