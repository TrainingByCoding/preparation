/*
===============================================================
Exercise 6: Palindromic Substrings
===============================================================
Question:
Count all palindromic substrings.

Example:
Input: s = "abc"
Output: 3 ("a", "b", "c")

Key Idea:
Expand around center for each position.

===============================================================
Exercise 7: Decode Ways
===============================================================
Question:
Count ways to decode string ('1'-'26' map to 'A'-'Z').

Example:
Input: s = "12"
Output: 2 ("AB" or "L")

Key Idea:
dp[i] = dp[i-1] + dp[i-2] (if valid 2-digit).

===============================================================
Exercise 8: Coin Change
===============================================================
Question:
Minimum coins to make amount.

Example:
Input: coins = [1,2,5], amount = 11
Output: 3 (5+5+1)

Key Idea:
dp[i] = min(dp[i], dp[i-coin] + 1).

===============================================================
Exercise 9: Maximum Product Subarray
===============================================================
Question:
Find subarray with maximum product.

Example:
Input: nums = [2,3,-2,4]
Output: 6 (2*3)

Key Idea:
Track both max and min (negatives can flip).

===============================================================
Exercise 10: Word Break
===============================================================
Question:
Check if string can be segmented into dictionary words.

Example:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true

Key Idea:
dp[i] = true if s[0:i] can be segmented.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Palindromic Substrings - Solution
// ===============================================================
func countSubstrings(s string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n^2)
// Space Complexity: O(1)

// ===============================================================
// Exercise 7: Decode Ways - Solution
// ===============================================================
func numDecodings(s string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 8: Coin Change - Solution
// ===============================================================
func coinChange(coins []int, amount int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(amount * n)
// Space Complexity: O(amount)

// ===============================================================
// Exercise 9: Maximum Product Subarray - Solution
// ===============================================================
func maxProduct(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 10: Word Break - Solution
// ===============================================================
func wordBreak(s string, wordDict []string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== 1-D DP Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
