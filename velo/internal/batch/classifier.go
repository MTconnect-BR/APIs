package batch

import (
	"net/http"
	"strings"
)

type Classifier interface {
	Classify(r *http.Request) string
}

type EndpointClassifier struct{}

func (c *EndpointClassifier) Classify(r *http.Request) string {
	path := r.URL.Path
	method := r.Method

	parts := strings.Split(path, "/")
	if len(parts) >= 4 {
		return method + ":" + strings.Join(parts[:4], "/")
	}
	return method + ":" + path
}

type MethodClassifier struct{}

func (c *MethodClassifier) Classify(r *http.Request) string {
	return r.Method
}

type MethodAndPathClassifier struct{}

func (c *MethodAndPathClassifier) Classify(r *http.Request) string {
	return r.Method + ":" + r.URL.Path
}

func NewClassifier(classifierType string) Classifier {
	switch classifierType {
	case "endpoint":
		return &EndpointClassifier{}
	case "method":
		return &MethodClassifier{}
	case "methodAndPath":
		return &MethodAndPathClassifier{}
	default:
		return &EndpointClassifier{}
	}
}

func IsBatchable(r *http.Request) bool {
	return r.Method == http.MethodGet
}
