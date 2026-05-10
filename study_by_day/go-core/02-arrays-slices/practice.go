package main

import "fmt"

// ========================================
// Exercise 1: Arrays vs Slices
// ========================================
func arrayVsSlice() {
	// TODO: Create an array of 5 integers
	// TODO: Create a slice from positions 1-4
	// TODO: Modify an element in the slice
	// TODO: Print both array and slice to see the relationship
}

// ========================================
// Exercise 2: Reverse a Slice
// ========================================
func reverseSlice(arr []int) {
	// TODO: Reverse the slice in-place
	// Hint: Swap elements from both ends moving toward center
}

// ========================================
// Exercise 3: Rotate Array Left
// ========================================
func rotateLeft(arr []int, k int) []int {
	// TODO: Rotate array left by k positions
	// Example: [1,2,3,4,5] rotated left by 2 = [3,4,5,1,2]
	return arr
}

// ========================================
// Exercise 4: Remove Duplicates
// ========================================
func removeDuplicates(arr []int) []int {
	// TODO: Remove duplicates while maintaining order
	// Hint: Use a map to track seen elements
	return []int{}
}

// ========================================
// Main - Test Your Solutions
// ========================================
func main() {
	fmt.Println("=== Exercise 1: Arrays vs Slices ===")
	arrayVsSlice()

	fmt.Println("\n=== Exercise 2: Reverse ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", nums)
	reverseSlice(nums)
	fmt.Println("After:", nums)

	fmt.Println("\n=== Exercise 3: Rotate Left ===")
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", nums2)
	rotated := rotateLeft(nums2, 2)
	fmt.Println("After:", rotated)

	fmt.Println("\n=== Exercise 4: Remove Duplicates ===")
	nums3 := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println("Before:", nums3)
	unique := removeDuplicates(nums3)
	fmt.Println("After:", unique)
}
