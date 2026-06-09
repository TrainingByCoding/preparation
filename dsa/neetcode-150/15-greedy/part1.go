/*
===============================================================
Exercise 1: Maximum Subarray
===============================================================
Question:
Find contiguous subarray with largest sum.

Example:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6 ([4,-1,2,1])

Key Idea:
Kadane's algorithm: track current max and global max.

===============================================================
Exercise 2: Jump Game
===============================================================
Question:
Check if you can reach last index.

Example:
Input: nums = [2,3,1,1,4]
Output: true

Key Idea:
Track farthest reachable index.

===============================================================
Exercise 3: Jump Game II
===============================================================
Question:
Minimum jumps to reach last index.

Example:
Input: nums = [2,3,1,1,4]
Output: 2

Key Idea:
Greedy BFS: track current range and next range.

===============================================================
Exercise 4: Gas Station
===============================================================
Question:
Find starting station to complete circuit.

Example:
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3

Key Idea:
Track total gas and current tank, reset start when tank < 0.

===============================================================
Exercise 5: Hand of Straights
===============================================================
Question:
Check if cards can be rearranged into groups of consecutive cards.

Example:
Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true

Key Idea:
Greedy with sorted map: form groups starting from smallest.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Maximum Subarray - Solution
// ===============================================================
func maxSubArray(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 2: Jump Game - Solution
// ===============================================================
func canJump(nums []int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 3: Jump Game II - Solution
// ===============================================================
func jump(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Gas Station - Solution
// ===============================================================
func canCompleteCircuit(gas []int, cost []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Hand of Straights - Solution
// ===============================================================
func isNStraightHand(hand []int, groupSize int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n log n)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Greedy Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
