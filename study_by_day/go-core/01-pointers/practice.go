package main

import "fmt"

// ========================================
// Exercise 1: Swap Function
// ========================================
func swap(a, b *int) {
	// TODO: Implement swap using pointers
	*a, *b = *b, *a
}

// ========================================
// Exercise 2: Modify Struct
// ========================================
type Person struct {
	Name string
	Age  int
}

func birthday(p *Person) {
	// TODO: Increment the person's age
	p.Age++
}

// ========================================
// Exercise 3: Nil Pointer Handling
// ========================================
func handleNilPointer() {
	var p *int
	fmt.Println("Pointer value:", p)
	p = new(int)
	*p = 10
	fmt.Println("Pointer value after assignment:", p)
	// TODO: Initialize the pointer and set a value
	// Then print the dereferenced value
}

// ========================================
// Main - Test Your Solutions
// ========================================
func main() {
	fmt.Println("=== Exercise 1: Swap ===")
	x, y := 5, 10
	fmt.Printf("Before: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After: x=%d, y=%d\n", x, y)

	fmt.Println("\n=== Exercise 2: Birthday ===")
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("Before:", person)
	birthday(&person)
	fmt.Println("After:", person)

	fmt.Println("\n=== Exercise 3: Nil Pointer ===")
	handleNilPointer()

	fmt.Println("\n=== Bonus: Heap Escape (from master100) ===")
	result := process()
	fmt.Println("Value:", *result)
}

// ========================================
// Bonus: Heap Escape / Stack vs Heap (from master100)
// ========================================
// Q: When does Go move a variable from stack to heap?
// A: When a pointer to a local variable escapes the function.
//    Go's escape analysis detects this at compile time.
//    Run: go build -gcflags="-m" to see escape decisions.

func process() *int {
	value := 10
	// 'value' escapes to heap because we return a pointer to it
	return &value
}

// Try: go build -gcflags="-m" ./go-core/01-pointers/
// Look for: "moved to heap: value"
