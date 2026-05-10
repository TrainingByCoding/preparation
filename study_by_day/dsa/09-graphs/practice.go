package main

import "fmt"

// ========================================
// Graph represented as adjacency list
// ========================================

// 1. BFS — shortest path in unweighted graph
// Returns distance from start to each node (-1 if unreachable)

func bfs(graph map[int][]int, start int) map[int]int {
	dist := map[int]int{start: 0}
	queue := []int{start}
	// TODO: while queue not empty: dequeue, visit neighbors
	// if neighbor not visited: set dist, enqueue
	return dist
}

// ========================================
// 2. DFS — detect cycle in directed graph
// ========================================
// Use visited + recursion stack (inStack)

func hasCycle(graph map[int][]int, nodes int) bool {
	visited := make([]bool, nodes)
	inStack := make([]bool, nodes)
	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node] = true
		inStack[node] = true
		for _, neighbor := range graph[node] {
			// TODO: if not visited: recurse; if inStack: cycle found
		}
		inStack[node] = false
		return false
	}
	for i := 0; i < nodes; i++ {
		if !visited[i] && dfs(i) {
			return true
		}
	}
	return false
}

// ========================================
// 3. Number of Islands
// ========================================
// '1' = land, '0' = water. Count connected land regions.
// Input: grid of '1'/'0'

func numIslands(grid [][]byte) int {
	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0' // mark visited
		// TODO: recurse in 4 directions
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// ========================================
// 4. Clone Graph
// ========================================

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[*Node]*Node{}
	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if clone, ok := visited[n]; ok {
			return clone
		}
		clone := &Node{Val: n.Val}
		visited[n] = clone
		// TODO: for each neighbor, recurse and append to clone.Neighbors
		return clone
	}
	return dfs(node)
}

// ========================================
// 5. Topological Sort (Kahn's BFS algorithm)
// ========================================
// For directed acyclic graph, return nodes in dependency order

func topoSort(numNodes int, edges [][2]int) []int {
	inDegree := make([]int, numNodes)
	adj := make([][]int, numNodes)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		inDegree[e[1]]++
	}
	queue := []int{}
	for i, d := range inDegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	result := []int{}
	// TODO: BFS — process zero in-degree nodes
	// for each neighbor: decrement inDegree; if reaches 0, enqueue
	return result
}

func main() {
	fmt.Println("=== dsa/09: Graphs ===\n")

	graph := map[int][]int{0: {1, 2}, 1: {3}, 2: {3}, 3: {}}
	fmt.Println("BFS dist from 0:", bfs(graph, 0)) // map[0:0 1:1 2:1 3:2]

	cyclic := map[int][]int{0: {1}, 1: {2}, 2: {0}}
	acyclic := map[int][]int{0: {1}, 1: {2}, 2: {}}
	fmt.Println("Has cycle (cyclic):", hasCycle(cyclic, 3))   // true
	fmt.Println("Has cycle (acyclic):", hasCycle(acyclic, 3)) // false

	grid := [][]byte{
		{'1', '1', '0', '0'},
		{'1', '0', '0', '1'},
		{'0', '0', '0', '1'},
	}
	fmt.Println("Islands:", numIslands(grid)) // 2

	// Topo sort: 5→2, 5→0, 4→0, 4→1, 2→3, 3→1
	edges := [][2]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}}
	fmt.Println("Topo sort:", topoSort(6, edges)) // valid order e.g. [4 5 0 2 3 1]
}

/*
SOLUTIONS:

func bfs(graph map[int][]int, start int) map[int]int {
	dist := map[int]int{start: 0}; queue := []int{start}
	for len(queue) > 0 {
		node := queue[0]; queue = queue[1:]
		for _, nb := range graph[node] {
			if _, ok := dist[nb]; !ok { dist[nb] = dist[node]+1; queue = append(queue, nb) }
		}
	}
	return dist
}

func hasCycle dfs neighbor:
if !visited[neighbor] { if dfs(neighbor) { return true } } else if inStack[neighbor] { return true }

func numIslands dfs directions:
dfs(i+1, j); dfs(i-1, j); dfs(i, j+1); dfs(i, j-1)

func cloneGraph neighbor loop:
for _, nb := range n.Neighbors { clone.Neighbors = append(clone.Neighbors, dfs(nb)) }

func topoSort BFS:
for len(queue) > 0 {
	node := queue[0]; queue = queue[1:]; result = append(result, node)
	for _, nb := range adj[node] { inDegree[nb]--; if inDegree[nb] == 0 { queue = append(queue, nb) } }
}
*/
