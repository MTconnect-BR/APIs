package version

import (
	"net/http"
	"strings"

	"github.com/velo-api/velo/pkg/config"
)

type Versioning struct {
	config config.VersioningConfig
}

func New(cfg config.VersioningConfig) *Versioning {
	return &Versioning{config: cfg}
}

func (v *Versioning) Process(r *http.Request) string {
	version := v.extractVersion(r)
	r.Header.Set("X-API-Version", version)
	return version
}

func (v *Versioning) extractVersion(r *http.Request) string {
	switch v.config.Strategy {
	case "header":
		return v.fromHeader(r)
	case "path":
		return v.fromPath(r)
	case "query":
		return v.fromQuery(r)
	default:
		return v.config.Default
	}
}

func (v *Versioning) fromHeader(r *http.Request) string {
	version := r.Header.Get(v.config.Header)
	if version == "" {
		return v.config.Default
	}
	return version
}

func (v *Versioning) fromPath(r *http.Request) string {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, "v") && len(part) > 1 {
			return part
		}
	}
	return v.config.Default
}

func (v *Versioning) fromQuery(r *http.Request) string {
	version := r.URL.Query().Get("version")
	if version == "" {
		return v.config.Default
	}
	return version
}
