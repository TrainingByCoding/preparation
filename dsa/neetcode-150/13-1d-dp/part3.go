/*
===============================================================
Exercise 11: Longest Increasing Subsequence
===============================================================
Question:
Find length of longest increasing subsequence.

Example:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4 ([2,3,7,101])

Key Idea:
dp[i] = max length ending at i, or binary search optimization.

===============================================================
Exercise 12: Partition Equal Subset Sum
===============================================================
Question:
Check if array can be partitioned into two equal sum subsets.

Example:
Input: nums = [1,5,11,5]
Output: true (1+5+5 = 11)

Key Idea:
Subset sum DP: can we make target = sum/2.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 11: Longest Increasing Subsequence - Solution
// ===============================================================
func lengthOfLIS(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n log n) with binary search
// Space Complexity: O(n)

// ===============================================================
// Exercise 12: Partition Equal Subset Sum - Solution
// ===============================================================
func canPartition(nums []int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n * sum)
// Space Complexity: O(sum)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== 1-D DP Part 3 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
