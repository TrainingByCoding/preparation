package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ========================================
// Exercise 1: Basic Cancel
// ========================================

// TODO: Implement worker that loops, doing work every 500ms
// Stop when ctx.Done() is closed
// Print "worker: tick" each iteration, "worker: stopped" when done
func worker(ctx context.Context, id int) {
	// TODO: use for loop + select on ctx.Done() and time.After/time.Ticker
}

func exercise1() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(2 * time.Second)
	fmt.Println("main: cancelling context")
	cancel()
	time.Sleep(100 * time.Millisecond) // let goroutines print "stopped"
}

// ========================================
// Exercise 2: Timeout
// ========================================

// Simulates a slow DB query (always takes 3 seconds)
func slowDBQuery(ctx context.Context) (string, error) {
	select {
	case <-time.After(3 * time.Second):
		return "query result", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func exercise2() {
	// TODO: Create a context with 1 second timeout
	// Call slowDBQuery with that context
	// Print the result or the error

	// Expected output: error — context deadline exceeded
}

// ========================================
// Exercise 3: Context Chain (3-level propagation)
// ========================================

// Simulates DB layer
func dbLayer(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second): // slow
		return nil
	case <-ctx.Done():
		fmt.Println("  [DB] cancelled:", ctx.Err())
		return ctx.Err()
	}
}

// Simulates service layer
func serviceLayer(ctx context.Context) error {
	fmt.Println("  [Service] calling DB...")
	return dbLayer(ctx) // propagates same context
}

// Simulates HTTP handler
func handlerLayer(ctx context.Context) error {
	fmt.Println("  [Handler] calling service...")
	return serviceLayer(ctx) // propagates same context
}

func exercise3() {
	// TODO: Create context with 1 second timeout
	// Call handlerLayer — cancellation should propagate through all 3 layers
}

// ========================================
// Exercise 4: context.WithValue
// ========================================

type contextKey string

const requestIDKey contextKey = "requestID"

// TODO: Implement logWithRequestID — reads requestID from context and logs
// fmt.Printf("[%s] %s\n", requestID, message)
func logWithRequestID(ctx context.Context, message string) {
	// TODO: extract requestID using ctx.Value(requestIDKey)
	// If not found, use "unknown"
}

func processRequest(ctx context.Context) {
	logWithRequestID(ctx, "Processing request")
	time.Sleep(100 * time.Millisecond)
	logWithRequestID(ctx, "Request complete")
}

func exercise4() {
	// TODO: Create context with requestID = "req-12345"
	// Use context.WithValue
	// Call processRequest with that context

	ctx := context.Background()
	// TODO: add requestID to ctx
	processRequest(ctx)
}

// ========================================
// Exercise 5: Cancel Multiple Goroutines
// ========================================

func fetchData(ctx context.Context, source string, delay time.Duration) {
	select {
	case <-time.After(delay):
		fmt.Printf("[%s] data fetched\n", source)
	case <-ctx.Done():
		fmt.Printf("[%s] cancelled: %v\n", source, ctx.Err())
	}
}

func exercise5() {
	// TODO: Create one cancellable context
	// Start 4 goroutines using it (with different delays: 500ms, 1s, 2s, 3s)
	// Cancel after 1.2 seconds
	// Which goroutines finish? Which are cancelled?
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 39: Context Cancellation ===\n")

	fmt.Println("Exercise 1: Manual Cancel (2 workers, stop after 2s)")
	exercise1()

	fmt.Println("\nExercise 2: Timeout (1s timeout on 3s DB query)")
	exercise2()

	fmt.Println("\nExercise 3: Context Chain (3-layer propagation)")
	exercise3()

	fmt.Println("\nExercise 4: WithValue (request ID)")
	exercise4()

	fmt.Println("\nExercise 5: Cancel Multiple Goroutines")
	exercise5()

	_ = errors.Is // used in readme
}

/*
SOLUTIONS:

// Exercise 1:
func worker(ctx context.Context, id int) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: stopped (%v)\n", id, ctx.Err())
			return
		case <-ticker.C:
			fmt.Printf("worker %d: tick\n", id)
		}
	}
}

// Exercise 2:
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()
result, err := slowDBQuery(ctx)
if err != nil { fmt.Println("Error:", err) } else { fmt.Println("Result:", result) }

// Exercise 3:
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()
err := handlerLayer(ctx)
fmt.Println("Handler returned:", err)

// Exercise 4:
func logWithRequestID(ctx context.Context, message string) {
	id, ok := ctx.Value(requestIDKey).(string)
	if !ok { id = "unknown" }
	fmt.Printf("[%s] %s\n", id, message)
}
ctx = context.WithValue(ctx, requestIDKey, "req-12345")

// Exercise 5:
ctx, cancel := context.WithCancel(context.Background())
go fetchData(ctx, "DB", 500*time.Millisecond)
go fetchData(ctx, "Cache", 1*time.Second)
go fetchData(ctx, "API", 2*time.Second)
go fetchData(ctx, "Queue", 3*time.Second)
time.Sleep(1200 * time.Millisecond)
cancel()
time.Sleep(100 * time.Millisecond)
*/
