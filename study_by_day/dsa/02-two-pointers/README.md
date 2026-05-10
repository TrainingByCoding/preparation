# Day 42: Two Pointers

## 🎯 Goal
Master the **Two Pointers** technique — solve O(n²) brute-force problems in O(n).

## 🧠 Key Concepts
- Two pointers = use two index variables moving toward each other (or same direction)
- Works on **sorted arrays** (or when you can sort first)
- Classic problems: Two Sum (sorted), 3Sum, remove duplicates, container with most water
- Variants:
  - **Opposite ends** — left=0, right=n-1, move based on condition
  - **Same direction (fast/slow)** — detect cycles, find middle, remove duplicates

## 📖 Pattern

### Opposite Ends (Two Sum in sorted array)
```go
func twoSumSorted(arr []int, target int) (int, int) {
    left, right := 0, len(arr)-1
    for left < right {
        sum := arr[left] + arr[right]
        if sum == target { return left, right }
        if sum < target  { left++ }  // need bigger sum
        else             { right-- } // need smaller sum
    }
    return -1, -1
}
```

### Fast/Slow (remove duplicates in-place)
```go
func removeDuplicates(arr []int) int {
    if len(arr) == 0 { return 0 }
    slow := 0
    for fast := 1; fast < len(arr); fast++ {
        if arr[fast] != arr[slow] {
            slow++
            arr[slow] = arr[fast]
        }
    }
    return slow + 1 // new length
}
```

## ✅ Learning Checklist
- [ ] Two Sum on sorted array (opposite ends)
- [ ] 3Sum (sort + two pointers for each element)
- [ ] Remove duplicates in-place
- [ ] Container with most water
- [ ] Trapping rainwater

## 🛠️ Practice Exercises
1. Two Sum II (sorted array) → return indices
2. 3Sum → return all triplets summing to 0
3. Remove duplicates in-place, return new length
4. Container with most water (maximize area)
5. Trapping Rain Water
