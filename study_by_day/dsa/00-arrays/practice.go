/*
Exercise 1 - contains Duplicate
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
Example 1:
Input: nums = [1,2,3,1]
Output: true
Explanation: The element 1 occurs at the indices 0 and 3.
Example 2:
Input: nums = [1,2,3,4]
Output: false
Explanation: All elements are distinct.
Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
===============================================================
Exercise 2 - Valid anagram
An anagram:

	Two strings contain the same characters with the same frequency, but possibly in a different order.
	e.g. listen - silent, rat - tar, aab - aba

===============================================================
Exercise 3 - group anagram
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Explanation: There is no string in strs that can be rearranged to form "bat". The strings "nat" and "tan" are anagrams as they can be rearranged to form each other. The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
===============================================================
Exercise 4 - two sum
===============================================================
Exercise 5: Top K Frequent Elements
===============================================================
*/
package main

import (
	"fmt"
	"slices"
)

// Exercise 1 - contains Duplicate
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return true		//if already exists, immedietely return true
		}
		seen[n] = true  // this adds n in map
	}
	return false
}
✅ Optimal Solution (HashMap / Set)
Time → O(n)
Space → O(n)

// Exercise 2 - Valid anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

/* Exercise 3 - group anagram
✅ Key Idea
All anagrams become SAME after sorting.
*/
func groupAnagrams(arr []string) [][]string {
	m := make(map[string][]string)
	result := [][]string{}
	for _, word := range arr {
		chars := []rune(word)
		slices.Sort(chars)
		sortedWord := string(chars)
		m[sortedWord] = append(m[sortedWord], word)
	}
	for _, group := range m {
		result = append(result, group)
	}

	return result
}

/* Exercise 4 - two sum
✅ Core Idea
For every number:
target - currentNumber
Check:
"Did I already see this required number before?"
If yes → solution found.
*/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
/*
Time Complexity
O(n)
Single traversal.

Space Complexity
O(n)
*/

/*
Exercise 5: Top K Frequent Elements
✅ Core Idea
First:
count frequency of each number
Then:
pick top k frequent numbers
*/
func main() {
	// Exercise 1 - contains Duplicate
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	// Exercise 2 - Valid anagram
	// var s1, s2 string
	// fmt.Println("enter two strings to check anagram or not")
	// fmt.Scan(&s1)
	// fmt.Scan(&s2)
	// fmt.Printf("result : %v", isAnagram(s1, s2))

	// Exercise 3 - group anagrams
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("Before grouping: %v\n", arr)
	fmt.Printf("After grouping: %v\n", groupAnagrams(arr))

	// Exercise 4 - two sum
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0 1]

}
