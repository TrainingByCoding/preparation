/*
===============================================================
Exercise 1: Valid Palindrome
===============================================================
Question:
A phrase is a palindrome if, after converting all uppercase letters
into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true

Example 2:
Input: s = "race a car"
Output: false

Key Idea:
Use two pointers from both ends, skip non-alphanumeric chars, compare lowercase.

===============================================================
Exercise 2: Two Sum II - Input Array Is Sorted
===============================================================
Question:
Find two numbers in sorted array that add up to target. Return 1-indexed positions.

Example:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]

Key Idea:
Two pointers: left at start, right at end. Move based on sum vs target.

===============================================================
Exercise 3: 3Sum
===============================================================
Question:
Find all unique triplets that sum to zero.

Example:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Key Idea:
Sort array, fix first number, use two pointers for remaining two. Skip duplicates.

===============================================================
Exercise 4: Container With Most Water
===============================================================
Question:
Find two lines that form container with most water.

Example:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49

Key Idea:
Two pointers from ends. Move pointer with smaller height to find potentially larger area.

===============================================================
Exercise 5: Trapping Rain Water
===============================================================
Question:
Calculate water trapped after raining on elevation map.

Example:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

Key Idea:
Two pointers with left_max and right_max tracking.

===============================================================
*/
package main

import (
	"fmt"
	"sort"
)

// ===============================================================
// Exercise 1: Valid Palindrome - Solution
// ===============================================================
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 2: Two Sum II - Solution
// ===============================================================
func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if target == sum {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else if sum > target {
			right--
		}

	}
	return []int{}
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 3: 3Sum - Solution
// ===============================================================
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {

			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				if left < right && nums[right] == nums[right-1] {
					right--
				}

				if left < right && nums[left] == nums[left+1] {
					left++
				}
			} else if sum > 0 {
				right--

			} else {
				left++
			}
		}
	}
	return result
}

// Time Complexity: O(n²)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Container With Most Water - Solution
// ===============================================================
func maxArea(height []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Trapping Rain Water - Solution
// ===============================================================
func trap(height []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	for {
		// fmt.Println("===== Enter string =====")
		// var s string
		// fmt.Scan(&s)
		// fmt.Println("result : ", isPalindrome(s))

		fmt.Println("Enter length of array:")
		var n int
		var nums []int
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			var val int
			fmt.Scan(&val)
			nums = append(nums, val)
		}
		fmt.Println("Entered array: %v", nums)
		fmt.Println("result:")
		fmt.Println(threeSum(nums))
	}
}
