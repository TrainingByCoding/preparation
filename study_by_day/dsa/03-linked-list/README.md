# Day 43: Linked Lists — Reverse & Basics

## 🎯 Goal
Master linked list operations: reverse, find middle, merge — foundational interview problems.

## 🧠 Key Concepts
- Linked list: each node has `Val` + `Next` pointer
- No random access (unlike arrays) — must traverse
- Key operations: insert, delete, reverse, find middle, detect cycle
- Trick for many problems: **dummy head node** simplifies edge cases

## 📖 Pattern

### Reverse a Linked List
```go
func reverse(head *Node) *Node {
    var prev *Node
    curr := head
    for curr != nil {
        next := curr.Next  // save next
        curr.Next = prev   // reverse pointer
        prev = curr        // advance prev
        curr = next        // advance curr
    }
    return prev  // prev is new head
}
```

### Find Middle (slow/fast pointer)
```go
func findMiddle(head *Node) *Node {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next       // moves 1 step
        fast = fast.Next.Next  // moves 2 steps
    }
    return slow  // slow is at middle when fast reaches end
}
```

## ✅ Learning Checklist
- [ ] Reverse a linked list (iterative)
- [ ] Reverse a linked list (recursive)
- [ ] Find the middle node
- [ ] Merge two sorted linked lists
- [ ] Remove Nth node from end

## 🛠️ Practice Exercises
1. Reverse a linked list
2. Find the middle node
3. Merge two sorted linked lists
4. Remove Nth node from end (one pass, two pointers)
5. Check if linked list is a palindrome
