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

// ========================================
// Bonus 2: Even/Odd Splitter (from master100)
// ========================================
// Send nums to two separate channels: odd and even
// Two goroutines consume and print each
func evenOddSplitter(nums []int) {
	// TODO: create oddCh and evenCh
	// producer goroutine: route each num to correct channel, close both
	// two consumer goroutines: print from each channel
	// WaitGroup to wait for consumers
}

/*
SOLUTIONS:

func simpleChannel() {
	ch := make(chan string)
	go func() { ch <- "hello from goroutine" }()
	fmt.Println(<-ch)
}

func bufferedChannel() {
	ch := make(chan int, 3)
	ch <- 1; ch <- 2; ch <- 3
	fmt.Println(<-ch, <-ch, <-ch)
}

func fixedChannelDemo() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 5; i++ { channel1 <- "Every 100ms"; time.Sleep(100*time.Millisecond) }
		close(channel1)
	}()
	go func() {
		for i := 0; i < 5; i++ { channel2 <- "Every 1s"; time.Sleep(1*time.Second) }
		close(channel2)
	}()
	// Fix: use goroutines per channel so neither blocks the other
	go func() { defer wg.Done(); for msg := range channel1 { fmt.Println(msg) } }()
	go func() { defer wg.Done(); for msg := range channel2 { fmt.Println(msg) } }()
	wg.Wait()
}

func deadlockExample() {
	ch := make(chan int, 1) // buffered = no deadlock
	ch <- 42
	fmt.Println(<-ch)
}

func pingPong() {
	ping := make(chan string)
	pong := make(chan string)
	go func() {
		for i := 0; i < 5; i++ { ping <- "ping"; fmt.Println("sent:", <-pong) }
		close(ping)
	}()
	go func() {
		for msg := range ping { fmt.Println("got:", msg); pong <- "pong" }
		close(pong)
	}()
	time.Sleep(100 * time.Millisecond)
}

func evenOddSplitter(nums []int) {
	oddCh := make(chan int)
	evenCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for _, n := range nums {
			if n%2 == 0 { evenCh <- n } else { oddCh <- n }
		}
		close(oddCh); close(evenCh)
	}()
	go func() { defer wg.Done(); for n := range oddCh { fmt.Println(n, "is odd") } }()
	go func() { defer wg.Done(); for n := range evenCh { fmt.Println(n, "is even") } }()
	wg.Wait()
}
*/
