/*
===============================================================
Exercise 1: Binary Search
===============================================================
Question:
Given sorted array and target, return index of target or -1.

Example:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4

Key Idea:
Standard binary search: compare middle element, adjust search range.

===============================================================
Exercise 2: Search a 2D Matrix
===============================================================
Question:
Search for value in m x n matrix where each row is sorted.

Example:
Input: matrix = [[1,3,5,7],[10,11,16,20]], target = 3
Output: true

Key Idea:
Treat 2D matrix as 1D sorted array, apply binary search.

===============================================================
Exercise 3: Koko Eating Bananas
===============================================================
Question:
Find minimum eating speed k to finish all bananas within h hours.

Example:
Input: piles = [3,6,7,11], h = 8
Output: 4

Key Idea:
Binary search on answer space (speed from 1 to max pile).

===============================================================
Exercise 4: Find Minimum in Rotated Sorted Array
===============================================================
Question:
Find minimum element in rotated sorted array.

Example:
Input: nums = [3,4,5,1,2]
Output: 1

Key Idea:
Binary search: compare with rightmost element to determine which half.

===============================================================
Exercise 5: Search in Rotated Sorted Array
===============================================================
Question:
Search target in rotated sorted array.

Example:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Key Idea:
Find which half is sorted, then determine which half contains target.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Binary Search - Solution
// ===============================================================
func search(nums []int, target int) int {
	// TODO: Implement
	return -1
}

// Time Complexity: O(log n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 2: Search a 2D Matrix - Solution
// ===============================================================
func searchMatrix(matrix [][]int, target int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(log(m*n))
// Space Complexity: O(1)

// ===============================================================
// Exercise 3: Koko Eating Bananas - Solution
// ===============================================================
func minEatingSpeed(piles []int, h int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n log m) - m is max pile size
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Find Minimum in Rotated Sorted Array - Solution
// ===============================================================
func findMin(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(log n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Search in Rotated Sorted Array - Solution
// ===============================================================
func searchRotated(nums []int, target int) int {
	// TODO: Implement
	return -1
}

// Time Complexity: O(log n)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Binary Search Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
