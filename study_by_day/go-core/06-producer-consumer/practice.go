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
// Bonus 2: HTTP-backed Producer-Consumer (from master100)
// ========================================
// Pattern: HTTP handler enqueues to a buffered channel, background goroutine consumes
//   var queue = make(chan string, 10)
//   go consume()        // consumer runs in background
//   http.HandleFunc("/enqueue", handler)  // handler is the producer
//   http.ListenAndServe(":8080", nil)
//
// handler: reads query param "data", sends to queue channel
// consume: loops over queue channel and processes each item

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

	fmt.Println("\n=== Bonus: Rate Limited ===")
	ch := make(chan int)
	go rateLimitedProducer(ch, 500*time.Millisecond)

	// Receive first 5 numbers
	for i := 0; i < 5; i++ {
		fmt.Printf("Received: %d (at %v)\n", <-ch, time.Now().Format("15:04:05"))
	}
}

/*
SOLUTIONS:

func exercise1() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ { ch <- i }
		close(ch)
	}()
	for n := range ch { fmt.Println(n * n) }
}

func exercise2() {
	ch := make(chan int, 30)
	var wg sync.WaitGroup
	for p := 0; p < 3; p++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for i := start; i < start+10; i++ { ch <- i }
		}(p*10 + 1)
	}
	go func() { wg.Wait(); close(ch) }()
	sum := 0
	for n := range ch { sum += n }
	fmt.Println("Total sum:", sum) // 1+2+...+30 = 465
}

func exercise3() {
	jobs := make(chan int, 100)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs { fmt.Printf("Worker %d processed job %d\n", id, j) }
		}(w)
	}
	go func() {
		for i := 1; i <= 30; i++ { jobs <- i }
		close(jobs)
	}()
	wg.Wait()
}

func rateLimitedProducer(ch chan<- int, rate time.Duration) {
	ticker := time.NewTicker(rate)
	defer ticker.Stop()
	i := 1
	for range ticker.C { ch <- i; i++ }
}
*/
