package storage

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dgraph-io/badger/v3"
	"github.com/velo-api/velo/pkg/utils"
)

const (
	prefixUser    = "user:"
	prefixPost    = "post:"
	prefixComment = "comment:"
	prefixEmail   = "email:"
	prefixUserPost = "userpost:"
)

type Engine struct {
	db *badger.DB
}

func NewEngine(path string) (*Engine, error) {
	opts := badger.DefaultOptions(path).
		WithLogger(nil).
		WithNumVersionsToKeep(1)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to open badger: %w", err)
	}

	return &Engine{db: db}, nil
}

func (e *Engine) Close() error {
	return e.db.Close()
}

func (e *Engine) set(prefix, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return e.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(prefix+key), data)
	})
}

func (e *Engine) get(prefix, key string, dest interface{}) error {
	return e.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(prefix + key))
		if err == badger.ErrKeyNotFound {
			return fmt.Errorf("not found")
		}
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, dest)
		})
	})
}

func (e *Engine) delete(prefix, key string) error {
	return e.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(prefix + key))
	})
}

func (e *Engine) list(prefix string) ([][]byte, error) {
	var results [][]byte
	err := e.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = []byte(prefix)
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.ValidForPrefix([]byte(prefix)); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				cp := make([]byte, len(val))
				copy(cp, val)
				results = append(results, cp)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return results, err
}

func (e *Engine) setIndex(indexKey, value string) error {
	return e.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(indexKey), []byte(value))
	})
}

func (e *Engine) getIndex(indexKey string) (string, error) {
	var result string
	err := e.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(indexKey))
		if err == badger.ErrKeyNotFound {
			return fmt.Errorf("not found")
		}
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			result = string(val)
			return nil
		})
	})
	return result, err
}

func (e *Engine) count(prefix string) (int, error) {
	count := 0
	err := e.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = []byte(prefix)
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.ValidForPrefix([]byte(prefix)); it.Next() {
			count++
		}
		return nil
	})
	return count, err
}

// Users
func (e *Engine) GetUser(id string) (*User, error) {
	var user User
	err := e.get(prefixUser, id, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (e *Engine) ListUsers() ([]User, error) {
	data, err := e.list(prefixUser)
	if err != nil {
		return nil, err
	}
	var users []User
	for _, d := range data {
		var u User
		if err := json.Unmarshal(d, &u); err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

func (e *Engine) CreateUser(u *User) error {
	if err := e.set(prefixUser, u.ID, u); err != nil {
		return err
	}
	return e.setIndex(prefixEmail+u.Email, u.ID)
}

func (e *Engine) UpdateUser(id string, u *User) error {
	return e.set(prefixUser, id, u)
}

func (e *Engine) DeleteUser(id string) error {
	user, err := e.GetUser(id)
	if err != nil {
		return err
	}
	if err := e.delete(prefixUser, id); err != nil {
		return err
	}
	return e.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(prefixEmail + user.Email))
	})
}

func (e *Engine) GetUserByEmail(email string) (*User, error) {
	userID, err := e.getIndex(prefixEmail + email)
	if err != nil {
		return nil, err
	}
	return e.GetUser(userID)
}

// Posts
func (e *Engine) GetPost(id string) (*Post, error) {
	var post Post
	err := e.get(prefixPost, id, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (e *Engine) ListPosts() ([]Post, error) {
	data, err := e.list(prefixPost)
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, d := range data {
		var p Post
		if err := json.Unmarshal(d, &p); err != nil {
			continue
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (e *Engine) ListPostsByUser(userID string) ([]Post, error) {
	all, err := e.ListPosts()
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, p := range all {
		if p.UserID == userID {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (e *Engine) CreatePost(p *Post) error {
	return e.set(prefixPost, p.ID, p)
}

func (e *Engine) DeletePost(id string) error {
	return e.delete(prefixPost, id)
}

func (e *Engine) NextID(prefix string) string {
	return utils.GeneratePrefixedID(prefix[:len(prefix)-1])
}

// Comments
func (e *Engine) GetComments(postID string) ([]Comment, error) {
	data, err := e.list(prefixComment)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	for _, d := range data {
		var c Comment
		if err := json.Unmarshal(d, &c); err != nil {
			continue
		}
		if c.PostID == postID {
			comments = append(comments, c)
		}
	}
	return comments, nil
}

func (e *Engine) CreateComment(c *Comment) error {
	return e.set(prefixComment, c.ID, c)
}

// Stats
func (e *Engine) Stats() map[string]int64 {
	stats := make(map[string]int64)
	userCount, _ := e.count(prefixUser)
	postCount, _ := e.count(prefixPost)
	commentCount, _ := e.count(prefixComment)
	stats["users"] = int64(userCount)
	stats["posts"] = int64(postCount)
	stats["comments"] = int64(commentCount)
	return stats
}

// ParseID extracts ID from path like /api/v1/users/123 -> 123
func ParseID(path, prefix string) string {
	parts := strings.Split(path, "/")
	for i, p := range parts {
		if p == prefix && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
