/*
===============================================================
Exercise 1: Reverse Linked List
===============================================================
Question:
Reverse a singly linked list.

Example:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Key Idea:
Iterative: Use three pointers (prev, curr, next).
Recursive: Reverse rest, then fix pointers.

===============================================================
Exercise 2: Merge Two Sorted Lists
===============================================================
Question:
Merge two sorted linked lists into one sorted list.

Example:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Key Idea:
Use dummy node, compare and link smaller nodes.

===============================================================
Exercise 3: Reorder List
===============================================================
Question:
Reorder list: L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → ...

Example:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Key Idea:
1. Find middle 2. Reverse second half 3. Merge alternately.

===============================================================
Exercise 4: Remove Nth Node From End
===============================================================
Question:
Remove nth node from end of list.

Example:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Key Idea:
Two pointers with n gap between them.

===============================================================
Exercise 5: Copy List with Random Pointer
===============================================================
Question:
Deep copy linked list with random pointer.

Example:
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: Copy of the list

Key Idea:
Use hash map to store old -> new node mapping.

===============================================================
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeWithRandom struct {
	Val    int
	Next   *NodeWithRandom
	Random *NodeWithRandom
}

// ===============================================================
// Exercise 1: Reverse Linked List - Solution
// ===============================================================
func reverseList(head *ListNode) *ListNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(1) iterative, O(n) recursive

// ===============================================================
// Exercise 2: Merge Two Sorted Lists - Solution
// ===============================================================
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(m + n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 3: Reorder List - Solution
// ===============================================================
func reorderList(head *ListNode) {
	// TODO: Implement
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Remove Nth Node From End - Solution
// ===============================================================
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Copy List with Random Pointer - Solution
// ===============================================================
func copyRandomList(head *NodeWithRandom) *NodeWithRandom {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Linked List Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
