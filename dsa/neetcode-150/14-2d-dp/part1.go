/*
===============================================================
Exercise 1: Unique Paths
===============================================================
Question:
Count unique paths from top-left to bottom-right (only right/down).

Example:
Input: m = 3, n = 7
Output: 28

Key Idea:
dp[i][j] = dp[i-1][j] + dp[i][j-1].

===============================================================
Exercise 2: Longest Common Subsequence
===============================================================
Question:
Find length of longest common subsequence.

Example:
Input: text1 = "abcde", text2 = "ace"
Output: 3 ("ace")

Key Idea:
dp[i][j] = match ? dp[i-1][j-1]+1 : max(dp[i-1][j], dp[i][j-1]).

===============================================================
Exercise 3: Best Time to Buy/Sell Stock with Cooldown
===============================================================
Question:
Max profit with cooldown day after selling.

Example:
Input: prices = [1,2,3,0,2]
Output: 3 (buy at 1, sell at 2, buy at 0, sell at 2)

Key Idea:
Three states: hold, sold, rest.

===============================================================
Exercise 4: Coin Change 2
===============================================================
Question:
Count ways to make amount with coins.

Example:
Input: amount = 5, coins = [1,2,5]
Output: 4

Key Idea:
Unbounded knapsack: dp[i] += dp[i-coin].

===============================================================
Exercise 5: Target Sum
===============================================================
Question:
Count ways to assign +/- to reach target.

Example:
Input: nums = [1,1,1,1,1], target = 3
Output: 5

Key Idea:
Transform to subset sum: find subset with sum = (total+target)/2.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Unique Paths - Solution
// ===============================================================
func uniquePaths(m int, n int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 2: Longest Common Subsequence - Solution
// ===============================================================
func longestCommonSubsequence(text1 string, text2 string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 3: Best Time to Buy/Sell with Cooldown - Solution
// ===============================================================
func maxProfitCooldown(prices []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Coin Change 2 - Solution
// ===============================================================
func change(amount int, coins []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(amount * n)
// Space Complexity: O(amount)

// ===============================================================
// Exercise 5: Target Sum - Solution
// ===============================================================
func findTargetSumWays(nums []int, target int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n * sum)
// Space Complexity: O(sum)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== 2-D DP Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
