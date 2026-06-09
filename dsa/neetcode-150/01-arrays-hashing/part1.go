/*
===============================================================
Exercise 1: Contains Duplicate
===============================================================
Question:
Given an integer array nums, return true if any value appears at
least twice in the array, and return false if every element is distinct.

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

Key Idea:
Use a hash map to track seen elements. If we encounter an element
that's already in the map, we found a duplicate.

===============================================================
Exercise 2: Valid Anagram
===============================================================
Question:
Given two strings s and t, return true if t is an anagram of s,
and false otherwise.

An anagram: Two strings contain the same characters with the same
frequency, but possibly in a different order.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Example 3:
Input: s = "listen", t = "silent"
Output: true

Key Idea:
Count character frequencies in both strings using a hash map.
If all frequencies match (count becomes 0 for all chars), they're anagrams.

===============================================================
Exercise 3: Group Anagrams
===============================================================
Question:
Given an array of strings strs, group the anagrams together.
You can return the answer in any order.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Explanation:
- "bat" has no anagrams
- "nat" and "tan" are anagrams
- "ate", "eat", and "tea" are anagrams

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Key Idea:
All anagrams produce the SAME sorted string. Use sorted string as
hash map key to group anagrams together.

===============================================================
Exercise 4: Two Sum
===============================================================
Question:
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and
you may not use the same element twice.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: nums[0] + nums[1] == 9, so we return [0, 1]

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Key Idea:
For every number, calculate: complement = target - currentNumber
Check if we've already seen this complement in our hash map.
If yes → return indices. If no → store current number with its index.

===============================================================
Exercise 5: Top K Frequent Elements
===============================================================
Question:
Given an integer array nums and an integer k, return the k most
frequent elements. You may return the answer in any order.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]

Example 3:
Input: nums = [1,2,0,5,6,9,4,6,6,6,3,0,0], k = 3
Output: [6,0,1] (or any permutation)

Key Idea:
1. Count frequency of each number using hash map
2. Convert to array of (number, frequency) pairs
3. Sort by frequency in descending order
4. Return first k numbers

===============================================================
*/
package main

import (
	"fmt"
	"slices"
	"sort"
)

// ===============================================================
// Exercise 1: Contains Duplicate - Solution
// ===============================================================
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return true // if already exists, immediately return true
		}
		seen[n] = true // mark number as seen
	}
	return false
}

// Time Complexity: O(n) - single pass through array
// Space Complexity: O(n) - hash map stores up to n elements

// ===============================================================
// Exercise 2: Valid Anagram - Solution
// ===============================================================
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	// increment count for each char in s
	for _, ch := range s {
		count[ch]++
	}
	// decrement count for each char in t
	for _, ch := range t {
		count[ch]--
	}
	// if all counts are 0, strings are anagrams
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// Time Complexity: O(n) - iterate through both strings once
// Space Complexity: O(1) - map stores at most 26 characters (if lowercase English)

// ===============================================================
// Exercise 3: Group Anagrams - Solution
// ===============================================================
func groupAnagrams(arr []string) [][]string {
	m := make(map[string][]string)
	result := [][]string{}

	// for each word, sort it and use as key
	for _, word := range arr {
		chars := []rune(word)
		slices.Sort(chars)
		sortedWord := string(chars)
		m[sortedWord] = append(m[sortedWord], word)
	}

	// collect all groups
	for _, group := range m {
		result = append(result, group)
	}

	return result
}

// Time Complexity: O(n * k log k) - n strings, each of length k sorted
// Space Complexity: O(n * k) - storing all strings in hash map

// ===============================================================
// Exercise 4: Two Sum - Solution
// ===============================================================
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // value -> index

	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i} // found the pair
		}
		numMap[num] = i // store current number with its index
	}
	return nil
}

// Time Complexity: O(n) - single pass through array
// Space Complexity: O(n) - hash map stores up to n elements

// ===============================================================
// Exercise 5: Top K Frequent Elements - Solution
// ===============================================================
func topkFrequentElements(arr []int, k int) []int {
	// Step 1: Count frequency of each number
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}

	// Step 2: Convert to slice of (number, count) pairs
	type Pair struct {
		num   int
		count int
	}
	pairs := []Pair{}
	for key, val := range freq {
		pairs = append(pairs, Pair{key, val})
	}

	// Step 3: Sort by count in descending order
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].count > pairs[b].count
	})

	// Step 4: Return top k elements
	output := []int{}
	for i := 0; i < k; i++ {
		output = append(output, pairs[i].num)
	}
	return output
}

// Time Complexity: O(n log n) - counting O(n), sorting O(n log n)
// Space Complexity: O(n) - hash map and pairs array

// ===============================================================
// Main - Test all solutions
// ===============================================================
// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	// ===== Exercise 1: Contains Duplicate =====
	fmt.Println("\n===== Exercise 1: Contains Duplicate =====")
	fmt.Println("[1,2,3,1]:", containsDuplicate([]int{1, 2, 3, 1}))                               // true
	fmt.Println("[1,2,3,4]:", containsDuplicate([]int{1, 2, 3, 4}))                               // false
	fmt.Println("[1,1,1,3,3,4,3,2,4,2]:", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // true

	// ===== Exercise 2: Valid Anagram =====
	fmt.Println("\n===== Exercise 2: Valid Anagram =====")
	fmt.Println("anagram & nagaram:", isAnagram("anagram", "nagaram")) // true
	fmt.Println("rat & car:", isAnagram("rat", "car"))                 // false
	fmt.Println("listen & silent:", isAnagram("listen", "silent"))     // true

	// ===== Exercise 3: Group Anagrams =====
	fmt.Println("\n===== Exercise 3: Group Anagrams =====")
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("Input: %v\n", arr)
	fmt.Printf("Output: %v\n", groupAnagrams(arr))

	// ===== Exercise 4: Two Sum =====
	fmt.Println("\n===== Exercise 4: Two Sum =====")
	fmt.Println("[2,7,11,15], target=9:", twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
	fmt.Println("[3,2,4], target=6:", twoSum([]int{3, 2, 4}, 6))          // [1 2]
	fmt.Println("[3,3], target=6:", twoSum([]int{3, 3}, 6))               // [0 1]

	// ===== Exercise 5: Top K Frequent Elements =====
	fmt.Println("\n===== Exercise 5: Top K Frequent Elements =====")
	fmt.Println("[1,1,1,2,2,3], k=2:", topkFrequentElements([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println("[1,2,0,5,6,9,4,6,6,6,3,0,0], k=3:", topkFrequentElements([]int{1, 2, 0, 5, 6, 9, 4, 6, 6, 6, 3, 0, 0}, 3))
}
