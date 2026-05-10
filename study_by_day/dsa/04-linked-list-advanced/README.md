# Day 44: Linked List — Cycle Detection & Advanced

## 🎯 Goal
Detect cycles, find cycle start, and solve advanced linked list problems using Floyd's algorithm.

## 🧠 Key Concepts
- **Floyd's Cycle Detection (Tortoise & Hare)** — fast pointer moves 2x, slow 1x
  - If cycle: they meet inside the cycle
  - If no cycle: fast reaches nil
- **Find cycle start**: after meeting point, reset one pointer to head, move both 1 step — they meet at cycle start
- **Intersection**: find where two lists share a node (not same value, same reference)

## 📖 Pattern

### Cycle Detection
```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast { return true }
    }
    return false
}
```

### Find Cycle Start
```go
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next; fast = fast.Next.Next
        if slow == fast {
            slow = head  // reset one to head
            for slow != fast { slow = slow.Next; fast = fast.Next }
            return slow  // cycle start
        }
    }
    return nil
}
```

## ✅ Learning Checklist
- [ ] Detect cycle (Floyd's algorithm)
- [ ] Find cycle start node
- [ ] Find intersection of two lists
- [ ] Reorder list (L0→Ln→L1→Ln-1→...)
- [ ] Sort linked list (merge sort)

## 🛠️ Practice Exercises
1. Detect cycle (return true/false)
2. Find start of cycle (return node)
3. Intersection of two linked lists
4. Reorder list: L0→Ln→L1→Ln-1
5. Sort linked list using merge sort
