package batch

import (
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/velo-api/velo/pkg/config"
)

type BatchEngine struct {
	config     *BatchConfig
	classifier Classifier
	processor  Processor
	boxes      map[string]*RequestBox
	mu         sync.Mutex
	counter    uint64
	running    bool
	stopCh     chan struct{}
}

type BatchMetrics struct {
	BoxesTotal    uint64
	RequestsTotal uint64
	AvgBoxSize    float64
}

var metrics BatchMetrics

func NewBatchEngine(cfg config.BatchConfig, processor Processor) *BatchEngine {
	batchCfg := ParseConfig(cfg)

	return &BatchEngine{
		config:     batchCfg,
		classifier: NewClassifier("endpoint"),
		processor:  processor,
		boxes:      make(map[string]*RequestBox),
		stopCh:     make(chan struct{}),
	}
}

func (e *BatchEngine) Start() {
	if !e.config.Enabled {
		return
	}

	e.running = true
	go e.flushLoop()
	go e.cleanupLoop()

	log.Printf("Batch engine started (window: %v, maxBatchSize: %d)", e.config.Window, e.config.MaxBatchSize)
}

func (e *BatchEngine) Stop() {
	if !e.running {
		return
	}

	e.running = false
	close(e.stopCh)

	e.mu.Lock()
	defer e.mu.Unlock()

	for key, box := range e.boxes {
		e.processBox(box)
		delete(e.boxes, key)
	}
}

func (e *BatchEngine) Submit(r *http.Request) <-chan *BatchItemResult {
	if !e.config.Enabled || !IsBatchable(r) {
		resultCh := make(chan *BatchItemResult, 1)
		go func() {
			rec := &responseRecorder{}
			e.processSingle(r, rec)
			resultCh <- &BatchItemResult{
				StatusCode: rec.status,
				Body:       rec.body,
				Headers:    rec.headers,
			}
		}()
		return resultCh
	}

	key := e.classifier.Classify(r)
	resultCh := make(chan *BatchItemResult, 1)

	item := &BatchItem{
		Request:  r,
		Response: resultCh,
	}

	e.mu.Lock()
	box, exists := e.boxes[key]
	if !exists {
		box = &RequestBox{
			ID:  e.generateBoxID(),
			Key: key,
		}
		e.boxes[key] = box
	}

	box.Requests = append(box.Requests, item)
	shouldProcess := len(box.Requests) >= e.config.MaxBatchSize
	e.mu.Unlock()

	atomic.AddUint64(&metrics.RequestsTotal, 1)

	if shouldProcess {
		e.processBoxByKey(key)
	}

	return resultCh
}

func (e *BatchEngine) processBoxByKey(key string) {
	e.mu.Lock()
	box, exists := e.boxes[key]
	if !exists {
		e.mu.Unlock()
		return
	}
	delete(e.boxes, key)
	e.mu.Unlock()

	e.processBox(box)
}

func (e *BatchEngine) processBox(box *RequestBox) {
	if len(box.Requests) == 0 {
		return
	}

	start := time.Now()

	results, err := e.processor.ProcessBox(box)
	if err != nil {
		log.Printf("Batch processing error for box %s: %v", box.ID, err)
		for _, item := range box.Requests {
			item.Response <- &BatchItemResult{
				StatusCode: http.StatusInternalServerError,
				Body:       []byte(`{"error": "batch processing failed"}`),
			}
		}
		return
	}

	SplitResults(box, results)

	atomic.AddUint64(&metrics.BoxesTotal, 1)

	elapsed := time.Since(start)
	avgLatency := float64(elapsed.Milliseconds()) / float64(len(box.Requests))
	log.Printf("Batch box %s processed: %d requests in %v (avg: %.2fms per request)",
		box.ID, len(box.Requests), elapsed, avgLatency)
}

func (e *BatchEngine) flushLoop() {
	ticker := time.NewTicker(e.config.FlushInterval)
	defer ticker.Stop()

	for {
		select {
		case <-e.stopCh:
			return
		case <-ticker.C:
			e.flushExpiredBoxes()
		}
	}
}

func (e *BatchEngine) flushExpiredBoxes() {
	e.mu.Lock()
	var expiredKeys []string
	now := time.Now()

	for key, box := range e.boxes {
		if len(box.Requests) > 0 {
			expiredKeys = append(expiredKeys, key)
		}
		_ = now
	}

	for _, key := range expiredKeys {
		box := e.boxes[key]
		delete(e.boxes, key)
		go e.processBox(box)
	}

	e.mu.Unlock()
}

func (e *BatchEngine) cleanupLoop() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-e.stopCh:
			return
		case <-ticker.C:
			e.cleanup()
		}
	}
}

func (e *BatchEngine) cleanup() {
	e.mu.Lock()
	defer e.mu.Unlock()

	for key, box := range e.boxes {
		if len(box.Requests) == 0 {
			delete(e.boxes, key)
		}
	}
}

func (e *BatchEngine) generateBoxID() string {
	return "box_" + string(rune(atomic.AddUint64(&e.counter, 1)))
}

func (e *BatchEngine) GetMetrics() BatchMetrics {
	return BatchMetrics{
		BoxesTotal:    atomic.LoadUint64(&metrics.BoxesTotal),
		RequestsTotal: atomic.LoadUint64(&metrics.RequestsTotal),
	}
}

type responseRecorder struct {
	status  int
	body    []byte
	headers http.Header
}

func (r *responseRecorder) Header() http.Header {
	if r.headers == nil {
		r.headers = make(http.Header)
	}
	return r.headers
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body = append(r.body, b...)
	return len(b), nil
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
}

func (e *BatchEngine) processSingle(r *http.Request, w http.ResponseWriter) {
	e.processor.ProcessBox(&RequestBox{
		Requests: []*BatchItem{
			{Request: r},
		},
	})
}
