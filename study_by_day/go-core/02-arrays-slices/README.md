# Day 2: Arrays & Slices

## 🎯 Today's Goal
Master the difference between arrays and slices, and learn essential slice operations.

## 📚 Files to Study
1. `../../master100/arrayreverse.go` - Array reversal techniques
2. `../../master100/arrayrotation.go` - Array rotation patterns

## 🧠 Key Concepts
- Arrays vs Slices (fixed size vs dynamic)
- Slice operations: append, copy, slicing
- Array/Slice reversal algorithms
- Array rotation (left/right)
- In-place vs new array modifications

## ✅ Learning Checklist
- [ ] Understand array declaration: `[5]int`
- [ ] Understand slice declaration: `[]int`
- [ ] Can use `append()` correctly
- [ ] Can use slice notation `arr[start:end]`
- [ ] Can reverse an array in-place
- [ ] Can rotate an array

## 🛠️ Practice Exercises

Complete these in `practice.go`:

### Exercise 1: Array vs Slice
```go
// Create an array of 5 integers
// Create a slice from that array
// Modify the slice - what happens to the array?
```

### Exercise 2: Reverse a Slice
```go
// Reverse []int{1, 2, 3, 4, 5} in-place
// Expected: [5, 4, 3, 2, 1]
```

### Exercise 3: Rotate Array
```go
// Rotate []int{1, 2, 3, 4, 5} left by 2 positions
// Expected: [3, 4, 5, 1, 2]
```

### Exercise 4: Remove Duplicates
```go
// Remove duplicates from []int{1, 2, 2, 3, 4, 4, 5}
// Expected: [1, 2, 3, 4, 5]
```

## 📝 Study Steps (15-20 minutes)

### Step 1: Recall (5 min)
Without looking:
- What's the difference between `[5]int` and `[]int`?
- How do you append to a slice?
- How do you get a sub-slice?

### Step 2: Study (5 min)
Read both files and understand the reversal/rotation algorithms.

### Step 3: Practice (10 min)
Implement all 4 exercises.

## 🎓 Key Takeaways

**Array**: Fixed size, value type  
**Slice**: Dynamic size, reference type (points to underlying array)

**Common operations**:
```go
s := []int{1, 2, 3}      // Slice literal
s = append(s, 4)          // Add element
sub := s[1:3]             // Sub-slice [2, 3]
len(s)                    // Length
cap(s)                    // Capacity
```

## ✅ Completion
- [ ] Studied both files
- [ ] Completed all 4 exercises
- [ ] Can explain arrays vs slices
- [ ] Updated PROGRESS_TRACKER.md
- [ ] Confidence level: ___/10

## 🔄 Spaced Repetition
- Review Day 1 (Pointers) briefly today

## 🔜 Tomorrow: Day 3 - Basic Algorithms
