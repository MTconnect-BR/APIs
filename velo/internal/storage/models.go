package storage

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"createdAt"`
}

type Post struct {
	ID        string   `json:"id"`
	UserID    string   `json:"userId"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags,omitempty"`
	CreatedAt int64    `json:"createdAt"`
}

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"postId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"createdAt"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func NewUser(id, name, email, password string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Unix(),
	}
}

func NewPost(id, userID, title, body string, tags []string) *Post {
	return &Post{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Body:      body,
		Tags:      tags,
		CreatedAt: time.Now().Unix(),
	}
}

func NewComment(id, postID, name, email, body string) *Comment {
	return &Comment{
		ID:        id,
		PostID:    postID,
		Name:      name,
		Email:     email,
		Body:      body,
		CreatedAt: time.Now().Unix(),
	}
}
