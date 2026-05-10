package main

import (
	"fmt"
	"time"
)

// ========================================
// Exercise 1: Simple Channel
// ========================================
func simpleChannel() {
	// TODO: Create channel, send value in goroutine, receive in main
	ch := make(chan string)

	// TODO: Launch goroutine that sends a message

	// TODO: Receive and print the message
}

// ========================================
// Exercise 2: Buffered Channel
// ========================================
func bufferedChannel() {
	// TODO: Create buffered channel with capacity 3

	// TODO: Send 3 values (shouldn't block)

	// TODO: Receive and print all 3 values
}

// ========================================
// Exercise 3: Fix the Channel Demo
// ========================================
func fixedChannelDemo() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel1 <- "Every 100ms"
			time.Sleep(100 * time.Millisecond)
		}
		close(channel1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "Every 1s"
			time.Sleep(1 * time.Second)
		}
		close(channel2)
	}()

	// TODO: Fix the receiving logic
	// Should print messages as they arrive (not in forced order)
	// Hint: Use select or separate goroutines
}

// ========================================
// Exercise 4: Deadlock Example
// ========================================
func deadlockExample() {
	ch := make(chan int)

	// This will deadlock - why?
	// Uncomment to see:
	// ch <- 42  // Blocks forever waiting for receiver
	// fmt.Println(<-ch)

	// TODO: Fix by using a goroutine or buffered channel
	_ = ch
}

// ========================================
// Bonus: Ping Pong
// ========================================
func pingPong() {
	ping := make(chan string)
	pong := make(chan string)

	// TODO: Create two goroutines that play ping-pong
	// Goroutine 1: Sends "ping" to ping channel, waits for pong
	// Goroutine 2: Waits for ping, sends "pong" to pong channel
	// Do this 5 times then exit

	_ = ping
	_ = pong
}

// ========================================
// Main
// ========================================
func main() {
	fmt.Println("=== Exercise 1: Simple Channel ===")
	simpleChannel()

	fmt.Println("\n=== Exercise 2: Buffered Channel ===")
	bufferedChannel()

	fmt.Println("\n=== Exercise 3: Fixed Channel Demo ===")
	fixedChannelDemo()

	fmt.Println("\n=== Exercise 4: Deadlock ===")
	deadlockExample()

	fmt.Println("\n=== Bonus: Ping Pong ===")
	pingPong()
}
