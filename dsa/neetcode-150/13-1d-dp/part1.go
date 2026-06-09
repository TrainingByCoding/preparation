/*
===============================================================
Exercise 1: Climbing Stairs
===============================================================
Question:
How many ways to climb n stairs (1 or 2 steps at a time).

Example:
Input: n = 3
Output: 3 (1+1+1, 1+2, 2+1)

Key Idea:
dp[i] = dp[i-1] + dp[i-2] (Fibonacci).

===============================================================
Exercise 2: Min Cost Climbing Stairs
===============================================================
Question:
Minimum cost to reach top (can start from step 0 or 1).

Example:
Input: cost = [10,15,20]
Output: 15

Key Idea:
dp[i] = cost[i] + min(dp[i-1], dp[i-2]).

===============================================================
Exercise 3: House Robber
===============================================================
Question:
Maximum money robbing houses (can't rob adjacent).

Example:
Input: nums = [1,2,3,1]
Output: 4 (rob 1 and 3)

Key Idea:
dp[i] = max(dp[i-1], nums[i] + dp[i-2]).

===============================================================
Exercise 4: House Robber II
===============================================================
Question:
Houses arranged in circle (can't rob first and last).

Example:
Input: nums = [2,3,2]
Output: 3

Key Idea:
Solve twice: skip first house, skip last house.

===============================================================
Exercise 5: Longest Palindromic Substring
===============================================================
Question:
Find longest palindromic substring.

Example:
Input: s = "babad"
Output: "bab" or "aba"

Key Idea:
Expand around center or DP table.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Climbing Stairs - Solution
// ===============================================================
func climbStairs(n int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 2: Min Cost Climbing Stairs - Solution
// ===============================================================
func minCostClimbingStairs(cost []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 3: House Robber - Solution
// ===============================================================
func rob(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: House Robber II - Solution
// ===============================================================
func robII(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Longest Palindromic Substring - Solution
// ===============================================================
func longestPalindrome(s string) string {
	// TODO: Implement
	return ""
}

// Time Complexity: O(n^2)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== 1-D DP Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
