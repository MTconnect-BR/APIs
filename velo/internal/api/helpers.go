package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/velo-api/velo/internal/storage"
)

type API struct {
	storage *storage.Engine
}

func New(store *storage.Engine) *API {
	return &API{storage: store}
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, ErrorResponse{Error: message, Code: status})
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
