/*
===============================================================
Exercise 1: Single Number
===============================================================
Question:
Find the element that appears once (all others appear twice).

Example:
Input: nums = [2,2,1]
Output: 1

Key Idea:
XOR all numbers: a ^ a = 0, a ^ 0 = a.

===============================================================
Exercise 2: Number of 1 Bits
===============================================================
Question:
Count number of 1 bits in unsigned integer.

Example:
Input: n = 11 (binary: 1011)
Output: 3

Key Idea:
Use n & (n-1) to remove rightmost 1 bit.

===============================================================
Exercise 3: Counting Bits
===============================================================
Question:
For every number from 0 to n, count number of 1s in binary.

Example:
Input: n = 5
Output: [0,1,1,2,1,2]

Key Idea:
DP: count[i] = count[i >> 1] + (i & 1).

===============================================================
Exercise 4: Reverse Bits
===============================================================
Question:
Reverse bits of 32-bit unsigned integer.

Example:
Input: n = 00000010100101000001111010011100
Output:    964176192 (00111001011110000010100101000000)

Key Idea:
Build result bit by bit from right to left.

===============================================================
Exercise 5: Missing Number
===============================================================
Question:
Find missing number in array containing n distinct numbers.

Example:
Input: nums = [3,0,1]
Output: 2

Key Idea:
XOR all indices and values, or use sum formula.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Single Number - Solution
// ===============================================================
func singleNumber(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 2: Number of 1 Bits - Solution
// ===============================================================
func hammingWeight(num uint32) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(1) - max 32 iterations
// Space Complexity: O(1)

// ===============================================================
// Exercise 3: Counting Bits - Solution
// ===============================================================
func countBits(n int) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(1) - not counting output

// ===============================================================
// Exercise 4: Reverse Bits - Solution
// ===============================================================
func reverseBits(num uint32) uint32 {
	// TODO: Implement
	return 0
}

// Time Complexity: O(1)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Missing Number - Solution
// ===============================================================
func missingNumber(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Bit Manipulation Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
