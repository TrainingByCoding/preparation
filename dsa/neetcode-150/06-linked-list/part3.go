/*
===============================================================
Exercise 11: Reverse Nodes in k-Group
===============================================================
Question:
Reverse nodes of linked list k at a time.

Example:
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

Example 2:
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]

Key Idea:
Count k nodes, reverse them, recursively handle rest.

===============================================================
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// ===============================================================
// Exercise 11: Reverse Nodes in k-Group - Solution
// ===============================================================
func reverseKGroup(head *ListNode, k int) *ListNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Linked List Part 3 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
