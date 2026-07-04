package gateway

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/velo-api/velo/internal/api"
	"github.com/velo-api/velo/internal/auth"
	"github.com/velo-api/velo/internal/cache"
	"github.com/velo-api/velo/internal/docs"
	"github.com/velo-api/velo/internal/loadbalance"
	"github.com/velo-api/velo/internal/observe"
	"github.com/velo-api/velo/internal/ratelimit"
	"github.com/velo-api/velo/internal/storage"
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
	storage     *storage.Engine
	apiHandlers *api.API
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

	if cfg.Storage.Enabled {
		storePath := cfg.Storage.Path
		if storePath == "" {
			storePath = "./data/velo.db"
		}

		log.Printf("📦 Opening storage at %s...", storePath)
		store, err := storage.NewEngine(storePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open storage: %w", err)
		}
		gw.storage = store

		if cfg.Storage.Seed {
			storage.Seed(store)
		}

		gw.apiHandlers = api.New(store)
		log.Println("✅ Storage engine ready")
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
	mux.HandleFunc("/health", gw.handleHealth)

	if gw.apiHandlers != nil {
		mux.HandleFunc("/api/v1/users", gw.handleAPIRequest)
		mux.HandleFunc("/api/v1/users/", gw.handleAPIRequest)
		mux.HandleFunc("/api/v1/posts", gw.handleAPIRequest)
		mux.HandleFunc("/api/v1/posts/", gw.handleAPIRequest)
		mux.HandleFunc("/api/v1/comments", gw.handleAPIRequest)
		mux.HandleFunc("/api/v1/comments/", gw.handleAPIRequest)
		mux.HandleFunc("/api/v1/auth/login", gw.handleAPIRequest)
		mux.HandleFunc("/api/v1/auth/register", gw.handleAPIRequest)
	}

	mux.HandleFunc("/", gw.handleRequest)

	handler := gw.middleware.Then(mux)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", gw.config.Server.Host, gw.config.Server.Port),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("🚀 Velo API Gateway running on %s", server.Addr)
	log.Printf("📊 API Endpoints:")
	log.Printf("   GET    /api/v1/users")
	log.Printf("   GET    /api/v1/users/:id")
	log.Printf("   POST   /api/v1/users")
	log.Printf("   PUT    /api/v1/users/:id")
	log.Printf("   DELETE /api/v1/users/:id")
	log.Printf("   GET    /api/v1/posts")
	log.Printf("   GET    /api/v1/posts/:id")
	log.Printf("   POST   /api/v1/posts")
	log.Printf("   DELETE /api/v1/posts/:id")
	log.Printf("   GET    /api/v1/posts/:id/comments")
	log.Printf("   POST   /api/v1/comments")
	log.Printf("   POST   /api/v1/auth/login")
	log.Printf("   POST   /api/v1/auth/register")
	log.Printf("   GET    /health")
	log.Printf("   GET    /metrics")
	log.Printf("   GET    %s", gw.config.Docs.Path)

	return server.ListenAndServe()
}

func (gw *Gateway) handleHealth(w http.ResponseWriter, r *http.Request) {
	stats := map[string]interface{}{
		"status": "ok",
		"time":   time.Now().Unix(),
	}

	if gw.storage != nil {
		stats["storage"] = gw.storage.Stats()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","time":%d}`, time.Now().Unix())
}

func (gw *Gateway) handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if strings.HasPrefix(path, "/api/v1/auth/") {
		switch {
		case strings.HasSuffix(path, "/login"):
			gw.apiHandlers.Login(w, r)
		case strings.HasSuffix(path, "/register"):
			gw.apiHandlers.Register(w, r)
		default:
			http.NotFound(w, r)
		}
		return
	}

	if strings.HasPrefix(path, "/api/v1/users") {
		if path == "/api/v1/users" {
			gw.apiHandlers.Users(w, r)
		} else {
			gw.apiHandlers.UserByID(w, r)
		}
		return
	}

	if strings.HasPrefix(path, "/api/v1/posts") {
		if path == "/api/v1/posts" {
			gw.apiHandlers.Posts(w, r)
		} else {
			gw.apiHandlers.PostsByID(w, r)
		}
		return
	}

	if strings.HasPrefix(path, "/api/v1/comments") {
		gw.apiHandlers.Comments(w, r)
		return
	}

	http.NotFound(w, r)
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

type cachedResponseWriter struct {
	http.ResponseWriter
	body    []byte
	status  int
	isCache bool
}

func (crw *cachedResponseWriter) Write(b []byte) (int, error) {
	if crw.isCache {
		return len(b), nil
	}
	crw.body = append(crw.body, b...)
	return crw.ResponseWriter.Write(b)
}

func (crw *cachedResponseWriter) WriteHeader(code int) {
	crw.status = code
	if !crw.isCache {
		crw.ResponseWriter.WriteHeader(code)
	}
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

			crw := &cachedResponseWriter{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(crw, r)

			if crw.status == http.StatusOK && len(crw.body) > 0 {
				gw.cache.Set(r.URL.String(), crw.body, 5*time.Minute)
			}
			return
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
