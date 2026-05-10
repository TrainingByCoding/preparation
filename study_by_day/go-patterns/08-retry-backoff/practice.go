package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// ========================================
// Exercise 1: Fixed Retry
// ========================================

// TODO: Implement Retry — calls fn up to maxRetries times with fixedDelay between attempts
// Print: fmt.Printf("Attempt %d failed: %v\n", attempt+1, err)
// Return nil on success, last error if all retries exhausted
func Retry(maxRetries int, fixedDelay time.Duration, fn func() error) error {
	// TODO: implement
	return nil
}

// Simulated flaky service — succeeds on the Nth call
func flakyService(successOnAttempt int) func() func() error {
	count := 0
	return func() func() error {
		return func() error {
			count++
			if count < successOnAttempt {
				return fmt.Errorf("service unavailable (attempt %d)", count)
			}
			return nil
		}
	}
}

func exercise1() {
	factory := flakyService(3) // will succeed on 3rd attempt
	err := Retry(5, 200*time.Millisecond, factory())
	if err != nil {
		fmt.Println("All retries failed:", err)
	} else {
		fmt.Println("Success!")
	}
}

// ========================================
// Exercise 2: Exponential Backoff
// ========================================

// TODO: Implement RetryWithBackoff
// Start with initialDelay, double it each attempt (exponential)
// Cap at maxDelay if provided (use 0 to mean no cap)
// Print the delay before each retry
func RetryWithBackoff(maxRetries int, initialDelay, maxDelay time.Duration, fn func() error) error {
	delay := initialDelay
	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}
		if attempt == maxRetries {
			return fmt.Errorf("failed after %d attempts: %w", maxRetries+1, err)
		}
		// TODO: cap delay if maxDelay > 0
		fmt.Printf("  Attempt %d failed. Waiting %v before retry...\n", attempt+1, delay)
		time.Sleep(delay)
		delay *= 2 // exponential
	}
	return nil
}

func exercise2() {
	factory := flakyService(4)
	err := RetryWithBackoff(5, 100*time.Millisecond, 0, factory())
	if err != nil {
		fmt.Println("All retries failed:", err)
	} else {
		fmt.Println("Success!")
	}
}

// ========================================
// Exercise 3: With Jitter
// ========================================

// TODO: Implement RetryWithJitter
// Same as RetryWithBackoff but add random jitter of ±20% of current delay
// jitter := delay * (0.8 + 0.4*rand.Float64())   ← this gives ±20%
func RetryWithJitter(maxRetries int, initialDelay time.Duration, fn func() error) error {
	delay := initialDelay
	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}
		if attempt == maxRetries {
			return fmt.Errorf("failed after %d attempts: %w", maxRetries+1, err)
		}
		// TODO: apply jitter to delay
		jitteredDelay := delay // replace this with jitter calculation
		fmt.Printf("  Attempt %d failed. Waiting %v (with jitter)\n", attempt+1, jitteredDelay)
		time.Sleep(jitteredDelay)
		delay *= 2
	}
	_ = rand.Float64 // hint
	return nil
}

func exercise3() {
	factory := flakyService(3)
	err := RetryWithJitter(5, 200*time.Millisecond, factory())
	if err != nil {
		fmt.Println("All retries failed:", err)
	} else {
		fmt.Println("Success!")
	}
}

// ========================================
// Exercise 4: Context-Aware Retry
// ========================================

// TODO: Implement RetryWithContext
// Same as RetryWithBackoff but check context between retries
// If ctx is cancelled or deadline exceeded, return ctx.Err()
func RetryWithContext(ctx context.Context, maxRetries int, initialDelay time.Duration, fn func() error) error {
	delay := initialDelay
	for attempt := 0; attempt <= maxRetries; attempt++ {
		// TODO: check ctx.Err() first — if cancelled, return immediately

		err := fn()
		if err == nil {
			return nil
		}
		if attempt == maxRetries {
			return fmt.Errorf("failed after %d attempts: %w", maxRetries+1, err)
		}

		// TODO: use select with time.After(delay) and ctx.Done()
		// If context cancelled during sleep, return ctx.Err()
		delay *= 2
	}
	return nil
}

func exercise4() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	alwaysFails := func() error {
		return errors.New("always fails")
	}

	err := RetryWithContext(ctx, 10, 200*time.Millisecond, alwaysFails)
	fmt.Println("Result:", err)
	// Expected: context deadline exceeded (stops after ~500ms, not 10 retries)
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 38: Retry with Backoff ===\n")

	fmt.Println("Exercise 1: Fixed Retry")
	exercise1()

	fmt.Println("\nExercise 2: Exponential Backoff")
	exercise2()

	fmt.Println("\nExercise 3: With Jitter")
	exercise3()

	fmt.Println("\nExercise 4: Context-Aware Retry")
	exercise4()
}

/*
SOLUTIONS:

func Retry(maxRetries int, fixedDelay time.Duration, fn func() error) error {
	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := fn()
		if err == nil { return nil }
		if attempt == maxRetries {
			return fmt.Errorf("failed after %d attempts: %w", maxRetries+1, err)
		}
		fmt.Printf("Attempt %d failed: %v\n", attempt+1, err)
		time.Sleep(fixedDelay)
	}
	return nil
}

// Jitter:
jitterFactor := 0.8 + 0.4*rand.Float64()
jitteredDelay := time.Duration(float64(delay) * jitterFactor)

// Context retry:
func RetryWithContext(ctx context.Context, maxRetries int, initialDelay time.Duration, fn func() error) error {
	delay := initialDelay
	for attempt := 0; attempt <= maxRetries; attempt++ {
		if ctx.Err() != nil { return ctx.Err() }
		err := fn()
		if err == nil { return nil }
		if attempt == maxRetries {
			return fmt.Errorf("failed after %d attempts: %w", maxRetries+1, err)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
		delay *= 2
	}
	return nil
}
*/
