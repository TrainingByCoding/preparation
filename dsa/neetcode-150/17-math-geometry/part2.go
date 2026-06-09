/*
===============================================================
Exercise 6: Pow(x, n)
===============================================================
Question:
Implement pow(x, n).

Example:
Input: x = 2.0, n = 10
Output: 1024.0

Key Idea:
Fast exponentiation: x^n = (x^2)^(n/2).

===============================================================
Exercise 7: Multiply Strings
===============================================================
Question:
Multiply two non-negative numbers represented as strings.

Example:
Input: num1 = "2", num2 = "3"
Output: "6"

Key Idea:
Grade school multiplication with carry.

===============================================================
Exercise 8: Detect Squares
===============================================================
Question:
Design data structure to count axis-aligned squares.

Example:
add([3,10]);
count([10,10]); // return 1

Key Idea:
Store points with frequency, for each query check all possible squares.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Pow(x, n) - Solution
// ===============================================================
func myPow(x float64, n int) float64 {
	// TODO: Implement
	return 0.0
}

// Time Complexity: O(log n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 7: Multiply Strings - Solution
// ===============================================================
func multiply(num1 string, num2 string) string {
	// TODO: Implement
	return ""
}

// Time Complexity: O(m * n)
// Space Complexity: O(m + n)

// ===============================================================
// Exercise 8: Detect Squares - Solution
// ===============================================================
type DetectSquares struct {
	// TODO: Implement
}

func ConstructorDetect() DetectSquares {
	return DetectSquares{}
}

func (this *DetectSquares) Add(point []int) {
	// TODO: Implement
}

func (this *DetectSquares) Count(point []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: Add O(1), Count O(n)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Math & Geometry Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
