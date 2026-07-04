package observe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/velo-api/velo/pkg/config"
)

type Observer struct {
	config       config.ObserveConfig
	requestCount uint64
	errorCount   uint64
	metrics      map[string]*EndpointMetrics
	mu           sync.RWMutex
}

type EndpointMetrics struct {
	RequestCount uint64
	ErrorCount   uint64
	AvgLatency   time.Duration
}

type LogEntry struct {
	Timestamp string        `json:"timestamp"`
	Method    string        `json:"method"`
	Path      string        `json:"path"`
	Latency   time.Duration `json:"latency"`
	Status    int           `json:"status"`
	RequestID string        `json:"requestId,omitempty"`
}

func New(cfg config.ObserveConfig) *Observer {
	return &Observer{
		config:  cfg,
		metrics: make(map[string]*EndpointMetrics),
	}
}

func (o *Observer) IncRequests(method, path string) {
	atomic.AddUint64(&o.requestCount, 1)

	o.mu.Lock()
	defer o.mu.Unlock()

	key := fmt.Sprintf("%s %s", method, path)
	if _, ok := o.metrics[key]; !ok {
		o.metrics[key] = &EndpointMetrics{}
	}
	o.metrics[key].RequestCount++
}

func (o *Observer) LogRequest(r *http.Request, latency time.Duration) {
	if o.config.Logs.Enabled {
		entry := LogEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Method:    r.Method,
			Path:      r.URL.Path,
			Latency:   latency,
		}
		data, _ := json.Marshal(entry)
		fmt.Println(string(data))
	}
}

func (o *Observer) ServeMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	o.mu.RLock()
	defer o.mu.RUnlock()

	fmt.Fprintf(w, "# HELP velo_requests_total Total number of requests\n")
	fmt.Fprintf(w, "# TYPE velo_requests_total counter\n")
	fmt.Fprintf(w, "velo_requests_total %d\n", atomic.LoadUint64(&o.requestCount))

	for key, m := range o.metrics {
		fmt.Fprintf(w, "velo_endpoint_requests{endpoint=\"%s\"} %d\n", key, m.RequestCount)
	}
}
