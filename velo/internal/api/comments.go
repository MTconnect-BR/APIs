package api

import (
	"encoding/json"
	"net/http"

	"github.com/velo-api/velo/internal/storage"
)

func (a *API) Comments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		a.createComment(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (a *API) createComment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PostID string `json:"postId"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.PostID == "" || req.Name == "" || req.Email == "" || req.Body == "" {
		respondError(w, http.StatusBadRequest, "postId, name, email and body are required")
		return
	}

	_, err := a.storage.GetPost(req.PostID)
	if err != nil {
		respondError(w, http.StatusBadRequest, "post not found")
		return
	}

	id := a.storage.NextID("comment:")
	comment := &storage.Comment{
		ID:     id,
		PostID: req.PostID,
		Name:   req.Name,
		Email:  req.Email,
		Body:   req.Body,
	}

	if err := a.storage.CreateComment(comment); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondCreated(w, comment)
}
