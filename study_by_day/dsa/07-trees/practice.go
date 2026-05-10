package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// ========================================
// 1. Inorder (recursive + iterative)
// ========================================
// Inorder: Left → Root → Right

func inorder(root *TreeNode) []int {
	// TODO: recursive
	return nil
}

func inorderIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	curr := root
	// TODO: while curr != nil or stack not empty:
	//   go as left as possible (push to stack)
	//   pop, add val, move to right
	_ = curr
	_ = stack
	return result
}

// ========================================
// 2. Level Order (BFS)
// ========================================
// Returns each level as a slice: [[3],[9,20],[15,7]]

func levelOrder(root *TreeNode) [][]int {
	// TODO: queue-based BFS
	// process size = len(queue) nodes per level
	return nil
}

// ========================================
// 3. Max Depth
// ========================================
// Depth = 1 + max(left depth, right depth)

func maxDepth(root *TreeNode) int {
	// TODO: one-liner recursive
	return 0
}

// ========================================
// 4. Is Balanced
// ========================================
// Balanced = height difference of any subtree's children <= 1

func isBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// TODO: get left and right heights
		// if either returns -1 (unbalanced signal) or diff > 1: return -1
		// else return 1 + max(left, right)
		return 0
	}
	return height(root) != -1
}

// ========================================
// 5. Diameter
// ========================================
// Longest path between any two nodes (may not pass through root)
// diameter at node = leftHeight + rightHeight

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiam := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// TODO: left and right depths
		// update maxDiam = max(maxDiam, left+right)
		// return 1 + max(left, right)
		return 0
	}
	depth(root)
	return maxDiam
}

func main() {
	fmt.Println("=== dsa/07: Tree Traversals ===\n")

	//       3
	//      / \
	//     9  20
	//        / \
	//       15   7
	root := build([]interface{}{3, 9, 20, nil, nil, 15, 7})

	fmt.Println("Inorder (recursive):", inorder(root))          // [9 3 15 20 7]
	fmt.Println("Inorder (iterative):", inorderIterative(root)) // [9 3 15 20 7]
	fmt.Println("Level order:", levelOrder(root))               // [[3] [9 20] [15 7]]
	fmt.Println("Max depth:", maxDepth(root))                   // 3
	fmt.Println("Is balanced:", isBalanced(root))               // true
	fmt.Println("Diameter:", diameterOfBinaryTree(root))        // 4
}

/*
SOLUTIONS:

func inorder(root *TreeNode) []int {
	if root == nil { return nil }
	res := inorder(root.Left)
	res = append(res, root.Val)
	return append(res, inorder(root.Right)...)
}

func inorderIterative(root *TreeNode) []int {
	result := []int{}; stack := []*TreeNode{}; curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil { stack = append(stack, curr); curr = curr.Left }
		curr = stack[len(stack)-1]; stack = stack[:len(stack)-1]
		result = append(result, curr.Val); curr = curr.Right
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil { return nil }
	result := [][]int{}; queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue); level := []int{}
		for i := 0; i < size; i++ {
			n := queue[0]; queue = queue[1:]
			level = append(level, n.Val)
			if n.Left != nil { queue = append(queue, n.Left) }
			if n.Right != nil { queue = append(queue, n.Right) }
		}
		result = append(result, level)
	}
	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil { return 0 }
	l, r := maxDepth(root.Left), maxDepth(root.Right)
	if l > r { return 1+l }; return 1+r
}

// isBalanced height:
left, right := height(node.Left), height(node.Right)
if left == -1 || right == -1 { return -1 }
diff := left-right; if diff < 0 { diff = -diff }
if diff > 1 { return -1 }
if left > right { return 1+left }; return 1+right

// diameter depth:
l, r := depth(node.Left), depth(node.Right)
if l+r > maxDiam { maxDiam = l+r }
if l > r { return 1+l }; return 1+r
*/
