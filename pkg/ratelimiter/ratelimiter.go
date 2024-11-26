package ratelimiter

type RateLimiter interface {
	Refill()
	Consume() bool
}
