/*
===============================================================
Exercise 6: Time Based Key-Value Store
===============================================================
Question:
Design time-based key-value store that can store multiple values 
for same key at different timestamps.

Example:
store.set("foo", "bar", 1);
store.get("foo", 1); // returns "bar"
store.get("foo", 3); // returns "bar"

Key Idea:
Use hash map + binary search on timestamps.

===============================================================
Exercise 7: Median of Two Sorted Arrays
===============================================================
Question:
Find median of two sorted arrays.

Example:
Input: nums1 = [1,3], nums2 = [2]
Output: 2.0

Key Idea:
Binary search on smaller array to partition both arrays.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Time Based Key-Value Store - Solution
// ===============================================================
type TimeMap struct {
	// TODO: Implement
}

func ConstructorTimeMap() TimeMap {
	return TimeMap{}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// TODO: Implement
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// TODO: Implement
	return ""
}

// Time Complexity: Set O(1), Get O(log n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 7: Median of Two Sorted Arrays - Solution
// ===============================================================
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// TODO: Implement
	return 0.0
}

// Time Complexity: O(log(min(m,n)))
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Binary Search Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
