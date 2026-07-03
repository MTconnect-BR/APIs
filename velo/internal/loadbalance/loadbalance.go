package loadbalance

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/velo-api/velo/pkg/config"
)

type LoadBalancer struct {
	config   config.LoadBalancerConfig
	backends []*Backend
	counter  uint64
	mu       sync.RWMutex
}

type Backend struct {
	URL    *url.URL
	Weight int
	Alive  bool
	Proxy  *httputil.ReverseProxy
}

func New(cfg config.LoadBalancerConfig) *LoadBalancer {
	lb := &LoadBalancer{
		config:  cfg,
		backends: make([]*Backend, 0, len(cfg.Backends)),
	}

	for _, b := range cfg.Backends {
		u, err := url.Parse(b.URL)
		if err != nil {
			continue
		}

		proxy := httputil.NewSingleHostReverseProxy(u)
		backend := &Backend{
			URL:    u,
			Weight: b.Weight,
			Alive:  true,
			Proxy:  proxy,
		}
		lb.backends = append(lb.backends, backend)
	}

	go lb.healthCheck()

	return lb
}

func (lb *LoadBalancer) Forward(w http.ResponseWriter, r *http.Request) {
	backend := lb.nextBackend()
	if backend == nil {
		http.Error(w, `{"error": "no backends available"}`, http.StatusServiceUnavailable)
		return
	}

	backend.Proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) nextBackend() *Backend {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	aliveBackends := make([]*Backend, 0)
	for _, b := range lb.backends {
		if b.Alive {
			aliveBackends = append(aliveBackends, b)
		}
	}

	if len(aliveBackends) == 0 {
		return nil
	}

	switch lb.config.Strategy {
	case "round-robin":
		idx := atomic.AddUint64(&lb.counter, 1)
		return aliveBackends[idx%uint64(len(aliveBackends))]
	case "least-connections":
		return aliveBackends[0]
	default:
		idx := atomic.AddUint64(&lb.counter, 1)
		return aliveBackends[idx%uint64(len(aliveBackends))]
	}
}

func (lb *LoadBalancer) healthCheck() {
	for {
		lb.mu.Lock()
		for _, b := range lb.backends {
			resp, err := http.Get(b.URL.String())
			if err != nil {
				b.Alive = false
			} else {
				b.Alive = resp.StatusCode < 500
				resp.Body.Close()
			}
		}
		lb.mu.Unlock()
	}
}
