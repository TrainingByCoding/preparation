# Day 41: Arrays & Sliding Window

## 🎯 Goal
Master **array fundamentals** and the **Sliding Window** technique — one of the most common interview patterns.

## 🧠 Key Concepts
- Sliding Window = maintain a window (subarray) and slide it, avoiding recomputation
- Use when: "find max/min/sum of subarray of size K" or "longest subarray with condition"
- Two variants:
  - **Fixed window** — size K never changes
  - **Variable window** — expand right, shrink left based on condition

## 📖 Pattern

### Fixed Window (max sum of subarray of size K)
```go
func maxSumWindow(arr []int, k int) int {
    windowSum := 0
    for i := 0; i < k; i++ { windowSum += arr[i] }
    maxSum := windowSum
    for i := k; i < len(arr); i++ {
        windowSum += arr[i] - arr[i-k]  // slide: add right, remove left
        if windowSum > maxSum { maxSum = windowSum }
    }
    return maxSum
}
```

### Variable Window (longest subarray with sum ≤ target)
```go
func longestSubarrayWithSum(arr []int, target int) int {
    left, sum, maxLen := 0, 0, 0
    for right := 0; right < len(arr); right++ {
        sum += arr[right]
        for sum > target { // shrink window
            sum -= arr[left]
            left++
        }
        if right-left+1 > maxLen { maxLen = right-left+1 }
    }
    return maxLen
}
```

## ✅ Learning Checklist
- [ ] Max sum subarray of size K
- [ ] Longest subarray with sum ≤ target
- [ ] Longest substring without repeating characters
- [ ] Minimum window substring
- [ ] Know time complexity: O(n) for all sliding window problems

## 🛠️ Practice Exercises
1. Max sum of subarray of size K
2. Smallest subarray with sum ≥ target (variable window)
3. Longest substring without repeating characters
4. Count subarrays with sum equal to K
5. Max of all subarrays of size K (deque approach)
