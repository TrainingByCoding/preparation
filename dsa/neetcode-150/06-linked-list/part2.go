/*
===============================================================
Exercise 6: Add Two Numbers
===============================================================
Question:
Add two numbers represented by linked lists (digits in reverse order).

Example:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807

Key Idea:
Simulate addition with carry, create new nodes for result.

===============================================================
Exercise 7: Linked List Cycle
===============================================================
Question:
Determine if linked list has a cycle.

Example:
Input: head = [3,2,0,-4], pos = 1
Output: true

Key Idea:
Floyd's cycle detection: fast and slow pointers.

===============================================================
Exercise 8: Find the Duplicate Number
===============================================================
Question:
Find duplicate number in array (treated as linked list).

Example:
Input: nums = [1,3,4,2,2]
Output: 2

Key Idea:
Floyd's algorithm: treat array as linked list where nums[i] points to nums[nums[i]].

===============================================================
Exercise 9: LRU Cache
===============================================================
Question:
Design LRU cache with get and put in O(1) time.

Example:
LRUCache cache = new LRUCache(2);
cache.put(1, 1);
cache.get(1); // returns 1

Key Idea:
HashMap + doubly linked list for O(1) operations.

===============================================================
Exercise 10: Merge k Sorted Lists
===============================================================
Question:
Merge k sorted linked lists into one sorted list.

Example:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]

Key Idea:
Use min heap or divide and conquer.

===============================================================
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// ===============================================================
// Exercise 6: Add Two Numbers - Solution
// ===============================================================
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(max(m, n))
// Space Complexity: O(max(m, n))

// ===============================================================
// Exercise 7: Linked List Cycle - Solution
// ===============================================================
func hasCycle(head *ListNode) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 8: Find the Duplicate Number - Solution
// ===============================================================
func findDuplicate(nums []int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 9: LRU Cache - Solution
// ===============================================================
type LRUCache struct {
	// TODO: Implement
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{}
}

func (this *LRUCache) Get(key int) int {
	// TODO: Implement
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// TODO: Implement
}

// Time Complexity: O(1) for both get and put
// Space Complexity: O(capacity)

// ===============================================================
// Exercise 10: Merge k Sorted Lists - Solution
// ===============================================================
func mergeKLists(lists []*ListNode) *ListNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(N log k) - N total nodes, k lists
// Space Complexity: O(k)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Linked List Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
