/*
===============================================================
Exercise 1: Invert Binary Tree
===============================================================
Question:
Invert a binary tree (swap left and right children recursively).

Example:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Key Idea:
Recursively swap left and right subtrees.

===============================================================
Exercise 2: Maximum Depth of Binary Tree
===============================================================
Question:
Find maximum depth (number of nodes from root to farthest leaf).

Example:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Key Idea:
Recursion: 1 + max(left depth, right depth).

===============================================================
Exercise 3: Diameter of Binary Tree
===============================================================
Question:
Find diameter (longest path between any two nodes).

Example:
Input: root = [1,2,3,4,5]
Output: 3 (path: 4-2-1-3)

Key Idea:
At each node, diameter = left height + right height.

===============================================================
Exercise 4: Balanced Binary Tree
===============================================================
Question:
Check if tree is height-balanced.

Example:
Input: root = [3,9,20,null,null,15,7]
Output: true

Key Idea:
For each node, check if |left height - right height| <= 1.

===============================================================
Exercise 5: Same Tree
===============================================================
Question:
Check if two trees are identical.

Example:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Key Idea:
Recursively check if values match and subtrees are identical.

===============================================================
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ===============================================================
// Exercise 1: Invert Binary Tree - Solution
// ===============================================================
func invertTree(root *TreeNode) *TreeNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(h) - recursion stack

// ===============================================================
// Exercise 2: Maximum Depth - Solution
// ===============================================================
func maxDepth(root *TreeNode) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(h)

// ===============================================================
// Exercise 3: Diameter of Binary Tree - Solution
// ===============================================================
func diameterOfBinaryTree(root *TreeNode) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(h)

// ===============================================================
// Exercise 4: Balanced Binary Tree - Solution
// ===============================================================
func isBalanced(root *TreeNode) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(h)

// ===============================================================
// Exercise 5: Same Tree - Solution
// ===============================================================
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(min(m, n))
// Space Complexity: O(min(h1, h2))

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Trees Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
