package gateway

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/velo-api/velo/internal/auth"
	"github.com/velo-api/velo/internal/cache"
	"github.com/velo-api/velo/internal/docs"
	"github.com/velo-api/velo/internal/loadbalance"
	"github.com/velo-api/velo/internal/observe"
	"github.com/velo-api/velo/internal/ratelimit"
	"github.com/velo-api/velo/internal/version"
	"github.com/velo-api/velo/pkg/config"
	"github.com/velo-api/velo/pkg/middleware"
)

type Gateway struct {
	config      *config.Config
	router      *http.ServeMux
	rateLimiter *ratelimit.RateLimiter
	cache       *cache.Cache
	auth        *auth.Auth
	lb          *loadbalance.LoadBalancer
	observer    *observe.Observer
	versioning  *version.Versioning
	docsGen     *docs.DocsGenerator
	middleware  *middleware.Chain
}

func New(cfg *config.Config) (*Gateway, error) {
	gw := &Gateway{
		config: cfg,
		router: http.NewServeMux(),
	}

	if cfg.RateLimit.Enabled {
		gw.rateLimiter = ratelimit.New(cfg.RateLimit)
	}

	if cfg.Cache.Enabled {
		gw.cache = cache.New(cfg.Cache)
	}

	if cfg.Auth.Enabled {
		gw.auth = auth.New(cfg.Auth)
	}

	if cfg.LoadBalancer.Enabled {
		gw.lb = loadbalance.New(cfg.LoadBalancer)
	}

	if cfg.Observe.Metrics.Enabled || cfg.Observe.Logs.Enabled {
		gw.observer = observe.New(cfg.Observe)
	}

	gw.versioning = version.New(cfg.Versioning)
	gw.docsGen = docs.New(cfg.Docs)

	gw.middleware = middleware.NewChain(
		gw.corsMiddleware,
		gw.loggingMiddleware,
		gw.metricsMiddleware,
		gw.rateLimitMiddleware,
		gw.authMiddleware,
		gw.cacheMiddleware,
		gw.versionMiddleware,
	)

	return gw, nil
}

func (gw *Gateway) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", gw.handleMetrics)
	mux.HandleFunc(gw.config.Docs.Path, gw.handleDocs)
	mux.HandleFunc("/", gw.handleRequest)

	handler := gw.middleware.Then(mux)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", gw.config.Server.Host, gw.config.Server.Port),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server listening on %s", server.Addr)
	return server.ListenAndServe()
}

func (gw *Gateway) handleRequest(w http.ResponseWriter, r *http.Request) {
	if gw.config.LoadBalancer.Enabled && gw.lb != nil {
		gw.lb.Forward(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Velo API Gateway", "status": "ok"}`))
}

func (gw *Gateway) handleMetrics(w http.ResponseWriter, r *http.Request) {
	if gw.observer != nil {
		gw.observer.ServeMetrics(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (gw *Gateway) handleDocs(w http.ResponseWriter, r *http.Request) {
	gw.docsGen.Serve(w, r)
}

func (gw *Gateway) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		if gw.observer != nil {
			gw.observer.LogRequest(r, time.Since(start))
		}
	})
}

func (gw *Gateway) metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if gw.observer != nil {
			gw.observer.IncRequests(r.Method, r.URL.Path)
		}
		next.ServeHTTP(w, r)
	})
}

func (gw *Gateway) rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if gw.rateLimiter != nil {
			if !gw.rateLimiter.Allow(r) {
				http.Error(w, `{"error": "rate limit exceeded"}`, http.StatusTooManyRequests)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (gw *Gateway) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if gw.auth != nil {
			if !gw.auth.Validate(r) {
				http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (gw *Gateway) cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if gw.cache != nil && r.Method == http.MethodGet {
			if cached, ok := gw.cache.Get(r.URL.String()); ok {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-Cache", "HIT")
				w.Write(cached)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (gw *Gateway) versionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if gw.versioning != nil {
			gw.versioning.Process(r)
		}
		next.ServeHTTP(w, r)
	})
}

func (gw *Gateway) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Version, X-Request-ID")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
