/*
===============================================================
Exercise 6: Product of Array Except Self
===============================================================
Question:
Given an integer array nums, return an array answer such that answer[i]
is equal to the product of all the elements of nums except nums[i].

You must write an algorithm that runs in O(n) time and without using
the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Key Idea:
Use prefix and suffix products. For each index, multiply product of all
elements before it with product of all elements after it.

===============================================================
Exercise 7: Valid Sudoku
===============================================================
Question:
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need
to be validated according to the following rules:
1. Each row must contain digits 1-9 without repetition
2. Each column must contain digits 1-9 without repetition
3. Each of the nine 3x3 sub-boxes must contain digits 1-9 without repetition

Example 1:
Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

Key Idea:
Use hash sets to track seen numbers in each row, column, and 3x3 box.

===============================================================
Exercise 8: Encode and Decode Strings
===============================================================
Question:
Design an algorithm to encode a list of strings to a single string.
The encoded string is then decoded back to the original list of strings.

Example:
Input: ["Hello","World"]
Output: ["Hello","World"]

Input: arr[] = ["abc", "!@"]
Output: ["abc","!@"]

The encoding rule used here is: Each string is stored as length + "/:" + actual string. During decoding, we first read the length before the delimiter "/:", and then extract that many characters as the original string.

===============================================================
Exercise 9: Longest Consecutive Sequence
===============================================================
Question:
Given an unsorted array of integers nums, return the length of the
longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive sequence is [1, 2, 3, 4]

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Key Idea:
Use a hash set. For each number, check if it's the start of a sequence
(num-1 not in set). If yes, count consecutive numbers from there.

===============================================================
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ===============================================================
// Exercise 6: Product of Array Except Self - Solution
// ===============================================================
func productExceptSelf(nums []int) []int {
	n := len(nums)

	prefix := make([]int, n)
	suffix := make([]int, n)
	answer := make([]int, n)

	prefix[0] = 1

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		answer[i] = prefix[i] * suffix[i]
	}

	return answer
}

// Time Complexity: O(n) - two passes through array
// Space Complexity: O(1) - not counting output array

// ===============================================================
// Exercise 7: Valid Sudoku - Solution
// ===============================================================
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			val := board[r][c]
			boxIdx := (r/3)*3 + c/3

			if rows[r][val] || cols[c][val] || boxes[boxIdx][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			boxes[boxIdx][val] = true
		}
	}

	return true
}

// Time Complexity: O(1) - always 9x9 grid
// Space Complexity: O(1) - fixed size hash sets

// ===============================================================
// Exercise 8: Encode and Decode Strings - Solution
// ===============================================================
func encode(strs []string) string {
	var result strings.Builder
	for _, val := range strs {
		result.WriteString(strconv.Itoa(len(val)))
		result.WriteString(":/")
		result.WriteString(val)
	}
	return result.String()
}

func decode(s string) []string {
	var result []string
	i := 0
	for i < len(s) {
		j := i
		if s[j] != ':' || s[j] != '/' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		start := j + 2
		end := start + length
		result = append(result, s[start:end])
		i = end
	}
	return result
}

// Time Complexity: O(n) - n is total length of all strings
// Space Complexity: O(n) - for encoded/decoded strings

// ===============================================================
// Exercise 9: Longest Consecutive Sequence - Solution
// ===============================================================
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}
	longest := 0
	for num := range set {
		// Start only from beginning of sequence
		if _, exists := set[num-1]; !exists {

			current := num
			length := 1

			for {
				if _, exists := set[current+1]; !exists {
					break
				}
				current++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

// Time Complexity: O(n) - each number visited at most twice
// Space Complexity: O(n) - hash set storage

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	// ===== Exercise 6: Product of Array Except Self =====
	fmt.Println("\n===== Exercise 6: Product of Array Except Self =====")
	fmt.Println("[1,2,3,4]:", productExceptSelf([]int{1, 2, 3, 4}))          // [24,12,8,6]
	fmt.Println("[-1,1,0,-3,3]:", productExceptSelf([]int{-1, 1, 0, -3, 3})) // [0,0,9,0,0]

	// ===== Exercise 7: Valid Sudoku =====
	fmt.Println("\n===== Exercise 7: Valid Sudoku =====")
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("Valid board:", isValidSudoku(board)) // true

	// ===== Exercise 8: Encode and Decode Strings =====
	fmt.Println("\n===== Exercise 8: Encode and Decode Strings =====")
	strs := []string{"Hello", "World"}
	encoded := encode(strs)
	decoded := decode(encoded)
	fmt.Println("Original:", strs)
	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decoded)

	// ===== Exercise 9: Longest Consecutive Sequence =====
	fmt.Println("\n===== Exercise 9: Longest Consecutive Sequence =====")
	fmt.Println("[100,4,200,1,3,2]:", longestConsecutive([]int{100, 4, 200, 1, 3, 2}))             // 4
	fmt.Println("[0,3,7,2,5,8,4,6,0,1]:", longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9
}
