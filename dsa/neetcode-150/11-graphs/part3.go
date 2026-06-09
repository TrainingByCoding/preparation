/*
===============================================================
Exercise 11: Number of Connected Components
===============================================================
Question:
Count connected components in undirected graph.

Example:
Input: n = 5, edges = [[0,1],[1,2],[3,4]]
Output: 2

Key Idea:
Union-Find or DFS to count components.

===============================================================
Exercise 12: Graph Valid Tree
===============================================================
Question:
Check if graph is valid tree (connected + no cycles).

Example:
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true

Key Idea:
Must have exactly n-1 edges and be connected.

===============================================================
Exercise 13: Word Ladder
===============================================================
Question:
Find shortest transformation sequence from beginWord to endWord.

Example:
Input: beginWord = "hit", endWord = "cog",

	wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5 (hit -> hot -> dot -> dog -> cog)

Key Idea:
BFS treating words as nodes, edges between 1-char difference.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 11: Number of Connected Components - Solution
// ===============================================================
func countComponents(n int, edges [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(V + E)
// Space Complexity: O(V)

// ===============================================================
// Exercise 12: Graph Valid Tree - Solution
// ===============================================================
func validTree(n int, edges [][]int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(V + E)
// Space Complexity: O(V + E)

// ===============================================================
// Exercise 13: Word Ladder - Solution
// ===============================================================
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(m^2 * n) - m = word length, n = # words
// Space Complexity: O(m * n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Graphs Part 3 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
