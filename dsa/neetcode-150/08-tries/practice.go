/*
===============================================================
Exercise 1: Implement Trie (Prefix Tree)
===============================================================
Question:
Implement trie with insert, search, and startsWith methods.

Example:
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // true
trie.startsWith("app"); // true

Key Idea:
Use tree structure where each node has 26 children (a-z).

===============================================================
Exercise 2: Design Add and Search Words Data Structure
===============================================================
Question:
Design data structure supporting addWord and search with '.' wildcard.

Example:
addWord("bad");
search("pad"); // false
search("b.."); // true

Key Idea:
Trie + DFS for wildcard matching.

===============================================================
Exercise 3: Word Search II
===============================================================
Question:
Find all words from dictionary that exist in 2D board.

Example:
Input: board = [["o","a","a","n"],["e","t","a","e"]],

	words = ["oath","pea","eat","rain"]

Output: ["eat","oath"]

Key Idea:
Build trie from words, DFS on board with trie traversal.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Implement Trie - Solution
// ===============================================================
type Trie struct {
	// TODO: Implement
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	// TODO: Implement
}

func (this *Trie) Search(word string) bool {
	// TODO: Implement
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(m) for all operations (m = word length)
// Space Complexity: O(n * m) - n words, m average length

// ===============================================================
// Exercise 2: Add and Search Words - Solution
// ===============================================================
type WordDictionary struct {
	// TODO: Implement
}

func ConstructorWordDict() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	// TODO: Implement
}

func (this *WordDictionary) Search(word string) bool {
	// TODO: Implement
	return false
}

// Time Complexity: AddWord O(m), Search O(m) or O(26^m) with wildcards
// Space Complexity: O(n * m)

// ===============================================================
// Exercise 3: Word Search II - Solution
// ===============================================================
func findWords(board [][]byte, words []string) []string {
	// TODO: Implement
	return nil
}

// Time Complexity: O(m * n * 4^L) - m*n cells, L = max word length
// Space Complexity: O(total chars in all words)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Tries Practice =====")
	fmt.Println("Complete the TODO sections above")
}
