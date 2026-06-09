/*
===============================================================
Exercise 1: Subsets
===============================================================
Question:
Return all possible subsets (power set).

Example:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Key Idea:
Backtracking: for each element, choose to include or not.

===============================================================
Exercise 2: Combination Sum
===============================================================
Question:
Find all unique combinations that sum to target (can reuse elements).

Example:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]

Key Idea:
Backtracking with index to allow reusing current element.

===============================================================
Exercise 3: Permutations
===============================================================
Question:
Return all possible permutations.

Example:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Key Idea:
Backtracking: swap elements or use visited array.

===============================================================
Exercise 4: Subsets II
===============================================================
Question:
Return all subsets (with duplicates in input).

Example:
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Key Idea:
Sort array, skip duplicates at same level.

===============================================================
Exercise 5: Combination Sum II
===============================================================
Question:
Find combinations summing to target (each number used once).

Example:
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: [[1,1,6],[1,2,5],[1,7],[2,6]]

Key Idea:
Sort, backtrack with start index, skip duplicates.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Subsets - Solution
// ===============================================================
func subsets(nums []int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(2^n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 2: Combination Sum - Solution
// ===============================================================
func combinationSum(candidates []int, target int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(2^target)
// Space Complexity: O(target)

// ===============================================================
// Exercise 3: Permutations - Solution
// ===============================================================
func permute(nums []int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n!)
// Space Complexity: O(n)

// ===============================================================
// Exercise 4: Subsets II - Solution
// ===============================================================
func subsetsWithDup(nums []int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(2^n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 5: Combination Sum II - Solution
// ===============================================================
func combinationSum2(candidates []int, target int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(2^n)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Backtracking Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
