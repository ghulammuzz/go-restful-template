package rate

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	requestCounts map[string]int
	mutex         sync.Mutex
	requestLimit  int
	interval      time.Duration
	cleanupDelay  time.Duration
}

func NewRateLimiter(limit int, interval time.Duration, cleanupDelay time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requestCounts: make(map[string]int),
		requestLimit:  limit,
		interval:      interval,
		cleanupDelay:  cleanupDelay,
	}
	go rl.cleanup()
	return rl
}

func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(rl.cleanupDelay)
		rl.mutex.Lock()
		rl.requestCounts = make(map[string]int)
		rl.mutex.Unlock()
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		rl.mutex.Lock()
		requests, exists := rl.requestCounts[ip]
		if !exists {
			rl.requestCounts[ip] = 1
		} else if requests >= rl.requestLimit {
			rl.mutex.Unlock()
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		} else {
			rl.requestCounts[ip]++
		}
		rl.mutex.Unlock()

		next.ServeHTTP(w, r)
	})
}
