package router

import (
	"net/http"
	"sync"
)

type Router struct {
	routes map[string]http.HandlerFunc
	mu     sync.RWMutex
}

func New() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Handle(pattern string, handler http.HandlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.routes[pattern] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mu.RLock()
	handler, ok := r.routes[req.URL.Path]
	r.mu.RUnlock()

	if ok {
		handler(w, req)
		return
	}

	http.NotFound(w, req)
}
