# Day 46: Binary Search

## 🎯 Goal
Master **Binary Search** and its variations — reduces O(n) to O(log n) for sorted data.

## 🧠 Key Concepts
- Works ONLY on sorted (or monotonic) data
- Template: left=0, right=n-1, mid=(left+right)/2
- Three variations:
  - **Find exact** — standard binary search
  - **Find leftmost** — first occurrence, first true
  - **Find rightmost** — last occurrence, last true
- "Binary search on answer" — when you can define a monotonic check function

## 📖 Pattern

### Standard
```go
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    for left <= right {
        mid := left + (right-left)/2  // avoid overflow
        if arr[mid] == target { return mid }
        if arr[mid] < target  { left = mid + 1 }
        else                  { right = mid - 1 }
    }
    return -1
}
```

### Find First (Leftmost) Occurrence
```go
func firstOccurrence(arr []int, target int) int {
    left, right, result := 0, len(arr)-1, -1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target { result = mid; right = mid-1 } // keep searching left
        else if arr[mid] < target { left = mid+1 }
        else { right = mid-1 }
    }
    return result
}
```

## ✅ Learning Checklist
- [ ] Standard binary search (exact match)
- [ ] Find first and last occurrence (search range)
- [ ] Search in rotated sorted array
- [ ] Find minimum in rotated sorted array
- [ ] Binary search on answer (e.g., minimum days to ship packages)

## 🛠️ Practice Exercises
1. Classic binary search
2. First and last position of element in sorted array
3. Search in rotated sorted array
4. Find minimum in rotated sorted array
5. Koko eating bananas (binary search on answer)
