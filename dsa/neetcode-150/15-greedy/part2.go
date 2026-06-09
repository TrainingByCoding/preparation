/*
===============================================================
Exercise 6: Merge Triplets to Form Target
===============================================================
Question:
Check if target triplet can be formed by merging triplets.

Example:
Input: triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
Output: true

Key Idea:
Greedy: track if we can reach each target value without exceeding.

===============================================================
Exercise 7: Partition Labels
===============================================================
Question:
Partition string into max partitions where each letter appears in only one part.

Example:
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]

Key Idea:
Track last occurrence of each char, extend partition until reaching last.

===============================================================
Exercise 8: Valid Parenthesis String
===============================================================
Question:
Check if string is valid with '*' as empty, '(', or ')'.

Example:
Input: s = "(*)"
Output: true

Key Idea:
Track min and max possible open brackets count.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Merge Triplets - Solution
// ===============================================================
func mergeTriplets(triplets [][]int, target []int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 7: Partition Labels - Solution
// ===============================================================
func partitionLabels(s string) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 8: Valid Parenthesis String - Solution
// ===============================================================
func checkValidString(s string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Greedy Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
