package main

import (
	"fmt"
	"sync"
	"time"
)

// ========================================
// Exercise 1: Basic Fan-In
// ========================================

// TODO: Implement fanIn — takes any number of <-chan int, returns one <-chan int
// All values from all input channels should appear in the output channel
// Output channel should be closed when all inputs are exhausted
func fanIn(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	// TODO: for each channel, launch a goroutine that reads and forwards values
	// TODO: close merged when all goroutines finish

	_ = wg // remove this when implemented
	return merged
}

// helper: creates a channel that sends values then closes
func produce(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func exercise1() {
	ch1 := produce(1, 2, 3)
	ch2 := produce(4, 5, 6)
	ch3 := produce(7, 8, 9)

	merged := fanIn(ch1, ch2, ch3)
	for val := range merged {
		fmt.Println(val) // order will vary — that's fine
	}
}

// ========================================
// Exercise 2: Parallel Search Fan-In
// ========================================

type SearchResult struct {
	Source string
	Data   string
}

// Simulates a slow search source
func searchSource(source, query string, delay time.Duration) <-chan SearchResult {
	ch := make(chan SearchResult, 1)
	go func() {
		time.Sleep(delay)
		ch <- SearchResult{Source: source, Data: fmt.Sprintf("[%s result for '%s']", source, query)}
		close(ch)
	}()
	return ch
}

// TODO: Implement fanInResults — merge multiple SearchResult channels into one
func fanInResults(channels ...<-chan SearchResult) <-chan SearchResult {
	// TODO: implement (same pattern as fanIn above)
	return nil
}

func exercise2() {
	query := "golang patterns"
	web := searchSource("Web", query, 300*time.Millisecond)
	img := searchSource("Images", query, 150*time.Millisecond)
	news := searchSource("News", query, 500*time.Millisecond)

	// TODO: fan-in all 3 sources and print results as they arrive
	merged := fanInResults(web, img, news)
	_ = merged
}

// ========================================
// Exercise 3: Fan-In with Timeout
// ========================================

func slowProducer(name string, values []int, delay time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			time.Sleep(delay)
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func exercise3() {
	ch1 := slowProducer("A", []int{1, 2, 3}, 400*time.Millisecond)
	ch2 := slowProducer("B", []int{4, 5, 6}, 700*time.Millisecond)
	ch3 := slowProducer("C", []int{7, 8, 9}, 200*time.Millisecond)

	merged := fanIn(ch1, ch2, ch3)
	timeout := time.After(1 * time.Second)

	// TODO: use select to receive from merged or timeout
	// Print values as they arrive
	// Stop when timeout fires, print "Timeout — stopping"
	_ = timeout
	_ = merged
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 33: Fan-In Pattern ===\n")

	fmt.Println("Exercise 1: Basic Fan-In (9 values from 3 channels)")
	exercise1()

	fmt.Println("\nExercise 2: Parallel Search")
	exercise2()

	fmt.Println("\nExercise 3: Fan-In with Timeout")
	exercise3()
}

/*
SOLUTIONS:

// Exercise 1 — fanIn:
func fanIn(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	forward := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			merged <- val
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go forward(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

// Exercise 2 — fanInResults:
func fanInResults(channels ...<-chan SearchResult) <-chan SearchResult {
	merged := make(chan SearchResult)
	var wg sync.WaitGroup

	forward := func(ch <-chan SearchResult) {
		defer wg.Done()
		for val := range ch {
			merged <- val
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go forward(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func exercise2() {
	query := "golang patterns"
	web  := searchSource("Web", query, 300*time.Millisecond)
	img  := searchSource("Images", query, 150*time.Millisecond)
	news := searchSource("News", query, 500*time.Millisecond)

	for result := range fanInResults(web, img, news) {
		fmt.Printf("Source: %s → %s\n", result.Source, result.Data)
	}
}

// Exercise 3:
for {
	select {
	case val, ok := <-merged:
		if !ok { return }
		fmt.Println("Got:", val)
	case <-timeout:
		fmt.Println("Timeout — stopping")
		return
	}
}
*/
