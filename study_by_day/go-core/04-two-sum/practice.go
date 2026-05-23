package main

import "fmt"

// ========================================
// Exercise 1: Two Sum
// ========================================
func twoSum(nums []int, target int) []int {
	// TODO: Implement using hash map
	// Return indices of two numbers that sum to target
	return []int{}
}

// ========================================
// Exercise 2: Three Sum
// ========================================
func threeSum(nums []int) [][]int {
	// TODO: Find all unique triplets that sum to 0
	// Return array of triplets
	// Example: [-1,0,1,2,-1,-4] → [[-1,-1,2],[-1,0,1]]
	return [][]int{}
}

// ========================================
// Exercise 3: Two Sum - Sorted Array
// ========================================
func twoSumSorted(nums []int, target int) []int {
	// TODO: Use two pointers (O(1) space!)
	// Array is already sorted
	return []int{}
}

// ========================================
// Exercise 4: Count Pairs
// ========================================
func countPairs(nums []int, target int) int {
	// TODO: Return count of pairs that sum to target
	// Don't return indices, just count
	return 0
}

// ========================================
// Exercise 5: Two Sum - All Pairs
// ========================================
func twoSumAllPairs(nums []int, target int) [][]int {
	// TODO: Return all pairs (not just first one)
	// Example: [1,2,3,4,3], target=6
	// Output: [[2,4], [3,3]] (indices)
	return [][]int{}
}

// ========================================
// Helper: Print Test Case
// ========================================
func testTwoSum(nums []int, target int) {
	result := twoSum(nums, target)
	if len(result) == 2 {
		fmt.Printf("Input: %v, Target: %d → Indices: %v (nums[%d]=%d + nums[%d]=%d = %d)\n",
			nums, target, result, result[0], nums[result[0]], result[1], nums[result[1]], target)
	} else {
		fmt.Printf("Input: %v, Target: %d → No solution found\n", nums, target)
	}
}

// ========================================
// Main
// ========================================
func main() {
	fmt.Println("=== Exercise 1: Two Sum ===")
	testTwoSum([]int{2, 7, 11, 15}, 9)
	testTwoSum([]int{3, 2, 4}, 6)
	testTwoSum([]int{3, 3}, 6)
	testTwoSum([]int{-3, 4, 3, 90}, 0)

	fmt.Println("\n=== Exercise 2: Three Sum ===")
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Triplets that sum to 0: %v\n", result)

	fmt.Println("\n=== Exercise 3: Two Sum - Sorted ===")
	sorted := []int{1, 2, 3, 4, 5, 6}
	result2 := twoSumSorted(sorted, 7)
	fmt.Printf("Sorted array: %v, Target: 7 → %v\n", sorted, result2)

	fmt.Println("\n=== Exercise 4: Count Pairs ===")
	nums2 := []int{1, 2, 3, 4, 5}
	count := countPairs(nums2, 6)
	fmt.Printf("Array: %v, Target: 6 → Pairs: %d\n", nums2, count)

	fmt.Println("\n=== Exercise 5: All Pairs ===")
	nums3 := []int{1, 2, 3, 4, 3}
	allPairs := twoSumAllPairs(nums3, 6)
	fmt.Printf("Array: %v, Target: 6 → All pairs: %v\n", nums3, allPairs)
}

// ========================================
// Complexity Analysis + Solutions
// ========================================
/*
TWO SUM:
- Brute Force: O(n²) time, O(1) space
- Hash Map: O(n) time, O(n) space ✅
- Two Pointers (if sorted): O(n) time, O(1) space

THREE SUM:
- Brute Force: O(n³) time
- Sort + Two Pointers: O(n²) time ✅

KEY PATTERNS:
1. Hash map for O(n) lookups
2. Two pointers for sorted arrays
3. Sort first if space allows

SOLUTIONS:

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		if j, ok := seen[target-n]; ok { return []int{j, i} }
		seen[n] = i
	}
	return []int{}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { continue }
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] { left++ }
				for left < right && nums[right] == nums[right-1] { right-- }
				left++; right--
			} else if sum < 0 { left++ } else { right-- }
		}
	}
	return result
}

func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target { return []int{left, right} }
		if sum < target { left++ } else { right-- }
	}
	return []int{}
}

func countPairs(nums []int, target int) int {
	seen := make(map[int]int)
	count := 0
	for _, n := range nums {
		count += seen[target-n]
		seen[n]++
	}
	return count
}

func twoSumAllPairs(nums []int, target int) [][]int {
	seen := make(map[int][]int)
	result := [][]int{}
	for i, n := range nums {
		for _, j := range seen[target-n] {
			result = append(result, []int{j, i})
		}
		seen[n] = append(seen[n], i)
	}
	return result
}
*/
