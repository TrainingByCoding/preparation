package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Helper: build list from slice
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper: print list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" → ")
		}
		head = head.Next
	}
	fmt.Println()
}

// ========================================
// Exercise 1: Reverse Linked List
// ========================================
// Input: 1→2→3→4→5 → Output: 5→4→3→2→1

func reverseList(head *ListNode) *ListNode {
	// TODO: iterative with prev, curr, next
	return nil
}

// ========================================
// Exercise 2: Find Middle Node
// ========================================
// Input: 1→2→3→4→5 → Output: node with val 3
// Input: 1→2→3→4   → Output: node with val 3 (second middle)

func findMiddle(head *ListNode) *ListNode {
	// TODO: slow/fast pointer
	return nil
}

// ========================================
// Exercise 3: Merge Two Sorted Lists
// ========================================
// Input: 1→2→4 and 1→3→4 → Output: 1→1→2→3→4→4

func mergeSortedLists(l1, l2 *ListNode) *ListNode {
	// TODO: use a dummy head node to simplify
	// compare l1.Val and l2.Val, attach smaller to result
	return nil
}

// ========================================
// Exercise 4: Remove Nth Node From End
// ========================================
// Input: 1→2→3→4→5, n=2 → Output: 1→2→3→5

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// TODO: two pointer trick — advance fast by n steps
	// then move both fast and slow until fast.Next == nil
	// slow.Next is the node to remove
	// Use dummy head to handle edge case of removing head
	return nil
}

// ========================================
// Exercise 5: Palindrome Check
// ========================================
// Input: 1→2→2→1 → true
// Input: 1→2→3   → false

func isPalindrome(head *ListNode) bool {
	// TODO:
	// 1. Find middle
	// 2. Reverse second half
	// 3. Compare first and reversed second half
	return false
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 43: Linked Lists ===\n")

	fmt.Print("Ex1 Reverse: ")
	printList(reverseList(buildList([]int{1, 2, 3, 4, 5})))
	// Expected: 5 → 4 → 3 → 2 → 1

	mid := findMiddle(buildList([]int{1, 2, 3, 4, 5}))
	if mid != nil {
		fmt.Println("Ex2 Middle:", mid.Val) // Expected: 3
	}

	fmt.Print("Ex3 Merge: ")
	printList(mergeSortedLists(buildList([]int{1, 2, 4}), buildList([]int{1, 3, 4})))
	// Expected: 1 → 1 → 2 → 3 → 4 → 4

	fmt.Print("Ex4 Remove 2nd from end: ")
	printList(removeNthFromEnd(buildList([]int{1, 2, 3, 4, 5}), 2))
	// Expected: 1 → 2 → 3 → 5

	fmt.Println("Ex5 Palindrome 1→2→2→1:", isPalindrome(buildList([]int{1, 2, 2, 1})))
	fmt.Println("Ex5 Palindrome 1→2→3:", isPalindrome(buildList([]int{1, 2, 3})))
}

/*
SOLUTIONS:

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next; fast = fast.Next.Next
	}
	return slow
}

func mergeSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val { curr.Next = l1; l1 = l1.Next } else { curr.Next = l2; l2 = l2.Next }
		curr = curr.Next
	}
	if l1 != nil { curr.Next = l1 } else { curr.Next = l2 }
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i <= n; i++ { fast = fast.Next }
	for fast != nil { fast = fast.Next; slow = slow.Next }
	slow.Next = slow.Next.Next
	return dummy.Next
}

func isPalindrome(head *ListNode) bool {
	mid := findMiddle(head)
	second := reverseList(mid)
	p1, p2 := head, second
	for p2 != nil {
		if p1.Val != p2.Val { return false }
		p1 = p1.Next; p2 = p2.Next
	}
	return true
}
*/
