package main

import "fmt"

// ========================================
// Exercise 1: Fibonacci Iterative
// ========================================
func fibonacci(n int) []int {
	// TODO: Generate first n Fibonacci numbers
	// Return as a slice
	return []int{}
}

// ========================================
// Exercise 2: Even Fibonacci Generator
// ========================================
func evenFibonacci(n int) chan int {
	ch := make(chan int)
	go func() {
		// TODO: Generate only even Fibonacci numbers
		// Close channel when done
	}()
	return ch
}

// ========================================
// Exercise 3: XOR Swap
// ========================================
func xorSwap(a, b int) (int, int) {
	// TODO: Swap a and b using XOR without temp variable
	return a, b
}

// ========================================
// Exercise 4: Find Single Number
// ========================================
func singleNumber(nums []int) int {
	// TODO: Find the number that appears once
	// All others appear twice
	// Example: [4,1,2,1,2] -> 4
	// Hint: Use XOR properties
	return 0
}

// ========================================
// Main
// ========================================
func main() {
	fmt.Println("=== Exercise 1: Fibonacci ===")
	fibs := fibonacci(10)
	fmt.Println("First 10:", fibs)

	fmt.Println("\n=== Exercise 2: Even Fibonacci ===")
	fmt.Println("Even Fibonacci numbers:")
	for num := range evenFibonacci(10) {
		fmt.Print(num, " ")
	}
	fmt.Println()

	fmt.Println("\n=== Exercise 3: XOR Swap ===")
	x, y := 5, 10
	fmt.Printf("Before: x=%d, y=%d\n", x, y)
	x, y = xorSwap(x, y)
	fmt.Printf("After: x=%d, y=%d\n", x, y)

	fmt.Println("\n=== Exercise 4: Single Number ===")
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	fmt.Printf("Single number in %v is: %d\n", nums, result)

	fmt.Println("\n=== Exercise 5: Word Count ===")
	input := []rune("  the earth  is blue  ")
	fmt.Printf("Word count: %d\n", countWords(input)) // Expected: 4
}

// ========================================
// Exercise 5: Word Count (from master100)
// ========================================
// Count words in a char slice (words separated by spaces)
// "  the earth  is blue  " → 4
func countWords(chars []rune) int {
	// TODO: count transitions from space to non-space
	return 0
}

/*
SOLUTIONS:

func fibonacci(n int) []int {
	result := make([]int, n)
	if n >= 1 { result[0] = 0 }
	if n >= 2 { result[1] = 1 }
	for i := 2; i < n; i++ { result[i] = result[i-1] + result[i-2] }
	return result
}

func evenFibonacci(n int) chan int {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			if a%2 == 0 { ch <- a }
			a, b = b, a+b
		}
		close(ch)
	}()
	return ch
}

func xorSwap(a, b int) (int, int) {
	a = a ^ b; b = a ^ b; a = a ^ b
	return a, b
}

func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums { result ^= n }
	return result
}

func countWords(chars []rune) int {
	count := 0
	for i, c := range chars {
		if c != ' ' && (i == 0 || chars[i-1] == ' ') { count++ }
	}
	return count
}
*/
