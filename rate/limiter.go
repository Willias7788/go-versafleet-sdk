package rate

import (
	"context"

	"golang.org/x/time/rate"
)

// Limiter defines the interface for rate limiting
type Limiter interface {
	Wait(ctx context.Context) error
}

// Params defines the rate limit parameters
type Params struct {
	RPS   float64
	Burst int
}

// New creates a new rate limiter. 
// For 100 req/min, RPS = 100/60 = 1.666...
func New(params Params) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(params.RPS), params.Burst)
}

// Default creates a limiter matching VersaFleet's 100 req/min standard
func Default() *rate.Limiter {
	// 100 requests per minute
	// Limit is events/second. 100/60.
	limit := rate.Limit(100.0 / 60.0)
	// Burst of 10 to allow some parallel operations
	return rate.NewLimiter(limit, 10)
}
