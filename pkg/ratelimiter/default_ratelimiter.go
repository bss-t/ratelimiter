package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity   int        // Max tokens the bucket can hold
	Tokens     int        // Current tokens in the bucket
	RefillRate int        // Tokens added per second
	LastRefill time.Time  // Last refill timestamp
	Mu         sync.Mutex // Mutex to protect bucket state
}

// NewTokenBucket creates a new token bucket
func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		Capacity:   capacity,
		Tokens:     capacity,
		RefillRate: refillRate,
		LastRefill: time.Now(),
	}
}

// Refill refills the bucket based on elapsed time
func (tb *TokenBucket) Refill() {
	tb.Mu.Lock()
	defer tb.Mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.LastRefill).Seconds()
	newTokens := int(elapsed) * tb.RefillRate

	if newTokens > 0 {
		tb.Tokens = min(tb.Capacity, tb.Tokens+newTokens)
		tb.LastRefill = now
	}
}

// Consume attempts to consume a token from the bucket
func (tb *TokenBucket) Consume() bool {
	tb.Mu.Lock()
	defer tb.Mu.Unlock()

	tb.Refill()

	if tb.Tokens > 0 {
		tb.Tokens--
		return true
	}

	return false
}

// Helper to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
