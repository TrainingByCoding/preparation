package main

import "fmt"

// ========================================
// Exercise 1: Max Sum Subarray of Size K
// ========================================
// Given: [2, 1, 5, 1, 3, 2], k=3
// Expected: 9 (subarray [5,1,3])

func maxSumWindow(arr []int, k int) int {
	// TODO: compute initial window sum for first k elements
	// Then slide: add arr[i], remove arr[i-k]
	return 0
}

// ========================================
// Exercise 2: Smallest Subarray with Sum >= Target
// ========================================
// Given: [2, 3, 1, 2, 4, 3], target=7
// Expected: 2 (subarray [4,3])

func minLenSubarray(arr []int, target int) int {
	// TODO: variable window
	// expand right, shrink left when sum >= target
	// track minimum window length
	return 0
}

// ========================================
// Exercise 3: Longest Substring Without Repeating Characters
// ========================================
// Given: "abcabcbb"
// Expected: 3 ("abc")

func lengthOfLongestSubstring(s string) int {
	// TODO: use map[byte]int to track last seen index of each char
	// left pointer, expand right
	// when char repeats: move left to max(left, lastSeen[char]+1)
	return 0
}

// ========================================
// Exercise 4: Count Subarrays with Sum == K
// ========================================
// Given: [1, 1, 1], k=2
// Expected: 2

func subarraySum(nums []int, k int) int {
	// TODO: use prefix sum + hashmap
	// count[prefixSum] = number of times this prefix sum has occurred
	// for each index: check if (currentSum - k) exists in map
	return 0
}

// ========================================
// Exercise 5: Maximum of All Subarrays of Size K
// ========================================
// Given: [1, 3, -1, -3, 5, 3, 6, 7], k=3
// Expected: [3, 3, 5, 5, 6, 7]

func maxSlidingWindow(nums []int, k int) []int {
	// TODO: use a deque ([]int storing indices)
	// maintain: deque front = index of current window max
	// remove indices out of window from front
	// remove smaller elements from back (they'll never be max)
	return nil
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 41: Arrays & Sliding Window ===\n")

	fmt.Println("Ex1: Max sum window k=3:", maxSumWindow([]int{2, 1, 5, 1, 3, 2}, 3))
	// Expected: 9

	fmt.Println("Ex2: Min len subarray sum>=7:", minLenSubarray([]int{2, 3, 1, 2, 4, 3}, 7))
	// Expected: 2

	fmt.Println("Ex3: Longest no-repeat substring:", lengthOfLongestSubstring("abcabcbb"))
	// Expected: 3

	fmt.Println("Ex4: Subarrays with sum==2:", subarraySum([]int{1, 1, 1}, 2))
	// Expected: 2

	fmt.Println("Ex5: Max sliding window k=3:", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// Expected: [3 3 5 5 6 7]
}

/*
SOLUTIONS:

func maxSumWindow(arr []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ { sum += arr[i] }
	max := sum
	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if sum > max { max = sum }
	}
	return max
}

func minLenSubarray(arr []int, target int) int {
	left, sum, minLen := 0, 0, len(arr)+1
	for right := 0; right < len(arr); right++ {
		sum += arr[right]
		for sum >= target {
			if right-left+1 < minLen { minLen = right-left+1 }
			sum -= arr[left]; left++
		}
	}
	if minLen == len(arr)+1 { return 0 }
	return minLen
}

func lengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	left, maxLen := 0, 0
	for right := 0; right < len(s); right++ {
		if idx, ok := last[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		last[s[right]] = right
		if right-left+1 > maxLen { maxLen = right-left+1 }
	}
	return maxLen
}

func subarraySum(nums []int, k int) int {
	count := map[int]int{0: 1}
	sum, result := 0, 0
	for _, n := range nums {
		sum += n
		result += count[sum-k]
		count[sum]++
	}
	return result
}

func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{} // stores indices
	result := []int{}
	for i, n := range nums {
		for len(deque) > 0 && deque[0] < i-k+1 { deque = deque[1:] }
		for len(deque) > 0 && nums[deque[len(deque)-1]] < n { deque = deque[:len(deque)-1] }
		deque = append(deque, i)
		if i >= k-1 { result = append(result, nums[deque[0]]) }
	}
	return result
}
*/
