package main

import "fmt"

// ========================================
// Exercise 1: Valid Parentheses
// ========================================
// Input: "()[]{}" → true, "([)]" → false, "{[]}" → true

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		// TODO: if opening bracket: push
		// if closing bracket: check top of stack matches
		_ = pairs
		_ = stack
	}
	return len(stack) == 0
}

// ========================================
// Exercise 2: Next Greater Element
// ========================================
// Input: [4, 1, 2] and [1, 3, 4, 2]
// For each element in nums1, find next greater in nums2
// Output: [-1, 3, -1]

func nextGreaterElement(nums1, nums2 []int) []int {
	// TODO: build map using monotonic stack on nums2
	// nextGreater[n] = next greater element of n in nums2 (-1 if none)
	// then for each in nums1, look up in map
	return nil
}

// ========================================
// Exercise 3: Daily Temperatures
// ========================================
// Input: [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]
// (how many days until warmer temperature)

func dailyTemperatures(temps []int) []int {
	result := make([]int, len(temps))
	stack := []int{} // stores indices
	// TODO: monotonic stack (decreasing)
	// when current temp > temp at stack top: pop and compute days
	return result
}

// ========================================
// Exercise 4: Queue using Two Stacks
// ========================================

type MyQueue struct {
	inbox  []int // push stack
	outbox []int // pop stack
}

// TODO: Implement Push(x int)
func (q *MyQueue) Push(x int) {
	// TODO: push to inbox
}

// TODO: Implement Pop() int
// Move all from inbox to outbox if outbox empty, then pop outbox
func (q *MyQueue) Pop() int {
	// TODO: implement
	return -1
}

// TODO: Implement Peek() int
func (q *MyQueue) Peek() int {
	// TODO: implement (similar to pop but don't remove)
	return -1
}

func (q *MyQueue) Empty() bool {
	return len(q.inbox) == 0 && len(q.outbox) == 0
}

func exercise4() {
	q := &MyQueue{}
	q.Push(1)
	q.Push(2)
	fmt.Println("Peek:", q.Peek())   // Expected: 1
	fmt.Println("Pop:", q.Pop())     // Expected: 1
	fmt.Println("Empty:", q.Empty()) // Expected: false
}

// ========================================
// Exercise 5: Largest Rectangle in Histogram
// ========================================
// Input: [2,1,5,6,2,3] → 10

func largestRectangle(heights []int) int {
	// TODO: monotonic increasing stack (stores indices)
	// when we find a bar shorter than stack top: compute area
	// width = i - stack[top-1] - 1
	// Add sentinel 0 at end to flush remaining stack
	return 0
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 45: Stack & Queue ===\n")

	fmt.Println("Ex1 Valid parens:")
	fmt.Println("  ()[]{}:", isValid("()[]{}")) // true
	fmt.Println("  ([)]:", isValid("([)]"))     // false
	fmt.Println("  {[]}:", isValid("{[]}"))     // true

	fmt.Println("\nEx2 Next Greater Element:")
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	// Expected: [-1 3 -1]

	fmt.Println("\nEx3 Daily Temperatures:")
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	// Expected: [1 1 4 2 1 1 0 0]

	fmt.Println("\nEx4 Queue from two stacks:")
	exercise4()

	fmt.Println("\nEx5 Largest Rectangle:")
	fmt.Println(largestRectangle([]int{2, 1, 5, 6, 2, 3}))
	// Expected: 10
}

/*
SOLUTIONS:

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] { return false }
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func nextGreaterElement(nums1, nums2 []int) []int {
	ng := map[int]int{}
	stack := []int{}
	for _, n := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < n {
			ng[stack[len(stack)-1]] = n
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}
	result := make([]int, len(nums1))
	for i, n := range nums1 {
		if v, ok := ng[n]; ok { result[i] = v } else { result[i] = -1 }
	}
	return result
}

func dailyTemperatures(temps []int) []int {
	result := make([]int, len(temps))
	stack := []int{}
	for i, t := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < t {
			idx := stack[len(stack)-1]; stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return result
}

func (q *MyQueue) Push(x int) { q.inbox = append(q.inbox, x) }
func (q *MyQueue) transfer() {
	if len(q.outbox) == 0 {
		for len(q.inbox) > 0 {
			top := q.inbox[len(q.inbox)-1]; q.inbox = q.inbox[:len(q.inbox)-1]
			q.outbox = append(q.outbox, top)
		}
	}
}
func (q *MyQueue) Pop() int { q.transfer(); top := q.outbox[len(q.outbox)-1]; q.outbox = q.outbox[:len(q.outbox)-1]; return top }
func (q *MyQueue) Peek() int { q.transfer(); return q.outbox[len(q.outbox)-1] }

func largestRectangle(heights []int) int {
	heights = append(heights, 0) // sentinel
	stack := []int{}; maxArea := 0
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]; stack = stack[:len(stack)-1]
			width := i; if len(stack) > 0 { width = i - stack[len(stack)-1] - 1 }
			if height*width > maxArea { maxArea = height * width }
		}
		stack = append(stack, i)
	}
	return maxArea
}
*/
