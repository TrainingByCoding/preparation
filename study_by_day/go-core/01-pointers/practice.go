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
}
