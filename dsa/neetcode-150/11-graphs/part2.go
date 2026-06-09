/*
===============================================================
Exercise 6: Rotting Oranges
===============================================================
Question:
Find minimum time for all oranges to rot (multi-source BFS).

Example:
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

Key Idea:
Multi-source BFS from all initially rotten oranges.

===============================================================
Exercise 7: Walls and Gates
===============================================================
Question:
Fill each empty room with distance to nearest gate.

Example:
Input: rooms = [[INF,-1,0,INF],[INF,INF,INF,-1]]
Output: [[3,-1,0,1],[2,2,1,-1]]

Key Idea:
Multi-source BFS from all gates.

===============================================================
Exercise 8: Course Schedule
===============================================================
Question:
Check if you can finish all courses (detect cycle in DAG).

Example:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true

Key Idea:
Topological sort using DFS or Kahn's algorithm.

===============================================================
Exercise 9: Course Schedule II
===============================================================
Question:
Return ordering of courses to take.

Example:
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]

Key Idea:
Topological sort, track ordering.

===============================================================
Exercise 10: Redundant Connection
===============================================================
Question:
Find edge that can be removed to make tree.

Example:
Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]

Key Idea:
Union-Find to detect cycle.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Rotting Oranges - Solution
// ===============================================================
func orangesRotting(grid [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 7: Walls and Gates - Solution
// ===============================================================
func wallsAndGates(rooms [][]int) {
	// TODO: Implement
}

// Time Complexity: O(m * n)
// Space Complexity: O(m * n)

// ===============================================================
// Exercise 8: Course Schedule - Solution
// ===============================================================
func canFinish(numCourses int, prerequisites [][]int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(V + E)
// Space Complexity: O(V + E)

// ===============================================================
// Exercise 9: Course Schedule II - Solution
// ===============================================================
func findOrder(numCourses int, prerequisites [][]int) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(V + E)
// Space Complexity: O(V + E)

// ===============================================================
// Exercise 10: Redundant Connection - Solution
// ===============================================================
func findRedundantConnection(edges [][]int) []int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n * α(n)) - α is inverse Ackermann
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Graphs Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
