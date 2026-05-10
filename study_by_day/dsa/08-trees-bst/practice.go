package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(v int) *TreeNode { return &TreeNode{Val: v} }

// ========================================
// 1. Validate BST
// ========================================
// Every node: all left < node.Val < all right (not just direct children)

func isValidBST(root *TreeNode) bool {
	var validate func(node *TreeNode, min, max int) bool
	validate = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		// TODO: check node.Val > min && node.Val < max
		// recurse: left gets max=node.Val, right gets min=node.Val
		return false
	}
	return validate(root, -1<<63, 1<<63-1)
}

// ========================================
// 2. Lowest Common Ancestor (BST)
// ========================================
// In BST: if both p,q < node → go left; both > node → go right; else node is LCA

func lcaBST(root, p, q *TreeNode) *TreeNode {
	// TODO: no recursion stack needed — use the BST property to walk directly
	return nil
}

// ========================================
// 3. Kth Smallest Element in BST
// ========================================
// Inorder of BST is sorted → kth element of inorder = answer

func kthSmallest(root *TreeNode, k int) int {
	// TODO: inorder traversal, count down k
	// can do iterative to stop early
	return 0
}

// ========================================
// 4. Right Side View
// ========================================
// Level order, take last element of each level
// Input: [1,2,3,nil,5,nil,4] → [1,3,4]

func rightSideView(root *TreeNode) []int {
	// TODO: BFS, append last element of each level
	return nil
}

// ========================================
// 5. Path Sum II
// ========================================
// Find all root-to-leaf paths where sum == target

func pathSum(root *TreeNode, target int) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode, remaining int, path []int)
	dfs = func(node *TreeNode, remaining int, path []int) {
		if node == nil {
			return
		}
		// TODO: append node.Val to path
		// if leaf and remaining-node.Val == 0: append copy of path to result
		// recurse left and right
		// backtrack: remove last element from path
	}
	dfs(root, target, []int{})
	return result
}

func main() {
	fmt.Println("=== dsa/08: BST ===\n")

	//      5
	//     / \
	//    3   7
	//   / \ / \
	//  2  4 6  8
	root := newNode(5)
	root.Left = newNode(3)
	root.Right = newNode(7)
	root.Left.Left = newNode(2)
	root.Left.Right = newNode(4)
	root.Right.Left = newNode(6)
	root.Right.Right = newNode(8)

	fmt.Println("Valid BST:", isValidBST(root)) // true
	lca := lcaBST(root, newNode(2), newNode(4))
	if lca != nil {
		fmt.Println("LCA(2,4):", lca.Val)
	} // 3
	fmt.Println("3rd smallest:", kthSmallest(root, 3)) // 4
	fmt.Println("Right view:", rightSideView(root))    // [5 7 8]

	//    5
	//   / \
	//  4   8
	// /   / \
	// 11  13  4
	// /\      \
	// 7  2     1
	r2 := newNode(5)
	r2.Left = newNode(4)
	r2.Right = newNode(8)
	r2.Left.Left = newNode(11)
	r2.Right.Left = newNode(13)
	r2.Right.Right = newNode(4)
	r2.Left.Left.Left = newNode(7)
	r2.Left.Left.Right = newNode(2)
	r2.Right.Right.Right = newNode(1)
	fmt.Println("Path sum 22:", pathSum(r2, 22)) // [[5 4 11 2]]
}

/*
SOLUTIONS:

func isValidBST(root *TreeNode) bool {
	var validate func(*TreeNode, int, int) bool
	validate = func(n *TreeNode, min, max int) bool {
		if n == nil { return true }
		if n.Val <= min || n.Val >= max { return false }
		return validate(n.Left, min, n.Val) && validate(n.Right, n.Val, max)
	}
	return validate(root, -1<<63, 1<<63-1)
}

func lcaBST(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val { root = root.Left } else if p.Val > root.Val && q.Val > root.Val { root = root.Right } else { return root }
	}
	return nil
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}; curr := root
	for {
		for curr != nil { stack = append(stack, curr); curr = curr.Left }
		curr = stack[len(stack)-1]; stack = stack[:len(stack)-1]
		k--; if k == 0 { return curr.Val }
		curr = curr.Right
	}
}

func rightSideView(root *TreeNode) []int {
	if root == nil { return nil }
	result := []int{}; queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			n := queue[0]; queue = queue[1:]
			if i == size-1 { result = append(result, n.Val) }
			if n.Left != nil { queue = append(queue, n.Left) }
			if n.Right != nil { queue = append(queue, n.Right) }
		}
	}
	return result
}

func pathSum(root *TreeNode, target int) [][]int {
	result := [][]int{}
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, rem int, path []int) {
		if node == nil { return }
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && rem == node.Val {
			tmp := make([]int, len(path)); copy(tmp, path); result = append(result, tmp)
		}
		dfs(node.Left, rem-node.Val, path); dfs(node.Right, rem-node.Val, path)
	}
	dfs(root, target, []int{})
	return result
}
*/
