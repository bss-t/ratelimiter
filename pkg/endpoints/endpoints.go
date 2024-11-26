package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/bss-t/ratelimiter/pkg/ratelimiter"
)

var bucket = ratelimiter.NewTokenBucket(10, 2)

func HandleConsume(w http.ResponseWriter, r *http.Request) {
	if bucket.Consume() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Token granted"))
	} else {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Rate limit exceeded"))
	}
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	bucket.Mu.Lock()
	defer bucket.Mu.Unlock()

	status := map[string]interface{}{
		"capacity":   bucket.Capacity,
		"tokens":     bucket.Tokens,
		"refillRate": bucket.RefillRate,
		"lastRefill": bucket.LastRefill,
	}
	json.NewEncoder(w).Encode(status)
}
