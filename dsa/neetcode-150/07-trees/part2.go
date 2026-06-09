/*
===============================================================
Exercise 6: Subtree of Another Tree
===============================================================
Question:
Check if tree s is subtree of tree t.

Example:
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true

Key Idea:
For each node in t, check if it matches s using isSameTree.

===============================================================
Exercise 7: Lowest Common Ancestor of BST
===============================================================
Question:
Find LCA of two nodes in Binary Search Tree.

Example:
Input: root = [6,2,8,0,4,7,9], p = 2, q = 8
Output: 6

Key Idea:
Use BST property: if both < root, go left; if both > root, go right.

===============================================================
Exercise 8: Binary Tree Level Order Traversal
===============================================================
Question:
Return level order traversal (BFS).

Example:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Key Idea:
Use queue for BFS, track level sizes.

===============================================================
Exercise 9: Binary Tree Right Side View
===============================================================
Question:
Return values visible from right side.

Example:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Key Idea:
BFS level order, take last node of each level.

===============================================================
Exercise 10: Count Good Nodes in Binary Tree
===============================================================
Question:
Count nodes where path from root has no greater values.

Example:
Input: root = [3,1,4,3,null,1,5]
Output: 4

Key Idea:
DFS with max value seen so far on path.

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
// Exercise 6: Subtree of Another Tree - Solution
// ===============================================================
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(m * n)
// Space Complexity: O(h)

// ===============================================================
// Exercise 7: Lowest Common Ancestor of BST - Solution
// ===============================================================
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// TODO: Implement
	return nil
}

// Time Complexity: O(h)
// Space Complexity: O(1)

// ===============================================================
// Exercise 8: Binary Tree Level Order Traversal - Solution
// ===============================================================
func levelOrder(root *TreeNode) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 9: Binary Tree Right Side View - Solution
// ===============================================================
func rightSideView(root *TreeNode) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 10: Count Good Nodes - Solution
// ===============================================================
func goodNodes(root *TreeNode) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(h)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Trees Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
