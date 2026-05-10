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
}
