package auth

import (
	"net/http"
	"strings"

	"github.com/velo-api/velo/pkg/config"
)

type Auth struct {
	config config.AuthConfig
}

func New(cfg config.AuthConfig) *Auth {
	return &Auth{config: cfg}
}

func (a *Auth) Validate(r *http.Request) bool {
	for _, provider := range a.config.Providers {
		switch provider.Type {
		case "jwt":
			if a.validateJWT(r, provider) {
				return true
			}
		case "apikey":
			if a.validateAPIKey(r, provider) {
				return true
			}
		}
	}
	return false
}

func (a *Auth) validateJWT(r *http.Request, provider config.AuthProvider) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return false
	}

	token := parts[1]
	return token != ""
}

func (a *Auth) validateAPIKey(r *http.Request, provider config.AuthProvider) bool {
	header := provider.Header
	if header == "" {
		header = "X-API-Key"
	}

	apiKey := r.Header.Get(header)
	return apiKey != ""
}
