package docs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/velo-api/velo/pkg/config"
)

type DocsGenerator struct {
	config config.DocsConfig
}

type OpenAPISpec struct {
	OpenAPI string                 `json:"openapi"`
	Info    map[string]string      `json:"info"`
	Paths   map[string]interface{} `json:"paths"`
}

func New(cfg config.DocsConfig) *DocsGenerator {
	return &DocsGenerator{config: cfg}
}

func (d *DocsGenerator) Serve(w http.ResponseWriter, r *http.Request) {
	spec := OpenAPISpec{
		OpenAPI: "3.1.0",
		Info: map[string]string{
			"title":   d.config.Title,
			"version": d.config.Version,
		},
		Paths: map[string]interface{}{
			"/": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Health check",
					"description": "Returns gateway status",
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Success",
						},
					},
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spec)
}

func (d *DocsGenerator) GetSpecJSON() string {
	spec := OpenAPISpec{
		OpenAPI: "3.1.0",
		Info: map[string]string{
			"title":   d.config.Title,
			"version": d.config.Version,
		},
		Paths: map[string]interface{}{},
	}
	data, _ := json.Marshal(spec)
	return fmt.Sprintf("%s", data)
}
