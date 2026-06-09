/*
===============================================================
Exercise 6: Word Search
===============================================================
Question:
Check if word exists in 2D board.

Example:
Input: board = [["A","B","C","E"],["S","F","C","S"]], word = "ABCCED"
Output: true

Key Idea:
DFS + backtracking from each cell, mark visited.

===============================================================
Exercise 7: Palindrome Partitioning
===============================================================
Question:
Partition string into palindromic substrings.

Example:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Key Idea:
Backtracking: at each position, try all palindromic prefixes.

===============================================================
Exercise 8: Letter Combinations of Phone Number
===============================================================
Question:
Return all letter combinations for digit string.

Example:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Key Idea:
Backtracking with digit-to-letters mapping.

===============================================================
Exercise 9: N-Queens
===============================================================
Question:
Place n queens on n×n board so none attack each other.

Example:
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

Key Idea:
Backtracking: try placing queen in each column of current row.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Word Search - Solution
// ===============================================================
func exist(board [][]byte, word string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(m * n * 4^L) - L is word length
// Space Complexity: O(L)

// ===============================================================
// Exercise 7: Palindrome Partitioning - Solution
// ===============================================================
func partition(s string) [][]string {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 8: Letter Combinations - Solution
// ===============================================================
func letterCombinations(digits string) []string {
	// TODO: Implement
	return nil
}

// Time Complexity: O(4^n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 9: N-Queens - Solution
// ===============================================================
func solveNQueens(n int) [][]string {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n!)
// Space Complexity: O(n^2)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Backtracking Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
