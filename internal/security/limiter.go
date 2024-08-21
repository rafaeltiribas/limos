package security

import (
	"sync"
	"time"
)

type RateLimiter struct {
	userRequestCount map[int64]int
	userLastRequest  map[int64]time.Time
	mu               sync.Mutex
	requestLimit     int
	banDuration      time.Duration
	windowDuration   time.Duration
}

func NewRateLimiter(requestLimit int, banDuration, windowDuration time.Duration) *RateLimiter {
	return &RateLimiter{
		userRequestCount: make(map[int64]int),
		userLastRequest:  make(map[int64]time.Time),
		requestLimit:     requestLimit,
		banDuration:      banDuration,
		windowDuration:   windowDuration,
	}
}

func (rl *RateLimiter) IsAllowed(chatID int64) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	count, exists := rl.userRequestCount[chatID]
	lastRequest, _ := rl.userLastRequest[chatID]

	if exists && time.Since(lastRequest) > rl.windowDuration {
		count = 0
	}

	count++

	if count > rl.requestLimit {
		if time.Since(lastRequest) <= rl.banDuration {
			return false
		}
		count = 0
	}

	rl.userRequestCount[chatID] = count
	rl.userLastRequest[chatID] = time.Now()

	return true
}
