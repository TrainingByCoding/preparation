/*
===============================================================
Exercise 1: Valid Parentheses
===============================================================
Question:
Given a string containing just '(', ')', '{', '}', '[' and ']', 
determine if the input string is valid.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "([)]"
Output: false

Example 3:
Input: s = "{[]}"
Output: true

Key Idea:
Use stack to match opening brackets with closing brackets.

===============================================================
Exercise 2: Min Stack
===============================================================
Question:
Design a stack that supports push, pop, top, and retrieving 
the minimum element in constant time.

Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3

Key Idea:
Maintain two stacks: one for values, one for minimums.

===============================================================
Exercise 3: Evaluate Reverse Polish Notation
===============================================================
Question:
Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Example:
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Key Idea:
Use stack to process operands and operators.

===============================================================
Exercise 4: Generate Parentheses
===============================================================
Question:
Given n pairs of parentheses, generate all combinations of 
well-formed parentheses.

Example:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Key Idea:
Backtracking with constraints: open count <= n, close count <= open.

===============================================================
Exercise 5: Daily Temperatures
===============================================================
Question:
Given array of daily temperatures, return array where answer[i] 
is the number of days until a warmer temperature.

Example:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Key Idea:
Monotonic decreasing stack to track indices.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Valid Parentheses - Solution
// ===============================================================
func isValid(s string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 2: Min Stack - Solution
// ===============================================================
type MinStack struct {
	// TODO: Implement
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	// TODO: Implement
}

func (this *MinStack) Pop() {
	// TODO: Implement
}

func (this *MinStack) Top() int {
	// TODO: Implement
	return 0
}

func (this *MinStack) GetMin() int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(1) for all operations
// Space Complexity: O(n)

// ===============================================================
// Exercise 3: Evaluate Reverse Polish Notation - Solution
// ===============================================================
func evalRPN(tokens []string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 4: Generate Parentheses - Solution
// ===============================================================
func generateParenthesis(n int) []string {
	// TODO: Implement
	return nil
}

// Time Complexity: O(4^n / sqrt(n))
// Space Complexity: O(n)

// ===============================================================
// Exercise 5: Daily Temperatures - Solution
// ===============================================================
func dailyTemperatures(temperatures []int) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Stack Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
