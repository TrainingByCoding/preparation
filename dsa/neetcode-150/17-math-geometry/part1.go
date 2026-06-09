/*
===============================================================
Exercise 1: Rotate Image
===============================================================
Question:
Rotate n×n matrix 90 degrees clockwise in-place.

Example:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Key Idea:
Transpose then reverse each row.

===============================================================
Exercise 2: Spiral Matrix
===============================================================
Question:
Return elements of matrix in spiral order.

Example:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Key Idea:
Track boundaries and shrink them after each layer.

===============================================================
Exercise 3: Set Matrix Zeroes
===============================================================
Question:
Set entire row and column to 0 if element is 0.

Example:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Key Idea:
Use first row/column as markers.

===============================================================
Exercise 4: Happy Number
===============================================================
Question:
Check if number eventually reaches 1 by sum of squares.

Example:
Input: n = 19
Output: true (1^2 + 9^2 = 82 -> ... -> 1)

Key Idea:
Floyd's cycle detection or hash set.

===============================================================
Exercise 5: Plus One
===============================================================
Question:
Add 1 to number represented as array of digits.

Example:
Input: digits = [1,2,3]
Output: [1,2,4]

Key Idea:
Handle carry from right to left.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Rotate Image - Solution
// ===============================================================
func rotate(matrix [][]int) {
	// TODO: Implement
}

// Time Complexity: O(n^2)
// Space Complexity: O(1)

// ===============================================================
// Exercise 2: Spiral Matrix - Solution
// ===============================================================
func spiralOrder(matrix [][]int) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(m * n)
// Space Complexity: O(1) - not counting output

// ===============================================================
// Exercise 3: Set Matrix Zeroes - Solution
// ===============================================================
func setZeroes(matrix [][]int) {
	// TODO: Implement
}

// Time Complexity: O(m * n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Happy Number - Solution
// ===============================================================
func isHappy(n int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(log n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Plus One - Solution
// ===============================================================
func plusOne(digits []int) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Math & Geometry Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
