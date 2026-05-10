# Day 47: Trees — Traversals & Basics

## 🎯 Goal
Master **binary tree traversals** — the foundation for all tree-based interview problems.

## 🧠 Key Concepts
- **Inorder** (Left → Root → Right) → gives sorted output for BST
- **Preorder** (Root → Left → Right) → useful for copying/serializing tree
- **Postorder** (Left → Right → Root) → useful for deleting tree, evaluating expressions
- **Level order (BFS)** → process level by level using a queue
- Recursive solutions are intuitive; iterative solutions use a stack

## 📖 Pattern

### Recursive Inorder
```go
func inorder(root *TreeNode, result *[]int) {
    if root == nil { return }
    inorder(root.Left, result)
    *result = append(*result, root.Val)
    inorder(root.Right, result)
}
```

### Level Order (BFS)
```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    result := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        level := []int{}
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]; queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil  { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        result = append(result, level)
    }
    return result
}
```

## ✅ Learning Checklist
- [ ] Inorder, preorder, postorder (recursive and iterative)
- [ ] Level order (BFS)
- [ ] Max depth of tree
- [ ] Check if balanced
- [ ] Diameter of binary tree

## 🛠️ Practice Exercises
1. Inorder traversal (recursive + iterative)
2. Level order traversal (return list of levels)
3. Max depth of binary tree
4. Check if binary tree is balanced
5. Diameter of binary tree
