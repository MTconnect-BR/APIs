package ratelimit

import (
	"net/http"
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
	tokens    float64
	maxTokens float64
	refillRate float64
	lastRefill time.Time
}

func New(cfg config.RateLimitConfig) *RateLimiter {
	rl := &RateLimiter{
		config:  cfg,
		buckets: make(map[string]*TokenBucket),
	}

	go rl.refillLoop()

	return rl
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

	bucket := &TokenBucket{
		tokens:     100,
		maxTokens:  100,
		refillRate: 100.0 / 60.0,
		lastRefill: time.Now(),
	}
	rl.buckets[key] = bucket
	return bucket
}

func (b *TokenBucket) Allow() bool {
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
