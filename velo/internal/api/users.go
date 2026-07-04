package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/velo-api/velo/internal/storage"
)

func (a *API) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.listUsers(w, r)
	case http.MethodPost:
		a.createUser(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (a *API) UserByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/v1/"), "/")
	if len(parts) < 2 || parts[0] != "users" {
		respondError(w, http.StatusBadRequest, "invalid path")
		return
	}

	id := parts[1]

	switch r.Method {
	case http.MethodGet:
		a.getUser(w, r, id)
	case http.MethodPut:
		a.updateUser(w, r, id)
	case http.MethodDelete:
		a.deleteUser(w, r, id)
	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (a *API) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.storage.ListUsers()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if users == nil {
		users = []storage.User{}
	}
	sanitized := make([]map[string]interface{}, len(users))
	for i, u := range users {
		sanitized[i] = sanitizeUser(u)
	}
	respondSuccess(w, sanitized)
}

func (a *API) getUser(w http.ResponseWriter, r *http.Request, id string) {
	user, err := a.storage.GetUser(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "user not found")
		return
	}
	respondSuccess(w, sanitizeUser(*user))
}

func (a *API) createUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

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

	respondCreated(w, sanitizeUser(*user))
}

func (a *API) updateUser(w http.ResponseWriter, r *http.Request, id string) {
	existing, err := a.storage.GetUser(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "user not found")
		return
	}

	var updates map[string]string
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if name, ok := updates["name"]; ok && name != "" {
		existing.Name = name
	}
	if email, ok := updates["email"]; ok && email != "" {
		existing.Email = email
	}
	if password, ok := updates["password"]; ok && password != "" {
		hashedPassword, err := storage.HashPassword(password)
		if err != nil {
			respondError(w, http.StatusInternalServerError, "failed to hash password")
			return
		}
		existing.Password = hashedPassword
	}

	if err := a.storage.UpdateUser(id, existing); err != nil {
		respondError(w, http.StatusInternalServerError, internalError(err, "update user"))
		return
	}

	respondSuccess(w, sanitizeUser(*existing))
}

func (a *API) deleteUser(w http.ResponseWriter, r *http.Request, id string) {
	_, err := a.storage.GetUser(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "user not found")
		return
	}

	if err := a.storage.DeleteUser(id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondSuccess(w, map[string]string{"message": "user deleted"})
}
