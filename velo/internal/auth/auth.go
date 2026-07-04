package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/velo-api/velo/pkg/config"
)

type Auth struct {
	config config.AuthConfig
}

type Claims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func New(cfg config.AuthConfig) *Auth {
	return &Auth{config: cfg}
}

func (a *Auth) GenerateToken(userID, email, secret string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "velo-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (a *Auth) ValidateToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
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
	_, err := a.ValidateToken(token, provider.Secret)
	return err == nil
}

func (a *Auth) validateAPIKey(r *http.Request, provider config.AuthProvider) bool {
	header := provider.Header
	if header == "" {
		header = "X-API-Key"
	}

	apiKey := r.Header.Get(header)
	return apiKey != ""
}

func ExtractToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return ""
}
