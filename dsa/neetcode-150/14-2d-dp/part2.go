/*
===============================================================
Exercise 6: Interleaving String
===============================================================
Question:
Check if s3 is interleaving of s1 and s2.

Example:
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true

Key Idea:
dp[i][j] = true if s3[0:i+j] is interleaving of s1[0:i] and s2[0:j].

===============================================================
Exercise 7: Longest Increasing Path in Matrix
===============================================================
Question:
Find length of longest increasing path in matrix.

Example:
Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4 (1->2->6->9)

Key Idea:
DFS + memoization from each cell.

===============================================================
Exercise 8: Distinct Subsequences
===============================================================
Question:
Count distinct subsequences of s that equal t.

Example:
Input: s = "rabbbit", t = "rabbit"
Output: 3

Key Idea:
dp[i][j] = count of t[0:j] in s[0:i].

===============================================================
Exercise 9: Edit Distance
===============================================================
Question:
Minimum operations to convert word1 to word2.

Example:
Input: word1 = "horse", word2 = "ros"
Output: 3 (horse -> rorse -> rose -> ros)

Key Idea:
dp[i][j] = min edit distance for word1[0:i] -> word2[0:j].

===============================================================
Exercise 10: Burst Balloons
===============================================================
Question:
Maximum coins from bursting balloons.

Example:
Input: nums = [3,1,5,8]
Output: 167

Key Idea:
Interval DP: dp[i][j] = max coins bursting balloons i to j.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Interleaving String - Solution
// ===============================================================
func isInterleave(s1 string, s2 string, s3 string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 7: Longest Increasing Path in Matrix - Solution
// ===============================================================
func longestIncreasingPath(matrix [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 8: Distinct Subsequences - Solution
// ===============================================================
func numDistinct(s string, t string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 9: Edit Distance - Solution
// ===============================================================
func minDistance(word1 string, word2 string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 10: Burst Balloons - Solution
// ===============================================================
func maxCoins(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n^3)
// Space Complexity: O(n^2)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== 2-D DP Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
