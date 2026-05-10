# Day 5: Two Sum Problem

## 🎯 Today's Goal
Master the Two Sum problem - one of the most common interview questions!

## 📚 Files to Study
1. `../../master100/1twosum.go` - Two Sum implementation

## 🧠 Key Concepts
- Hash map (map in Go) for O(1) lookup
- Trading space for time complexity
- Complement pattern
- Avoiding duplicate solutions

## ✅ Learning Checklist
- [ ] Understand brute force O(n²) solution
- [ ] Understand optimized O(n) solution with map
- [ ] Can explain time/space complexity
- [ ] Know how to handle edge cases
- [ ] Can code from memory

## 🔍 Problem Statement

**Given**: An array of integers and a target sum  
**Return**: Indices of two numbers that add up to target  
**Constraint**: Each input has exactly one solution  

**Example**:
```
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]  // Because nums[0] + nums[1] = 2 + 7 = 9
```

## 💡 Solution Approaches

### Approach 1: Brute Force (Don't use in interviews!)
```go
// Time: O(n²), Space: O(1)
for i := 0; i < len(nums); i++ {
    for j := i+1; j < len(nums); j++ {
        if nums[i] + nums[j] == target {
            return []int{i, j}
        }
    }
}
```

### Approach 2: Hash Map (The correct answer!)
```go
// Time: O(n), Space: O(n)
seen := make(map[int]int)  // value -> index
for i, num := range nums {
    complement := target - num
    if idx, ok := seen[complement]; ok {
        return []int{idx, i}
    }
    seen[num] = i
}
```

## 🛠️ Practice Exercises

### Exercise 1: Implement Two Sum
```go
func twoSum(nums []int, target int) []int {
    // Your implementation
}
```

### Exercise 2: Variations

**Three Sum**:
```go
// Find all unique triplets that sum to 0
// Input: [-1, 0, 1, 2, -1, -4]
// Output: [[-1, -1, 2], [-1, 0, 1]]
```

**Two Sum - Sorted Array**:
```go
// Input array is sorted
// Can you do better than O(n) space?
// Hint: Two pointers
```

**Two Sum - Count Pairs**:
```go
// Return count of pairs that sum to target
// (Not indices, just count)
```

## 📝 Study Steps

### Step 1: Understand the Pattern (5 min)
The key insight:
```
If nums[i] + nums[j] = target
Then nums[j] = target - nums[i]

So while iterating, check if (target - current) was seen before!
```

### Step 2: Study the File (5 min)
Read `1twosum.go` and identify:
- How the map is used
- Why we check the map before adding
- Edge cases handled

### Step 3: Practice (10 min)
Implement all variations from scratch

## 🎓 Interview Tips

**What interviewers look for**:
1. ✅ Can you identify the complement pattern?
2. ✅ Do you think about time/space complexity?
3. ✅ Do you handle edge cases?
4. ✅ Can you code without syntax errors?

**Common follow-ups**:
- "What if array is sorted?" → Two pointers
- "What if we need all pairs?" → Don't return early
- "What about duplicates?" → Check problem constraints
- "Can you do it in-place?" → Usually no, need the map

**Complexity analysis you MUST know**:
- Brute force: O(n²) time, O(1) space
- Hash map: O(n) time, O(n) space
- Two pointers (sorted): O(n) time, O(1) space

## 🧪 Test Cases to Consider

```go
// Normal case
nums = [2, 7, 11, 15], target = 9  → [0, 1]

// Negatives
nums = [-3, 4, 3, 90], target = 0  → [0, 2]

// Same number twice
nums = [3, 3], target = 6  → [0, 1]

// Large numbers
nums = [1000000000, 1000000000], target = 2000000000

// Two elements only (minimum)
nums = [1, 2], target = 3  → [0, 1]
```

## ✅ Completion Checklist
- [ ] Understood brute force solution
- [ ] Understood optimized solution
- [ ] Can draw the algorithm flow
- [ ] Implemented from scratch in <10 minutes
- [ ] Solved all variations
- [ ] Can explain to interviewer clearly
- [ ] Updated PROGRESS_TRACKER.md
- [ ] Confidence: ___/10

## 🔄 Spaced Repetition
- Review Day 2 (Arrays & Slices)
- Review Day 4 (JSON - map usage is similar)

## 🎯 Mastery Criteria

**You've mastered Two Sum when**:
- [ ] Can code the solution in 5 minutes
- [ ] Can explain why map is used
- [ ] Can state time/space complexity instantly
- [ ] Can identify the pattern in new problems

## 🔜 Tomorrow: Day 6 - String Problems
Parenthesis matching and word counting!
