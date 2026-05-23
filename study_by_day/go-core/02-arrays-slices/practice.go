package main

import (
	"fmt"
	"slices"
)

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

	fmt.Println("\n=== Exercise 5: Insert Operations ===")
	arr := []int{1, 2, 3, 4, 5}
	arr = insertAtBeginning(arr, 0)
	fmt.Println("Insert 0 at beginning:", arr)
	arr = insertAtEnd(arr, 6)
	fmt.Println("Insert 6 at end:", arr)

	fmt.Println("\n=== Exercise 6: Sort array ===")
	nums4 := []int{3, 4, 5, 1, 2}
	fmt.Println("Before:", nums4)
	slices.Sort(nums4) // or use sort.Ints(nums4) Both modify original slice directly, returns nothing thats why - fmt.Println("After:", slices.Sort(nums4)) this will give error -  slices.Sort(nums4) (no value) used as value
	fmt.Println("After:", nums4)
}

// ========================================
// Exercise 5: Insert Operations
// ========================================
func insertAtBeginning(arr []int, val int) []int {
	// TODO: return new slice with val prepended
	return nil
}

func insertAtEnd(arr []int, val int) []int {
	// TODO: return new slice with val appended
	return nil
}

/*
SOLUTIONS:

func arrayVsSlice() {
	var arr [5]int
	for i := range arr { arr[i] = i + 1 }
	s := arr[1:4]    // [2 3 4] — backed by same array
	s[0] = 99        // modifies arr[1] too
	fmt.Println("Array:", arr)  // [1 99 3 4 5]
	fmt.Println("Slice:", s)    // [99 3 4]
}

func reverseSlice(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func rotateLeft(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	return append(arr[k:], arr[:k]...)
}

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func insertAtBeginning(arr []int, val int) []int {
	return append([]int{val}, arr...)
}

func insertAtEnd(arr []int, val int) []int {
	return append(arr, val)
}
*/
