/*
===============================================================
Exercise 1: Number of Islands
===============================================================
Question:
Count number of islands in 2D grid.

Example:
Input: grid = [

	["1","1","0","0","0"],
	["1","1","0","0","0"],
	["0","0","1","0","0"]]

Output: 3

Key Idea:
DFS/BFS to mark connected components.

===============================================================
Exercise 2: Clone Graph
===============================================================
Question:
Deep copy undirected graph.

Example:
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: Copy of graph

Key Idea:
DFS/BFS with HashMap to track old -> new node mapping.

===============================================================
Exercise 3: Max Area of Island
===============================================================
Question:
Find maximum area of island in grid.

Example:
Input: grid = [[0,0,1,0,0],[0,1,1,0,0]]
Output: 3

Key Idea:
DFS from each unvisited cell, count connected 1s.

===============================================================
Exercise 4: Pacific Atlantic Water Flow
===============================================================
Question:
Find cells where water can flow to both oceans.

Example:
Input: heights = [[1,2,2,3,5],[3,2,3,4,4]]
Output: [[0,4],[1,3],[1,4]]

Key Idea:
DFS from both oceans, find intersection.

===============================================================
Exercise 5: Surrounded Regions
===============================================================
Question:
Capture all regions surrounded by 'X'.

Example:
Input: board = [["X","X","X"],["X","O","X"],["X","X","X"]]
Output: [["X","X","X"],["X","X","X"],["X","X","X"]]

Key Idea:
DFS from border 'O's to mark safe cells, flip rest.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Number of Islands - Solution
// ===============================================================
func numIslands(grid [][]byte) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 2: Clone Graph - Solution
// ===============================================================
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	// TODO: Implement
	return nil
}

// Time Complexity: O(V + E)
// Space Complexity: O(V)

// ===============================================================
// Exercise 3: Max Area of Island - Solution
// ===============================================================
func maxAreaOfIsland(grid [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 4: Pacific Atlantic Water Flow - Solution
// ===============================================================
func pacificAtlantic(heights [][]int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 5: Surrounded Regions - Solution
// ===============================================================
func solve(board [][]byte) {
	// TODO: Implement
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Graphs Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
