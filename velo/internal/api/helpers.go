package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/velo-api/velo/internal/auth"
	"github.com/velo-api/velo/internal/storage"
)

type API struct {
	storage     *storage.Engine
	authService *auth.Auth
	jwtSecret   string
	mu          sync.RWMutex
}

func New(store *storage.Engine) *API {
	return &API{storage: store}
}

func (a *API) SetAuth(authService *auth.Auth, jwtSecret string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.authService = authService
	a.jwtSecret = jwtSecret
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

type PaginatedResponse struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total,omitempty"`
	Page  int         `json:"page,omitempty"`
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, ErrorResponse{Error: message, Code: status})
}

func internalError(err error, msg string) string {
	log.Printf("[ERROR] %s: %v", msg, err)
	return "internal server error"
}

func respondSuccess(w http.ResponseWriter, data interface{}) {
	respondJSON(w, http.StatusOK, SuccessResponse{Data: data})
}

func respondCreated(w http.ResponseWriter, data interface{}) {
	respondJSON(w, http.StatusCreated, SuccessResponse{Data: data})
}

func sanitizeUser(u storage.User) map[string]interface{} {
	return map[string]interface{}{
		"id":        u.ID,
		"name":      u.Name,
		"email":     u.Email,
		"avatar":    u.Avatar,
		"createdAt": u.CreatedAt,
	}
}

func extractID(path, prefix string) string {
	parts := strings.Split(path, "/")
	for i, p := range parts {
		if p == prefix && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
