package main

import (
	"fmt"
	"sync"
	"time"
)

// ========================================
// Exercise 1: Basic Rate Limiter (1 per 500ms)
// ========================================

func exercise1() {
	requests := []int{1, 2, 3, 4, 5}

	// TODO: Create a time.Ticker that fires every 500ms
	// For each request, wait for the ticker before processing
	// Print: fmt.Printf("Request %d processed at %v\n", req, time.Now().Format("15:04:05.000"))

	_ = requests
}

// ========================================
// Exercise 2: Bursty Rate Limiter
// ========================================

func exercise2() {
	requests := []int{1, 2, 3, 4, 5, 6, 7}

	// TODO: Create a bursty limiter:
	// - buffer size 3 (burst capacity)
	// - pre-fill with 3 tokens
	// - ticker adds new token every 1 second

	// TODO: for each request, consume a token then process
	// Print: fmt.Printf("Request %d: processed at %s\n", req, time.Now().Format("15:04:05"))

	_ = requests
}

// ========================================
// Exercise 3: Per-User Rate Limiter
// ========================================

type UserRateLimiter struct {
	mu       sync.Mutex
	limiters map[string]*time.Ticker
	rate     time.Duration
}

func NewUserRateLimiter(rate time.Duration) *UserRateLimiter {
	return &UserRateLimiter{
		limiters: make(map[string]*time.Ticker),
		rate:     rate,
	}
}

// TODO: Implement Allow(userID string) — blocks until the user's rate allows
// If user has no limiter yet, create one
func (u *UserRateLimiter) Allow(userID string) {
	u.mu.Lock()
	// TODO: check if userID has a ticker; if not, create one
	ticker, ok := u.limiters[userID]
	if !ok {
		// TODO: create and store ticker
		_ = ticker
	}
	u.mu.Unlock()

	// TODO: wait for tick
}

func exercise3() {
	limiter := NewUserRateLimiter(500 * time.Millisecond)

	var wg sync.WaitGroup
	users := []string{"alice", "bob", "alice", "charlie", "bob", "alice"}

	for i, user := range users {
		wg.Add(1)
		go func(id int, u string) {
			defer wg.Done()
			limiter.Allow(u)
			fmt.Printf("User %-10s → Request %d at %s\n", u, id, time.Now().Format("15:04:05.000"))
		}(i+1, user)
	}
	wg.Wait()
}

// ========================================
// Exercise 4: Sliding Window Counter
// ========================================

type SlidingWindowLimiter struct {
	mu       sync.Mutex
	window   time.Duration
	maxCalls int
	calls    []time.Time
}

func NewSlidingWindowLimiter(window time.Duration, maxCalls int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		window:   window,
		maxCalls: maxCalls,
	}
}

// TODO: Implement Allow() bool
// Return true if the call is allowed (within maxCalls in window)
// Return false if limit exceeded
// Hint: remove entries older than window, then check len(calls)
func (s *SlidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	// TODO: filter out calls older than s.window
	// TODO: if len(calls) < maxCalls: append now, return true
	// else return false
	_ = now
	return false
}

func exercise4() {
	limiter := NewSlidingWindowLimiter(2*time.Second, 3)

	for i := 1; i <= 7; i++ {
		allowed := limiter.Allow()
		if allowed {
			fmt.Printf("Call %d: ALLOWED at %s\n", i, time.Now().Format("15:04:05"))
		} else {
			fmt.Printf("Call %d: REJECTED (rate limit exceeded)\n", i)
		}
		time.Sleep(400 * time.Millisecond)
	}
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 37: Rate Limiter Pattern ===\n")

	fmt.Println("Exercise 1: Basic Rate Limiter (1 per 500ms)")
	exercise1()

	fmt.Println("\nExercise 2: Bursty Rate Limiter (burst 3, then 1/sec)")
	exercise2()

	fmt.Println("\nExercise 3: Per-User Rate Limiter")
	exercise3()

	fmt.Println("\nExercise 4: Sliding Window (3 per 2s)")
	exercise4()
}

/*
SOLUTIONS:

// Exercise 1:
limiter := time.NewTicker(500 * time.Millisecond)
defer limiter.Stop()
for _, req := range requests {
	<-limiter.C
	fmt.Printf("Request %d processed at %v\n", req, time.Now().Format("15:04:05.000"))
}

// Exercise 2:
limiter := time.NewTicker(1 * time.Second)
burstyLimiter := make(chan time.Time, 3)
for i := 0; i < 3; i++ { burstyLimiter <- time.Now() }
go func() { for t := range limiter.C { burstyLimiter <- t } }()
for _, req := range requests {
	<-burstyLimiter
	fmt.Printf("Request %d: processed at %s\n", req, time.Now().Format("15:04:05"))
}

// Exercise 3 — Allow:
func (u *UserRateLimiter) Allow(userID string) {
	u.mu.Lock()
	ticker, ok := u.limiters[userID]
	if !ok {
		ticker = time.NewTicker(u.rate)
		u.limiters[userID] = ticker
	}
	u.mu.Unlock()
	<-ticker.C
}

// Exercise 4 — Allow:
func (s *SlidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	cutoff := now.Add(-s.window)
	filtered := s.calls[:0]
	for _, t := range s.calls {
		if t.After(cutoff) { filtered = append(filtered, t) }
	}
	s.calls = filtered
	if len(s.calls) < s.maxCalls {
		s.calls = append(s.calls, now)
		return true
	}
	return false
}
*/
