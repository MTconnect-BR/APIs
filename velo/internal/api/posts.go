package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/velo-api/velo/internal/storage"
)

func (a *API) Posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.listPosts(w, r)
	case http.MethodPost:
		a.createPost(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (a *API) PostsByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/v1/"), "/")

	if len(parts) >= 2 && parts[0] == "posts" {
		id := parts[1]

		if len(parts) >= 3 && parts[2] == "comments" {
			a.getCommentsByPost(w, r, id)
			return
		}

		switch r.Method {
		case http.MethodGet:
			a.getPost(w, r, id)
		case http.MethodDelete:
			a.deletePost(w, r, id)
		default:
			respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
		return
	}

	respondError(w, http.StatusBadRequest, "invalid path")
}

func (a *API) listPosts(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")

	var posts []storage.Post
	var err error

	if userID != "" {
		posts, err = a.storage.ListPostsByUser(userID)
	} else {
		posts, err = a.storage.ListPosts()
	}

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if posts == nil {
		posts = []storage.Post{}
	}
	respondSuccess(w, posts)
}

func (a *API) getPost(w http.ResponseWriter, r *http.Request, id string) {
	post, err := a.storage.GetPost(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "post not found")
		return
	}
	respondSuccess(w, post)
}

func (a *API) createPost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string   `json:"userId"`
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		Tags   []string `json:"tags"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.Title == "" || req.Body == "" {
		respondError(w, http.StatusBadRequest, "userId, title and body are required")
		return
	}

	_, err := a.storage.GetUser(req.UserID)
	if err != nil {
		respondError(w, http.StatusBadRequest, "user not found")
		return
	}

	id := a.storage.NextID("post:")
	post := &storage.Post{
		ID:     id,
		UserID: req.UserID,
		Title:  req.Title,
		Body:   req.Body,
		Tags:   req.Tags,
	}

	if err := a.storage.CreatePost(post); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondCreated(w, post)
}

func (a *API) deletePost(w http.ResponseWriter, r *http.Request, id string) {
	_, err := a.storage.GetPost(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "post not found")
		return
	}

	if err := a.storage.DeletePost(id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondSuccess(w, map[string]string{"message": "post deleted"})
}

func (a *API) getCommentsByPost(w http.ResponseWriter, r *http.Request, postID string) {
	comments, err := a.storage.GetComments(postID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if comments == nil {
		comments = []storage.Comment{}
	}
	respondSuccess(w, comments)
}
