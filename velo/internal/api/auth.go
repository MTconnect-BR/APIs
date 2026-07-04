package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/velo-api/velo/internal/storage"
)

var sessions = make(map[string]Session)

type Session struct {
	UserID    string
	ExpiresAt time.Time
}

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req storage.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		respondError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	user, err := a.storage.GetUserByEmail(req.Email)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	if user.Password != req.Password {
		respondError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token := generateToken()
	sessions[token] = Session{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	respondSuccess(w, map[string]interface{}{
		"token": token,
		"user":  sanitizeUser(*user),
	})
}

func (a *API) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req storage.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		respondError(w, http.StatusBadRequest, "name, email and password are required")
		return
	}

	existing, _ := a.storage.GetUserByEmail(req.Email)
	if existing != nil {
		respondError(w, http.StatusConflict, "email already exists")
		return
	}

	id := a.storage.NextID("user:")
	user := &storage.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := a.storage.CreateUser(user); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	token := generateToken()
	sessions[token] = Session{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	respondCreated(w, map[string]interface{}{
		"token": token,
		"user":  sanitizeUser(*user),
	})
}

func (a *API) ValidateToken(token string) bool {
	session, exists := sessions[token]
	if !exists {
		return false
	}
	return time.Now().Before(session.ExpiresAt)
}

func ExtractToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return ""
}
