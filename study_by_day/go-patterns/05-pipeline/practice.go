package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ========================================
// Exercise 1: Number Pipeline
// Stage: generate → square → filterEven → print
// ========================================

// TODO: generate sends nums 1-10 to a channel, then closes
func generate(nums ...int) <-chan int {
	out := make(chan int)
	// TODO: implement
	return out
}

// TODO: square reads from in, sends n*n to output channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	// TODO: implement
	return out
}

// TODO: filterEven reads from in, only forwards even numbers
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	// TODO: implement
	return out
}

func exercise1() {
	nums := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := square(nums)
	evens := filterEven(squared)
	for v := range evens {
		fmt.Println(v)
	}
	// Expected: 4, 16, 36, 64, 100 (squares of 2,4,6,8,10)
}

// ========================================
// Exercise 2: String Pipeline
// Stage: generateWords → toUpper → filterLong → print
// ========================================

func generateWords(words ...string) <-chan string {
	out := make(chan string)
	// TODO: send each word then close
	return out
}

// TODO: toUpper — converts each word to uppercase
func toUpper(in <-chan string) <-chan string {
	out := make(chan string)
	// TODO: implement (use strings.ToUpper)
	_ = strings.ToUpper // hint
	return out
}

// TODO: filterLong — only forward words with len > 4
func filterLong(in <-chan string) <-chan string {
	out := make(chan string)
	// TODO: implement
	return out
}

func exercise2() {
	words := generateWords("go", "pipeline", "channel", "fan", "goroutine", "sync", "select")
	upper := toUpper(words)
	long := filterLong(upper)
	for w := range long {
		fmt.Println(w)
	}
	// Expected: PIPELINE, CHANNEL, GOROUTINE (len > 4 after uppercase)
}

// ========================================
// Exercise 3: CSV-like Pipeline
// Stage: generateRows → parseRow → validateRow → print
// ========================================

type Row struct {
	Name  string
	Score int
}

func generateRows(rows ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, r := range rows {
			out <- r
		}
		close(out)
	}()
	return out
}

// TODO: parseRow — splits "name,score" string into Row struct
// Use fmt.Sscanf or strings.Split
func parseRow(in <-chan string) <-chan Row {
	out := make(chan Row)
	// TODO: implement
	_ = unicode.IsDigit // just a hint import is used
	return out
}

// TODO: validateRow — only forward rows where Score >= 50
func validateRow(in <-chan Row) <-chan Row {
	out := make(chan Row)
	// TODO: implement
	return out
}

func exercise3() {
	raw := generateRows("Alice,85", "Bob,42", "Charlie,91", "Dave,30", "Eve,67")
	parsed := parseRow(raw)
	valid := validateRow(parsed)
	for row := range valid {
		fmt.Printf("Name: %s, Score: %d\n", row.Name, row.Score)
	}
	// Expected: Alice 85, Charlie 91, Eve 67
}

// ========================================
// Exercise 4: Cancellable Pipeline
// ========================================

// TODO: generateInfinite sends 1, 2, 3, ... forever
// Stop when done channel is closed
func generateInfinite(done <-chan struct{}) <-chan int {
	out := make(chan int)
	// TODO: implement
	return out
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func exercise4() {
	done := make(chan struct{})
	nums := generateInfinite(done)

	count := 0
	for n := range nums {
		if isPrime(n) {
			fmt.Println("Prime:", n)
			count++
			if count == 10 {
				close(done) // signal generator to stop
				break
			}
		}
	}
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 35: Pipeline Pattern ===\n")

	fmt.Println("Exercise 1: Number Pipeline (squares of evens)")
	exercise1()

	fmt.Println("\nExercise 2: String Pipeline (upper + long)")
	exercise2()

	fmt.Println("\nExercise 3: CSV Pipeline (parse + validate)")
	exercise3()

	fmt.Println("\nExercise 4: First 10 Primes (cancellable pipeline)")
	exercise4()
}

/*
SOLUTIONS:

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() { for _, n := range nums { out <- n }; close(out) }()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() { for n := range in { out <- n * n }; close(out) }()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() { for n := range in { if n%2 == 0 { out <- n } }; close(out) }()
	return out
}

func generateWords(words ...string) <-chan string {
	out := make(chan string)
	go func() { for _, w := range words { out <- w }; close(out) }()
	return out
}

func toUpper(in <-chan string) <-chan string {
	out := make(chan string)
	go func() { for w := range in { out <- strings.ToUpper(w) }; close(out) }()
	return out
}

func filterLong(in <-chan string) <-chan string {
	out := make(chan string)
	go func() { for w := range in { if len(w) > 4 { out <- w } }; close(out) }()
	return out
}

func parseRow(in <-chan string) <-chan Row {
	out := make(chan Row)
	go func() {
		for s := range in {
			parts := strings.Split(s, ",")
			var score int
			fmt.Sscanf(parts[1], "%d", &score)
			out <- Row{Name: parts[0], Score: score}
		}
		close(out)
	}()
	return out
}

func validateRow(in <-chan Row) <-chan Row {
	out := make(chan Row)
	go func() { for r := range in { if r.Score >= 50 { out <- r } }; close(out) }()
	return out
}

func generateInfinite(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; ; i++ {
			select {
			case <-done: return
			case out <- i:
			}
		}
	}()
	return out
}
*/
