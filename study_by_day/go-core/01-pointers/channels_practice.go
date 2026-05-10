package main

import (
	"fmt"
	"time"
)

// ========================================
// Exercise 1: Fix the Deadlock (Unbuffered Channel)
// ========================================
func exercise1Broken() {
	// ❌ This will deadlock!
	ch := make(chan int)
	ch <- 42 // Blocks forever - no receiver
	fmt.Println(<-ch)
}

func exercise1Fixed() {
	// TODO: Fix the deadlock by using a goroutine
	// Hint: Move either send or receive to a separate goroutine
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
}

// ========================================
// Exercise 2: Buffered Channel Practice
// ========================================
func exercise2() {
	// TODO: Create a buffered channel with capacity 3
	// Send 3 values without blocking
	// Receive and print all 3 values
	// Expected output: 1, 2, 3
	var buffchan = make(chan int, 3)
	buffchan <- 1
	buffchan <- 2
	buffchan <- 3
	for i := 0; i < 3; i++ {
		fmt.Println(<-buffchan)
	}
}

// ========================================
// Exercise 3: Same Goroutine - Buffered vs Unbuffered
// ========================================
func exercise3Unbuffered() {
	// ❌ This will deadlock with unbuffered channel
	ch := make(chan string)
	ch <- "hello" // Blocks forever
	msg := <-ch
	fmt.Println(msg)
}

func exercise3Buffered() {
	// TODO: Make this work using a buffered channel
	// Send "hello" and receive it in the same goroutine
	ch := make(chan string, 1)
	ch <- "hello"
	msg := <-ch
	fmt.Println(msg)
}

// ========================================
// Exercise 4: Simple Producer-Consumer
// ========================================
func producer(ch chan int) {
	// TODO: Send numbers 1 to 5 to the channel
	// Don't forget to close the channel when done!
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	// TODO: Receive and print all numbers from the channel
	// Use range to automatically stop when channel is closed
	for val := range ch {
		fmt.Println(val)
	}
}

func exercise4() {
	ch := make(chan int)
	fmt.Println(ch)
	// TODO: Start producer and consumer
	// Make sure consumer runs to completion
	go producer(ch)
	consumer(ch)
}

// ========================================
// Exercise 5: Print Numbers Alternately (Two Goroutines)
// ========================================
func printEven(ch chan int, done chan bool) {
	// TODO: Print even numbers from 0 to 10
	// Wait for signal on ch before printing
	// Send signal back on ch after printing
	// Signal done when finished
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ch <- i
			fmt.Println(<-ch)
		}
		done <- true
	}
}

func printOdd(ch chan int, done chan bool) {
	// TODO: Print odd numbers from 1 to 9
	// Wait for signal on ch before printing
	// Send signal back on ch after printing
	// Signal done when finished
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			ch <- i
			fmt.Println(<-ch)
		}
	}
	done <- true
}

func exercise5() {
	ch := make(chan int)
	done := make(chan bool)
	// TODO: Start both goroutines
	// TODO: Start the alternation (who goes first?)
	// TODO: Wait for both to finish
	go printEven(ch, done)
	go printOdd(ch, done)
	<-done
	<-done
}

// ========================================
// Exercise 8: Select Statement
// ========================================
func exercise8() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: sends after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	// Goroutine 2: sends after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// TODO: Use select to receive from whichever channel is ready first
	// Print both messages (you'll need to select twice)
}

// ========================================
// Exercise 9: Channel Direction (Read-only, Write-only)
// ========================================
func sendOnly(ch chan<- int) {
	// TODO: Send numbers 1, 2, 3 to channel
	// This function can only SEND (chan<-)
}

func receiveOnly(ch <-chan int) {
	// TODO: Receive and print all numbers
	// This function can only RECEIVE (<-chan)
}

func exercise9() {
	ch := make(chan int, 3)
	fmt.Println(ch)
	// TODO: Call sendOnly, then receiveOnly
}

// ========================================
// Exercise 10: Ping-Pong Pattern
// ========================================
func ping(pings chan<- string, pongs <-chan string, count int) {
	// TODO: Send "ping" count times
	// Wait for "pong" response each time
	// Print each ping sent
}

func pong(pings <-chan string, pongs chan<- string) {
	// TODO: Wait for "ping"
	// Respond with "pong"
	// Print each pong sent
	// Continue until pings channel is closed
	fmt.Println(pings)
	fmt.Println(pongs)
}

func exercise10() {
	pings := make(chan string)
	pongs := make(chan string)

	// TODO: Start ping and pong goroutines
	// TODO: Play 5 rounds of ping-pong
	// TODO: Clean up channels properly
	fmt.Println(pings)
	fmt.Println(pongs)
}

// ========================================
// Exercise 11: Timeout Pattern
// ========================================
func exercise11() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Result"
	}()
	fmt.Println(ch)
	// TODO: Use select with time.After() to timeout after 1 second
	// If timeout, print "Operation timed out"
	// If successful, print the result
}

// ========================================
// Exercise 12: Non-blocking Channel Operations
// ========================================
func exercise12() {
	ch := make(chan int, 1)
	fmt.Println(ch)
	// TODO: Try non-blocking send using select with default
	// Send 42 if channel is ready, otherwise print "Channel full"

	// TODO: Try non-blocking receive using select with default
	// Receive if available, otherwise print "Channel empty"
}

// ========================================
// BONUS: Fix the Buffer Overflow
// ========================================
func bonusExercise() {
	// ❌ This will deadlock - buffer too small
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // Blocks! No receiver and buffer full
	fmt.Println(<-ch)

	// TODO: Fix this by either:
	// 1. Increasing buffer size, OR
	// 2. Using a goroutine to receive
}

// ========================================
// Main - Uncomment exercises to test
// ========================================
func main() {
	fmt.Println("=== Channel & Goroutine Practice ===\n")

	// Exercise 1: Fix Deadlock
	fmt.Println("Exercise 1:")
	exercise1Fixed()

	// Exercise 2: Buffered Channel
	fmt.Println("\nExercise 2:")
	exercise2()

	// Exercise 3: Buffered vs Unbuffered
	fmt.Println("\nExercise 3:")
	exercise3Buffered()

	// Exercise 4: Producer-Consumer
	fmt.Println("\nExercise 4:")
	exercise4()

	// Exercise 5: Alternating Numbers
	fmt.Println("\nExercise 5:")
	exercise5()

	// Exercise 7: Done Signal
	// fmt.Println("\nExercise 7:")
	// exercise7()

	// Exercise 8: Select Statement
	// fmt.Println("\nExercise 8:")
	// exercise8()

	// Exercise 9: Channel Direction
	// fmt.Println("\nExercise 9:")
	// exercise9()

	// Exercise 10: Ping-Pong
	// fmt.Println("\nExercise 10:")
	// exercise10()

	// Exercise 11: Timeout
	// fmt.Println("\nExercise 11:")
	// exercise11()

	// Exercise 12: Non-blocking
	// fmt.Println("\nExercise 12:")
	// exercise12()

	// Bonus: Fix Buffer Overflow
	// fmt.Println("\nBonus Exercise:")
	// bonusExercise()

	fmt.Println("\n=== Practice Complete ===")
}

// ========================================
// SOLUTIONS (Don't peek until you try!)
// ========================================

/*
SOLUTION 1:
func exercise1Fixed() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
}

SOLUTION 2:
func exercise2() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

SOLUTION 3:
func exercise3Buffered() {
	ch := make(chan string, 1)
	ch <- "hello"
	msg := <-ch
	fmt.Println(msg)
}

SOLUTION 4:
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func exercise4() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

SOLUTION 5:
func printEven(ch chan int, done chan bool) {
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
		ch <- 1
		<-ch
	}
	done <- true
}

func printOdd(ch chan int, done chan bool) {
	for i := 1; i <= 9; i += 2 {
		<-ch
		fmt.Println(i)
		ch <- 1
	}
	done <- true
}

func exercise5() {
	ch := make(chan int)
	done := make(chan bool)

	go printEven(ch, done)
	go printOdd(ch, done)

	<-done
	<-done
}


SOLUTION 7:
func longRunningTask(done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task completed")
	done <- true
}

func exercise7() {
	done := make(chan bool)
	go longRunningTask(done)
	<-done
	fmt.Println("All tasks finished")
}

SOLUTION 8:
func exercise8() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

SOLUTION 9:
func sendOnly(ch chan<- int) {
	ch <- 1
	ch <- 2
	ch <- 3
}

func receiveOnly(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func exercise9() {
	ch := make(chan int, 3)
	sendOnly(ch)
	close(ch)
	receiveOnly(ch)
}

SOLUTION 10:
func ping(pings chan<- string, pongs <-chan string, count int) {
	for i := 0; i < count; i++ {
		pings <- "ping"
		fmt.Println("Sent: ping")
		<-pongs
	}
	close(pings)
}

func pong(pings <-chan string, pongs chan<- string) {
	for range pings {
		fmt.Println("Sent: pong")
		pongs <- "pong"
	}
}

func exercise10() {
	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, pongs, 5)
	pong(pings, pongs)
}

SOLUTION 11:
func exercise11() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
}

SOLUTION 12:
func exercise12() {
	ch := make(chan int, 1)

	select {
	case ch <- 42:
		fmt.Println("Sent 42")
	default:
		fmt.Println("Channel full")
	}

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	default:
		fmt.Println("Channel empty")
	}
}

BONUS SOLUTION:
func bonusExercise() {
	// Option 1: Increase buffer
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)

	// Option 2: Use goroutine
	ch2 := make(chan int, 2)
	go func() {
		fmt.Println(<-ch2)
		fmt.Println(<-ch2)
		fmt.Println(<-ch2)
	}()
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	time.Sleep(100 * time.Millisecond)
}
*/
