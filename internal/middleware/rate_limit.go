package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/indraplrg/technical_test/internal/response"
)

const (
	// rateLimitCleanupInterval is how often idle client entries are swept.
	rateLimitCleanupInterval = 5 * time.Minute
	// rateLimitIdleTimeout is how long a client entry survives without requests.
	rateLimitIdleTimeout = 30 * time.Minute
)

type clientEntry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimiter enforces a per-client-IP token bucket limit so a single user
// cannot spam the API. Clients are keyed by their client IP.
type RateLimiter struct {
	mu      sync.Mutex
	clients map[string]*clientEntry
	rps     float64
	burst   int
}

// NewRateLimiter creates a RateLimiter that allows rps requests per second
// per client, with a burst allowance of burst. A background goroutine evicts
// idle clients to keep memory bounded.
func NewRateLimiter(rps float64, burst int) *RateLimiter {
	if rps <= 0 {
		rps = 1
	}
	if burst < 1 {
		burst = 1
	}
	rl := &RateLimiter{
		clients: make(map[string]*clientEntry),
		rps:     rps,
		burst:   burst,
	}
	go rl.cleanupLoop()
	return rl
}

// Middleware returns the Gin handler enforcing the rate limit.
func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rl.allow(c.ClientIP()) {
			c.Header("Retry-After", "1")
			response.Error(c, http.StatusTooManyRequests, "rate limit exceeded, please try again later")
			c.Abort()
			return
		}
		c.Next()
	}
}

func (rl *RateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	entry, ok := rl.clients[ip]
	if !ok {
		entry = &clientEntry{
			limiter:  rate.NewLimiter(rate.Limit(rl.rps), rl.burst),
			lastSeen: time.Now(),
		}
		rl.clients[ip] = entry
	}
	entry.lastSeen = time.Now()
	return entry.limiter.Allow()
}

func (rl *RateLimiter) cleanupLoop() {
	ticker := time.NewTicker(rateLimitCleanupInterval)
	defer ticker.Stop()
	for range ticker.C {
		rl.cleanup()
	}
}

func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cutoff := time.Now().Add(-rateLimitIdleTimeout)
	for ip, entry := range rl.clients {
		if entry.lastSeen.Before(cutoff) {
			delete(rl.clients, ip)
		}
	}
}