package main

import (
	"fmt"
	"sync"
	"time"
)

// ========================================
// Exercise 1: Single Producer, Single Consumer
// ========================================
func exercise1() {
	ch := make(chan int)

	// TODO: Producer generates 1-100
	go func() {
		// Your code here
	}()

	// TODO: Consumer prints squares
	// Your code here
}

// ========================================
// Exercise 2: Multiple Producers, Single Consumer
// ========================================
func exercise2() {
	ch := make(chan int, 30) // Buffer for all data
	var wg sync.WaitGroup

	// TODO: Start 3 producers
	// Each generates numbers (producer 1: 1-10, producer 2: 11-20, etc.)

	// TODO: Close channel after all producers done

	// TODO: Consumer sums all numbers
	sum := 0
	for num := range ch {
		sum += num
	}
	fmt.Println("Total sum:", sum)
}

// ========================================
// Exercise 3: Single Producer, Multiple Consumers
// ========================================
func exercise3() {
	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	// TODO: Start 3 consumers
	// Each consumer processes jobs from channel

	// TODO: Producer generates 30 jobs
	go func() {
		// Your code here
	}()

	wg.Wait()
}

// ========================================
// Exercise 4: Pipeline Pattern
// ========================================

// Stage 1: Generate numbers 1-20
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		// TODO: Send all nums to channel
		close(out)
	}()
	return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// TODO: Read from in, square, send to out
		close(out)
	}()
	return out
}

// Stage 3: Filter even numbers only
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// TODO: Read from in, only send evens to out
		close(out)
	}()
	return out
}

func exercise4() {
	// TODO: Connect the pipeline
	// generate → square → filterEven → print

	// Example:
	// nums := []int{1, 2, 3, 4, 5}
	// for result := range filterEven(square(generate(nums...))) {
	//     fmt.Println(result)
	// }
	// Expected: 4, 16 (only even squares)
}

// ========================================
// Bonus: Rate-Limited Producer
// ========================================
func rateLimitedProducer(ch chan<- int, rate time.Duration) {
	// TODO: Send numbers but with a rate limit
	// Send one number every 'rate' duration
	ticker := time.NewTicker(rate)
	defer ticker.Stop()

	// Your code here
}

// ========================================
// Main
// ========================================
func main() {
	fmt.Println("=== Exercise 1: Single Producer/Consumer ===")
	exercise1()

	fmt.Println("\n=== Exercise 2: Multiple Producers ===")
	exercise2()

	fmt.Println("\n=== Exercise 3: Multiple Consumers ===")
	exercise3()

	fmt.Println("\n=== Exercise 4: Pipeline ===")
	exercise4()

	fmt.Println("\n=== Bonus: Rate Limited ===")
	ch := make(chan int)
	go rateLimitedProducer(ch, 500*time.Millisecond)

	// Receive first 5 numbers
	for i := 0; i < 5; i++ {
		fmt.Printf("Received: %d (at %v)\n", <-ch, time.Now().Format("15:04:05"))
	}
}
