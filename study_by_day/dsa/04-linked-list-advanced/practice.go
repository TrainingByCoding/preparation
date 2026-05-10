package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print("→")
		}
		head = head.Next
	}
	fmt.Println()
}

// ========================================
// Exercise 1: Has Cycle
// ========================================

func hasCycle(head *ListNode) bool {
	// TODO: Floyd's — fast moves 2x, slow moves 1x
	// If they meet: cycle exists
	return false
}

// ========================================
// Exercise 2: Detect Cycle Start
// ========================================

func detectCycle(head *ListNode) *ListNode {
	// TODO: after fast/slow meet, reset slow to head
	// move both one step at a time until they meet → that's the cycle start
	return nil
}

// ========================================
// Exercise 3: Intersection of Two Lists
// ========================================
// Two lists may share a suffix. Find the first shared node.
// Trick: if len(A)+len(B) combined, both pointers travel same total distance

func getIntersection(headA, headB *ListNode) *ListNode {
	// TODO: pointer A walks A then B; pointer B walks B then A
	// They meet at intersection (or both nil if no intersection)
	return nil
}

// ========================================
// Exercise 4: Reorder List
// ========================================
// 1→2→3→4→5 becomes 1→5→2→4→3

func reorderList(head *ListNode) {
	// TODO:
	// 1. Find middle
	// 2. Reverse second half
	// 3. Merge first and reversed second half alternately
}

// ========================================
// Exercise 5: Sort Linked List (Merge Sort)
// ========================================

func sortList(head *ListNode) *ListNode {
	// TODO: base case: nil or single node
	// 1. Find middle, split into two halves
	// 2. Recursively sort each half
	// 3. Merge sorted halves
	return nil
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

func main() {
	fmt.Println("=== Day 44: Cycle Detection & Advanced ===\n")

	// Ex1: build list with cycle
	list := buildList([]int{3, 2, 0, -4})
	// create cycle: tail points to node at index 1
	nodes := []*ListNode{}
	for n := list; n != nil; n = n.Next {
		nodes = append(nodes, n)
	}
	nodes[len(nodes)-1].Next = nodes[1]           // create cycle
	fmt.Println("Ex1 Has Cycle:", hasCycle(list)) // Expected: true
	cycleStart := detectCycle(list)
	if cycleStart != nil {
		fmt.Println("Ex2 Cycle Start:", cycleStart.Val) // Expected: 2
	}

	// Ex3: intersection
	common := buildList([]int{8, 4, 5})
	listA := buildList([]int{4, 1})
	listB := buildList([]int{5, 6, 1})
	// attach common to both
	tailA := listA
	for tailA.Next != nil {
		tailA = tailA.Next
	}
	tailA.Next = common
	tailB := listB
	for tailB.Next != nil {
		tailB = tailB.Next
	}
	tailB.Next = common
	inter := getIntersection(listA, listB)
	if inter != nil {
		fmt.Println("Ex3 Intersection:", inter.Val) // Expected: 8
	}

	// Ex4: reorder
	l := buildList([]int{1, 2, 3, 4, 5})
	reorderList(l)
	fmt.Print("Ex4 Reordered: ")
	printList(l)
	// Expected: 1→5→2→4→3

	// Ex5: sort
	unsorted := buildList([]int{4, 2, 1, 3})
	fmt.Print("Ex5 Sorted: ")
	printList(sortList(unsorted))
	// Expected: 1→2→3→4
}

/*
SOLUTIONS:

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next; fast = fast.Next.Next
		if slow == fast { return true }
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next; fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast { slow = slow.Next; fast = fast.Next }
			return slow
		}
	}
	return nil
}

func getIntersection(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil { a = headB } else { a = a.Next }
		if b == nil { b = headA } else { b = b.Next }
	}
	return a
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil { return }
	slow, fast := head, head
	for fast != nil && fast.Next != nil { slow = slow.Next; fast = fast.Next.Next }
	var prev *ListNode
	curr := slow.Next; slow.Next = nil
	for curr != nil { next := curr.Next; curr.Next = prev; prev = curr; curr = next }
	p1, p2 := head, prev
	for p2 != nil {
		next1, next2 := p1.Next, p2.Next
		p1.Next = p2; p2.Next = next1; p1 = next1; p2 = next2
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return head }
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil { slow = slow.Next; fast = fast.Next.Next }
	mid := slow.Next; slow.Next = nil
	return merge(sortList(head), sortList(mid))
}
*/
