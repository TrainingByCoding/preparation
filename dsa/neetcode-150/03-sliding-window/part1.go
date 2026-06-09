/*
===============================================================
Exercise 1: Best Time to Buy and Sell Stock
===============================================================
Question:
Find max profit from buying and selling stock once.

Example:
Input: prices = [7,1,5,3,6,4]
Output: 5

Key Idea:
Track minimum price seen so far and maximum profit.

===============================================================
Exercise 2: Longest Substring Without Repeating Characters
===============================================================
Question:
Find length of longest substring without repeating characters.

Example:
Input: s = "abcabcbb"
Output: 3 ("abc")

Key Idea:
Sliding window with hash set. Expand right, contract left when duplicate found.

===============================================================
Exercise 3: Longest Repeating Character Replacement
===============================================================
Question:
Find longest substring with same letter after k replacements.

Example:
Input: s = "ABAB", k = 2
Output: 4

Key Idea:
Sliding window, track max frequency, window valid if size - maxFreq <= k.

===============================================================
Exercise 4: Permutation in String
===============================================================
Question:
Check if s2 contains permutation of s1.

Example:
Input: s1 = "ab", s2 = "eidbaooo"
Output: true

Key Idea:
Fixed sliding window of s1 length, compare character counts.

===============================================================
Exercise 5: Minimum Window Substring
===============================================================
Question:
Find minimum window in s containing all characters of t.

Example:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"

Key Idea:
Expanding window until valid, then contracting to find minimum.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Best Time to Buy and Sell Stock - Solution
// ===============================================================
func maxProfit(prices []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 2: Longest Substring Without Repeating - Solution
// ===============================================================
func lengthOfLongestSubstring(s string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(min(n, charset))

// ===============================================================
// Exercise 3: Longest Repeating Character Replacement - Solution
// ===============================================================
func characterReplacement(s string, k int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Permutation in String - Solution
// ===============================================================
func checkInclusion(s1 string, s2 string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Minimum Window Substring - Solution
// ===============================================================
func minWindow(s string, t string) string {
	// TODO: Implement
	return ""
}

// Time Complexity: O(n + m)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Sliding Window Practice =====")
	fmt.Println("Complete the TODO sections above")
}
