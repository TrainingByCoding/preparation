package main

import (
	"fmt"
	"sort"
)

// ========================================
// Exercise 1: Two Sum II (Sorted Array)
// ========================================
// Given sorted array + target, return 1-indexed positions
// Input: [2,7,11,15], target=9 → [1,2]

func twoSumSorted(nums []int, target int) []int {
	// TODO: left=0, right=len-1
	// if sum < target: left++, if sum > target: right--
	return nil
}

// ========================================
// Exercise 2: 3Sum
// ========================================
// Find all unique triplets summing to 0
// Input: [-1,0,1,2,-1,-4] → [[-1,-1,2],[-1,0,1]]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	// TODO: for each i (skip duplicates):
	//   two-pointer on nums[i+1..n-1] finding pairs that sum to -nums[i]
	//   skip duplicates after finding a match
	return result
}

// ========================================
// Exercise 3: Remove Duplicates In-Place
// ========================================
// Modify array in-place, return new length
// Input: [0,0,1,1,1,2,2,3,3,4] → 5 (first 5 elements: [0,1,2,3,4])

func removeDuplicates(nums []int) int {
	// TODO: slow/fast pointer
	// slow tracks last unique position, fast scans ahead
	return 0
}

// ========================================
// Exercise 4: Container with Most Water
// ========================================
// Given heights, find two lines forming max area container
// Input: [1,8,6,2,5,4,8,3,7] → 49

func maxWater(height []int) int {
	// TODO: left=0, right=n-1
	// area = min(height[left],height[right]) * (right-left)
	// move the shorter side inward (moving taller side can only decrease area)
	return 0
}

// ========================================
// Exercise 5: Trapping Rain Water
// ========================================
// Input: [0,1,0,2,1,0,1,3,2,1,2,1] → 6

func trap(height []int) int {
	// TODO: two pointer approach
	// Track maxLeft and maxRight
	// Water at position i = min(maxLeft, maxRight) - height[i]
	// Move the side with smaller max inward
	return 0
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 42: Two Pointers ===\n")

	fmt.Println("Ex1:", twoSumSorted([]int{2, 7, 11, 15}, 9))
	// Expected: [1 2]

	fmt.Println("Ex2:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// Expected: [[-1 -1 2] [-1 0 1]]

	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(arr)
	fmt.Println("Ex3 len:", n, "arr:", arr[:n])
	// Expected: 5 [0 1 2 3 4]

	fmt.Println("Ex4:", maxWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	// Expected: 49

	fmt.Println("Ex5:", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// Expected: 6
}

/*
SOLUTIONS:

func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target { return []int{left+1, right+1} }
		if sum < target { left++ } else { right-- }
	}
	return nil
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { continue }
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i]+nums[left]+nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i],nums[left],nums[right]})
				for left < right && nums[left] == nums[left+1] { left++ }
				for left < right && nums[right] == nums[right-1] { right-- }
				left++; right--
			} else if sum < 0 { left++ } else { right-- }
		}
	}
	return result
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 { return 0 }
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] { slow++; nums[slow] = nums[fast] }
	}
	return slow + 1
}

func maxWater(height []int) int {
	left, right, max := 0, len(height)-1, 0
	for left < right {
		h := height[left]; if height[right] < h { h = height[right] }
		area := h * (right - left)
		if area > max { max = area }
		if height[left] < height[right] { left++ } else { right-- }
	}
	return max
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxL, maxR, water := 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxL { maxL = height[left] } else { water += maxL - height[left] }
			left++
		} else {
			if height[right] >= maxR { maxR = height[right] } else { water += maxR - height[right] }
			right--
		}
	}
	return water
}
*/
