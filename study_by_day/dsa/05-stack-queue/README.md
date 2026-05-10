# Day 45: Stack & Queue

## 🎯 Goal
Implement and apply **Stack** and **Queue** patterns — used in DFS, BFS, expression evaluation, and many interview problems.

## 🧠 Key Concepts
- **Stack** — Last In First Out (LIFO). In Go: use `[]int` slice as stack
- **Queue** — First In First Out (FIFO). In Go: use `[]int` slice, or `container/list`
- **Monotonic Stack** — stack where elements are always increasing or decreasing
  - Used for: Next Greater Element, Previous Smaller Element, Largest Rectangle in Histogram

## 📖 Pattern

### Stack with slice
```go
stack := []int{}
stack = append(stack, 5)       // push
top := stack[len(stack)-1]     // peek
stack = stack[:len(stack)-1]   // pop
```

### Monotonic Stack (Next Greater Element)
```go
func nextGreater(nums []int) []int {
    result := make([]int, len(nums))
    stack := []int{} // stores indices

    for i, n := range nums {
        // pop elements smaller than current → n is their next greater
        for len(stack) > 0 && nums[stack[len(stack)-1]] < n {
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[idx] = n
        }
        stack = append(stack, i)
    }
    // remaining elements have no next greater
    for _, idx := range stack { result[idx] = -1 }
    return result
}
```

## ✅ Learning Checklist
- [ ] Implement stack using slice
- [ ] Valid parentheses using stack
- [ ] Next Greater Element using monotonic stack
- [ ] Implement queue using two stacks
- [ ] Largest Rectangle in Histogram

## 🛠️ Practice Exercises
1. Valid Parentheses — `()[]{}` check
2. Next Greater Element
3. Daily Temperatures — days until warmer day
4. Implement Queue using two Stacks
5. Largest Rectangle in Histogram
