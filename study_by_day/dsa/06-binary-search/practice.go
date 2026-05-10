package main

import "fmt"

// ========================================
// Exercise 1: Classic Binary Search
// ========================================
// Input: [-1,0,3,5,9,12], target=9 → 4

func search(nums []int, target int) int {
	// TODO: standard binary search
	return -1
}

// ========================================
// Exercise 2: First and Last Position
// ========================================
// Input: [5,7,7,8,8,10], target=8 → [3,4]
// Input: [5,7,7,8,8,10], target=6 → [-1,-1]

func searchRange(nums []int, target int) []int {
	// TODO: two binary searches
	// findFirst: when found, save and go left (right = mid-1)
	// findLast: when found, save and go right (left = mid+1)
	return []int{-1, -1}
}

// ========================================
// Exercise 3: Search in Rotated Sorted Array
// ========================================
// Input: [4,5,6,7,0,1,2], target=0 → 4
// Array was sorted then rotated at some pivot

func searchRotated(nums []int, target int) int {
	// TODO: one of the halves is always sorted
	// Check which half is sorted, determine if target is in that half
	// Adjust left/right accordingly
	return -1
}

// ========================================
// Exercise 4: Find Minimum in Rotated Array
// ========================================
// Input: [3,4,5,1,2] → 1

func findMin(nums []int) int {
	// TODO: binary search
	// If nums[mid] > nums[right]: minimum is in right half
	// Else: minimum is in left half (including mid)
	return 0
}

// ========================================
// Exercise 5: Koko Eating Bananas (Binary Search on Answer)
// ========================================
// Piles of bananas, Koko eats k bananas/hour
// Find minimum k such that she can eat all piles in h hours
// Input: piles=[3,6,7,11], h=8 → 4

func minEatingSpeed(piles []int, h int) int {
	// TODO: binary search on k (eating speed)
	// range: 1 to max(piles)
	// canFinish(k): check if sum(ceil(pile/k)) <= h
	// find minimum k where canFinish is true
	return 0
}

func canFinish(piles []int, k, h int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + k - 1) / k // ceil division
	}
	return hours <= h
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 46: Binary Search ===\n")

	fmt.Println("Ex1:", search([]int{-1, 0, 3, 5, 9, 12}, 9))
	// Expected: 4

	fmt.Println("Ex2:", searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	// Expected: [3 4]
	fmt.Println("Ex2:", searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	// Expected: [-1 -1]

	fmt.Println("Ex3:", searchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	// Expected: 4

	fmt.Println("Ex4:", findMin([]int{3, 4, 5, 1, 2}))
	// Expected: 1

	fmt.Println("Ex5:", minEatingSpeed([]int{3, 6, 7, 11}, 8))
	// Expected: 4
}

/*
SOLUTIONS:

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target { return mid }
		if nums[mid] < target { left = mid+1 } else { right = mid-1 }
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	findFirst := func() int {
		left, right, res := 0, len(nums)-1, -1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target { res = mid; right = mid-1 } else if nums[mid] < target { left = mid+1 } else { right = mid-1 }
		}
		return res
	}
	findLast := func() int {
		left, right, res := 0, len(nums)-1, -1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target { res = mid; left = mid+1 } else if nums[mid] < target { left = mid+1 } else { right = mid-1 }
		}
		return res
	}
	return []int{findFirst(), findLast()}
}

func searchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target { return mid }
		if nums[left] <= nums[mid] { // left half sorted
			if nums[left] <= target && target < nums[mid] { right = mid-1 } else { left = mid+1 }
		} else { // right half sorted
			if nums[mid] < target && target <= nums[right] { left = mid+1 } else { right = mid-1 }
		}
	}
	return -1
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] { left = mid+1 } else { right = mid }
	}
	return nums[left]
}

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, p := range piles { if p > right { right = p } }
	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, h) { right = mid } else { left = mid+1 }
	}
	return left
}
*/
