/*
===============================================================
Exercise 1: Reconstruct Itinerary
===============================================================
Question:
Reconstruct itinerary from list of tickets (Eulerian path).

Example:
Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
Output: ["JFK","MUC","LHR","SFO","SJC"]

Key Idea:
DFS with post-order to find Eulerian path.

===============================================================
Exercise 2: Min Cost to Connect All Points
===============================================================
Question:
Find minimum cost to connect all points (MST).

Example:
Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20

Key Idea:
Prim's or Kruskal's MST algorithm.

===============================================================
Exercise 3: Network Delay Time
===============================================================
Question:
Find time for signal to reach all nodes (shortest path).

Example:
Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2

Key Idea:
Dijkstra's algorithm from source node.

===============================================================
Exercise 4: Swim in Rising Water
===============================================================
Question:
Find minimum time to swim from top-left to bottom-right.

Example:
Input: grid = [[0,2],[1,3]]
Output: 3

Key Idea:
Binary search on time + BFS, or Dijkstra's.

===============================================================
Exercise 5: Alien Dictionary
===============================================================
Question:
Derive order of characters from sorted alien words.

Example:
Input: words = ["wrt","wrf","er","ett","rftt"]
Output: "wertf"

Key Idea:
Build graph from word pairs, topological sort.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Reconstruct Itinerary - Solution
// ===============================================================
func findItinerary(tickets [][]string) []string {
	// TODO: Implement
	return nil
}

// Time Complexity: O(E log E)
// Space Complexity: O(E)

// ===============================================================
// Exercise 2: Min Cost to Connect All Points - Solution
// ===============================================================
func minCostConnectPoints(points [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n^2 log n)
// Space Complexity: O(n^2)

// ===============================================================
// Exercise 3: Network Delay Time - Solution
// ===============================================================
func networkDelayTime(times [][]int, n int, k int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(E log V)
// Space Complexity: O(V + E)

// ===============================================================
// Exercise 4: Swim in Rising Water - Solution
// ===============================================================
func swimInWater(grid [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n^2 log n)
// Space Complexity: O(n^2)

// ===============================================================
// Exercise 5: Alien Dictionary - Solution
// ===============================================================
func alienOrder(words []string) string {
	// TODO: Implement
	return ""
}

// Time Complexity: O(C) - C is total chars in all words
// Space Complexity: O(1) - at most 26 chars

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Advanced Graphs Part 1 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
