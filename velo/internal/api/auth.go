package api

import (
	"encoding/json"
	"net/http"

	"github.com/velo-api/velo/internal/auth"
	"github.com/velo-api/velo/internal/storage"
)

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

	if !storage.CheckPassword(user.Password, req.Password) {
		respondError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	a.mu.RLock()
	secret := a.jwtSecret
	a.mu.RUnlock()

	token, err := a.authService.GenerateToken(user.ID, user.Email, secret)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "failed to generate token")
		return
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

	hashedPassword, err := storage.HashPassword(req.Password)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "failed to hash password")
		return
	}

	id := a.storage.NextID("user:")
	user := &storage.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := a.storage.CreateUser(user); err != nil {
		respondError(w, http.StatusInternalServerError, internalError(err, "create user"))
		return
	}

	a.mu.RLock()
	secret := a.jwtSecret
	a.mu.RUnlock()

	token, err := a.authService.GenerateToken(user.ID, user.Email, secret)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	respondCreated(w, map[string]interface{}{
		"token": token,
		"user":  sanitizeUser(*user),
	})
}

func (a *API) ValidateToken(token string) bool {
	a.mu.RLock()
	secret := a.jwtSecret
	a.mu.RUnlock()

	_, err := a.authService.ValidateToken(token, secret)
	return err == nil
}

func ExtractToken(r *http.Request) string {
	return auth.ExtractToken(r)
}
