package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ========================================
// Exercise 1 & 2: Circuit Breaker
// ========================================

type State int

const (
	StateClosed   State = iota // normal — requests pass through
	StateOpen                  // failing — reject immediately
	StateHalfOpen              // recovering — allow one test request
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF-OPEN"
	default:
		return "UNKNOWN"
	}
}

type CircuitBreaker struct {
	mu          sync.Mutex
	state       State
	failures    int
	maxFailures int
	openedAt    time.Time
	timeout     time.Duration

	// Exercise 3: metrics
	TotalCalls    int
	FailedCalls   int
	RejectedCalls int
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       StateClosed,
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

// TODO: Implement Call(fn func() error) error
// Logic:
//  1. Lock, check if Open → if timeout passed, switch to HalfOpen
//  2. If still Open → increment RejectedCalls, return error "circuit open"
//  3. Unlock, call fn()
//  4. Lock, update state based on result:
//     - success: reset failures, go to Closed
//     - failure in HalfOpen: go to Open
//     - failure in Closed: increment failures, if >= maxFailures go to Open
//  5. Return fn's error
func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	cb.TotalCalls++

	// TODO: Check if Open and if timeout has passed → transition to HalfOpen

	if cb.state == StateOpen {
		cb.RejectedCalls++
		cb.mu.Unlock()
		return errors.New("circuit breaker is OPEN — request rejected")
	}
	cb.mu.Unlock()

	// TODO: call fn()
	err := fn()
	_ = err

	cb.mu.Lock()
	// TODO: update state based on err
	cb.mu.Unlock()

	return err
}

func (cb *CircuitBreaker) Status() string {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return fmt.Sprintf("State: %s | Failures: %d | Total: %d | Rejected: %d",
		cb.state, cb.failures, cb.TotalCalls, cb.RejectedCalls)
}

// ========================================
// Simulated external service
// ========================================

type MockService struct {
	callCount int
	failUntil int // fail for first N calls
}

func (s *MockService) Call() error {
	s.callCount++
	if s.callCount <= s.failUntil {
		return fmt.Errorf("service error on call #%d", s.callCount)
	}
	return nil
}

// ========================================
// Exercise 2: Test state transitions
// ========================================

func exercise2() {
	cb := NewCircuitBreaker(3, 2*time.Second)
	svc := &MockService{failUntil: 5} // fail first 5 calls

	for i := 1; i <= 10; i++ {
		err := cb.Call(svc.Call)
		if err != nil {
			fmt.Printf("Call %d: ERROR — %v\n", i, err)
		} else {
			fmt.Printf("Call %d: SUCCESS\n", i)
		}
		fmt.Printf("  Status: %s\n", cb.Status())

		// After call 4, wait for circuit breaker timeout
		if i == 4 {
			fmt.Println("  ... waiting 2.5s for circuit to try half-open ...")
			time.Sleep(2500 * time.Millisecond)
		}
	}
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 36: Circuit Breaker Pattern ===\n")
	exercise2()
}

/*
SOLUTIONS:

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	cb.TotalCalls++

	// Check if open and timeout passed → half-open
	if cb.state == StateOpen && time.Since(cb.openedAt) >= cb.timeout {
		cb.state = StateHalfOpen
		fmt.Println("  [CB] → HALF-OPEN (testing recovery)")
	}

	if cb.state == StateOpen {
		cb.RejectedCalls++
		cb.mu.Unlock()
		return errors.New("circuit breaker is OPEN — request rejected")
	}
	cb.mu.Unlock()

	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.FailedCalls++
		if cb.state == StateHalfOpen {
			cb.state = StateOpen
			cb.openedAt = time.Now()
			fmt.Println("  [CB] → OPEN (half-open test failed)")
		} else {
			cb.failures++
			if cb.failures >= cb.maxFailures {
				cb.state = StateOpen
				cb.openedAt = time.Now()
				fmt.Println("  [CB] → OPEN (failure threshold reached)")
			}
		}
	} else {
		if cb.state == StateHalfOpen {
			fmt.Println("  [CB] → CLOSED (recovered!)")
		}
		cb.failures = 0
		cb.state = StateClosed
	}

	return err
}
*/
