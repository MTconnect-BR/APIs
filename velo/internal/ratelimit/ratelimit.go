package ratelimit

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/velo-api/velo/pkg/config"
)

type RateLimiter struct {
	config  config.RateLimitConfig
	buckets map[string]*TokenBucket
	mu      sync.RWMutex
}

type TokenBucket struct {
	tokens     float64
	maxTokens  float64
	refillRate float64
	lastRefill time.Time
	mu         sync.Mutex
}

func New(cfg config.RateLimitConfig) *RateLimiter {
	rl := &RateLimiter{
		config:  cfg,
		buckets: make(map[string]*TokenBucket),
	}

	go rl.refillLoop()
	go rl.cleanupLoop()

	return rl
}

func (rl *RateLimiter) parseRate(rate string) (float64, float64) {
	parts := strings.SplitN(rate, "/", 2)
	if len(parts) != 2 {
		return 100, 100.0 / 60.0
	}

	tokens, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		tokens = 100
	}

	var refillRate float64
	switch parts[1] {
	case "s", "sec", "second":
		refillRate = tokens
	case "m", "min", "minute":
		refillRate = tokens / 60.0
	case "h", "hr", "hour":
		refillRate = tokens / 3600.0
	default:
		refillRate = tokens / 60.0
	}

	return tokens, refillRate
}

func (rl *RateLimiter) Allow(r *http.Request) bool {
	key := rl.getKey(r)
	bucket := rl.getBucket(key)
	return bucket.Allow()
}

func (rl *RateLimiter) getKey(r *http.Request) string {
	return r.RemoteAddr
}

func (rl *RateLimiter) getBucket(key string) *TokenBucket {
	rl.mu.RLock()
	if bucket, ok := rl.buckets[key]; ok {
		rl.mu.RUnlock()
		return bucket
	}
	rl.mu.RUnlock()

	rl.mu.Lock()
	defer rl.mu.Unlock()

	if bucket, ok := rl.buckets[key]; ok {
		return bucket
	}

	maxTokens, refillRate := rl.parseRate(rl.config.Default)

	bucket := &TokenBucket{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
	rl.buckets[key] = bucket
	return bucket
}

func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.refill()
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

func (b *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.maxTokens {
		b.tokens = b.maxTokens
	}
	b.lastRefill = now
}

func (rl *RateLimiter) refillLoop() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		rl.mu.Lock()
		for _, bucket := range rl.buckets {
			bucket.refill()
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) cleanupLoop() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, bucket := range rl.buckets {
			bucket.mu.Lock()
			if now.Sub(bucket.lastRefill) > 10*time.Minute {
				delete(rl.buckets, key)
			}
			bucket.mu.Unlock()
		}
		rl.mu.Unlock()
	}
}
