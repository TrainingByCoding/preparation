# NeetCode 150 - 
## Problem Categories and Distribution

| no | Category | Problems | Difficulty Distribution | Key Patterns |
|----|----------|----------|-------------------------|--------------|
|1| **Arrays & Hashing** | 9 | 2 Easy, 6 Medium, 1 Hard | Two Pointers, Hash Maps, Prefix Sums |
|2| **Two Pointers** | 5 | 1 Easy, 4 Medium | Left-Right Pointers, Fast-Slow Pointers |
|3| **Sliding Window** | 6 | 2 Easy, 3 Medium, 1 Hard | Fixed/Variable Window, String Problems |
|4| **Stack** | 7 | 2 Easy, 4 Medium, 1 Hard | Monotonic Stack, Expression Parsing |
|5| **Binary Search** | 7 | 2 Easy, 5 Medium | Search Space Reduction, Peak Finding |
|6| **Linked List** | 11 | 2 Easy, 8 Medium, 1 Hard | Two Pointers, Reversal, Cycle Detection |
|7| **Trees** | 15 | 3 Easy, 10 Medium, 2 Hard | DFS, BFS, Tree Construction |
|8| **Tries** | 3 | 2 Medium, 1 Hard | Prefix Trees, Word Search |
|9| **Heap/Priority Queue** | 7 | 1 Easy, 5 Medium, 1 Hard | K-Elements, Merge Operations |
|10| **Backtracking** | 9 | 6 Medium, 3 Hard | Permutations, Combinations, Subsets |
|11| **Graphs** | 13 | 3 Medium, 10 Hard | DFS, BFS, Union Find, Topological Sort |
|12| **Advanced Graphs** | 6 | 6 Hard | Shortest Path, MST, Advanced Algorithms |
|13| **1-D Dynamic Programming** | 12 | 1 Easy, 8 Medium, 3 Hard | State Transitions, Optimization |
|14| **2-D Dynamic Programming** | 11 | 8 Medium, 3 Hard | Grid DP, String DP |
|15| **Greedy** | 8 | 2 Easy, 4 Medium, 2 Hard | Local Optimal Choices |
|16| **Intervals** | 5 | 1 Easy, 3 Medium, 1 Hard | Merging, Scheduling |
|17| **Math & Geometry** | 8 | 3 Easy, 4 Medium, 1 Hard | Number Theory, Computational Geometry |
|18| **Bit Manipulation** | 7 | 3 Easy, 4 Medium | Bitwise Operations, XOR Tricks |

---

## 📋 Complete Problem List

### Arrays & Hashing (9 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 1 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Easy | Hash Set |
| 2 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Easy | Character Count |
| 3 | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy | Hash Map |
| 4 | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Medium | Hash Map, Sorting | 
| 5 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Medium | Heap, Hash Map |
| 6 | [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) | Medium | Prefix/Suffix Product | 
| 7 | [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/) | Medium | Hash Set, Matrix | 
| 8 | [Encode and Decode Strings](https://leetcode.com/problems/encode-and-decode-strings/) | Medium | String Manipulation | 
| 9 | [Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/) | Hard | Union Find, Hash Set | 

### Two Pointers (5 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 10 | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | Easy | Two Pointers | 
| 11 | [Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | Medium | Two Pointers | 
| 12 | [3Sum](https://leetcode.com/problems/3sum/) | Medium | Two Pointers, Sort |
| 13 | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | Medium | Two Pointers |
| 14 | [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | Hard | Two Pointers | 

### Sliding Window (6 Problems)

| No. | Problem | Difficulty | Pattern | 
|-----|---------|------------|---------|
| 15 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | Easy | Single Pass |
| 16 | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | Medium | Variable Window |
| 17 | [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/) | Medium | Variable Window | 
| 18 | [Permutation in String](https://leetcode.com/problems/permutation-in-string/) | Medium | Fixed Window | 
| 19 | [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | Hard | Variable Window | 
| 20 | [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) | Hard | Deque, Fixed Window | 

### Stack (7 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 21 | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | Easy | Stack Matching |
| 22 | [Min Stack](https://leetcode.com/problems/min-stack/) | Medium | Stack Design | 
| 23 | [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | Medium | Stack Evaluation | 
| 24 | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) | Medium | Backtracking, Stack | 
| 25 | [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/) | Medium | Monotonic Stack | 
| 26 | [Car Fleet](https://leetcode.com/problems/car-fleet/) | Medium | Stack, Sorting | 
| 27 | [Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/) | Hard | Monotonic Stack | 

### Binary Search (7 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 28 | [Binary Search](https://leetcode.com/problems/binary-search/) | Easy | Basic Binary Search |
| 29 | [Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/) | Medium | Matrix Binary Search | 
| 30 | [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/) | Medium | Binary Search on Answer | 
| 31 | [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | Medium | Rotated Array |
| 32 | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | Medium | Rotated Array |
| 33 | [Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/) | Medium | Binary Search, Design | 
| 34 | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | Hard | Binary Search, Arrays | 

### Linked List (11 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 35 | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | Easy | Iterative/Recursive |
| 36 | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | Easy | Two Pointers |
| 37 | [Reorder List](https://leetcode.com/problems/reorder-list/) | Medium | Multiple Pointers | 
| 38 | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | Medium | Two Pointers | 
| 39 | [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/) | Medium | Hash Map, Cloning | 
| 40 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | Medium | Simulation | 
| 41 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | Easy | Floyd's Algorithm |
| 42 | [Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/) | Medium | Floyd's Algorithm |
| 43 | [LRU Cache](https://leetcode.com/problems/lru-cache/) | Medium | Design, Hash Map |
| 44 | [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) | Hard | Divide & Conquer |
| 45 | [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/) | Hard | Advanced Reversal | 

### Trees (15 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 46 | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) | Easy | DFS/BFS |
| 47 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | Easy | DFS/BFS |
| 48 | [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) | Easy | DFS |
| 49 | [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/) | Easy | DFS | 
| 50 | [Same Tree](https://leetcode.com/problems/same-tree/) | Easy | DFS/BFS |
| 51 | [Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/) | Easy | DFS | 
| 52 | [Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | Medium | BST Property |
| 53 | [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) | Medium | BFS |
| 54 | [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/) | Medium | BFS/DFS | 
| 55 | [Count Good Nodes in Binary Tree](https://leetcode.com/problems/count-good-nodes-in-binary-tree/) | Medium | DFS | 
| 56 | [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) | Medium | DFS, BST |
| 57 | [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) | Medium | Inorder Traversal |
| 58 | [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) | Medium | Tree Construction |
| 59 | [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) | Hard | DFS, Path Sum |
| 60 | [Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) | Hard | Tree Serialization |

### Tries (3 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 61 | [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) | Medium | Trie Design |
| 62 | [Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-word-data-structure-design/) | Medium | Trie, Backtracking | 
| 63 | [Word Search II](https://leetcode.com/problems/word-search-ii/) | Hard | Trie, Backtracking |

### Heap / Priority Queue (7 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 64 | [Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | Easy | Min Heap |
| 65 | [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/) | Easy | Max Heap | 
| 66 | [K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/) | Medium | Min Heap |
| 67 | [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) | Medium | Quickselect, Heap |
| 68 | [Task Scheduler](https://leetcode.com/problems/task-scheduler/) | Medium | Max Heap, Greedy |
| 69 | [Design Twitter](https://leetcode.com/problems/design-twitter/) | Medium | Heap, Design | 
| 70 | [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) | Hard | Two Heaps |

### Backtracking (9 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 71 | [Subsets](https://leetcode.com/problems/subsets/) | Medium | Backtracking |
| 72 | [Combination Sum](https://leetcode.com/problems/combination-sum/) | Medium | Backtracking |
| 73 | [Permutations](https://leetcode.com/problems/permutations/) | Medium | Backtracking |
| 74 | [Subsets II](https://leetcode.com/problems/subsets-ii/) | Medium | Backtracking, Duplicates | 
| 75 | [Combination Sum II](https://leetcode.com/problems/combination-sum-ii/) | Medium | Backtracking, Duplicates | 
| 76 | [Word Search](https://leetcode.com/problems/word-search/) | Medium | Backtracking, Matrix |
| 77 | [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/) | Medium | Backtracking, String | 
| 78 | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | Medium | Backtracking |
| 79 | [N-Queens](https://leetcode.com/problems/n-queens/) | Hard | Backtracking | 

### Graphs (13 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 80 | [Number of Islands](https://leetcode.com/problems/number-of-islands/) | Medium | DFS/BFS |
| 81 | [Clone Graph](https://leetcode.com/problems/clone-graph/) | Medium | DFS/BFS, Cloning |
| 82 | [Max Area of Island](https://leetcode.com/problems/max-area-of-island/) | Medium | DFS | 
| 83 | [Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/) | Medium | DFS/BFS | 
| 84 | [Surrounded Regions](https://leetcode.com/problems/surrounded-regions/) | Medium | DFS/BFS | 
| 85 | [Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) | Medium | BFS | 
| 86 | [Walls and Gates](https://leetcode.com/problems/walls-and-gates/) | Medium | BFS | 
| 87 | [Course Schedule](https://leetcode.com/problems/course-schedule/) | Medium | Topological Sort |
| 88 | [Course Schedule II](https://leetcode.com/problems/course-schedule-ii/) | Medium | Topological Sort |
| 89 | [Redundant Connection](https://leetcode.com/problems/redundant-connection/) | Medium | Union Find | 
| 90 | [Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/) | Medium | Union Find/DFS | 
| 91 | [Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/) | Medium | Union Find/DFS | 
| 92 | [Word Ladder](https://leetcode.com/problems/word-ladder/) | Hard | BFS |

### Advanced Graphs (6 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 93 | [Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/) | Hard | Eulerian Path | 
| 94 | [Min Cost to Connect All Points](https://leetcode.com/problems/min-cost-to-connect-all-points/) | Medium | MST (Prim's/Kruskal's) | 
| 95 | [Network Delay Time](https://leetcode.com/problems/network-delay-time/) | Medium | Dijkstra's Algorithm | 
| 96 | [Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/) | Hard | Dijkstra's/Binary Search | 
| 97 | [Alien Dictionary](https://leetcode.com/problems/alien-dictionary/) | Hard | Topological Sort | 
| 98 | [Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/) | Medium | Bellman-Ford/Dijkstra's | 

### 1-D Dynamic Programming (12 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 99 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | Easy | Basic DP |
| 100 | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/) | Easy | Basic DP | 
| 101 | [House Robber](https://leetcode.com/problems/house-robber/) | Medium | Linear DP |
| 102 | [House Robber II](https://leetcode.com/problems/house-robber-ii/) | Medium | Circular Array DP |
| 103 | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | Medium | String DP |
| 104 | [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) | Medium | String DP | 
| 105 | [Decode Ways](https://leetcode.com/problems/decode-ways/) | Medium | String DP |
| 106 | [Coin Change](https://leetcode.com/problems/coin-change/) | Medium | DP, Greedy |
| 107 | [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/) | Medium | DP |
| 108 | [Word Break](https://leetcode.com/problems/word-break/) | Medium | String DP |
| 109 | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | Medium | DP, Binary Search |
| 110 | [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/) | Medium | Subset Sum DP | 

### 2-D Dynamic Programming (11 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 111 | [Unique Paths](https://leetcode.com/problems/unique-paths/) | Medium | Grid DP |
| 112 | [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/) | Medium | String DP |
| 113 | [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | Medium | State Machine DP | 
| 114 | [Coin Change 2](https://leetcode.com/problems/coin-change-2/) | Medium | Unbounded Knapsack | 
| 115 | [Target Sum](https://leetcode.com/problems/target-sum/) | Medium | Subset Sum DP | 
| 116 | [Interleaving String](https://leetcode.com/problems/interleaving-string/) | Medium | String DP | 
| 117 | [Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) | Hard | DFS + Memoization | 
| 118 | [Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/) | Hard | String DP | 
| 119 | [Edit Distance](https://leetcode.com/problems/edit-distance/) | Hard | String DP | 
| 120 | [Burst Balloons](https://leetcode.com/problems/burst-balloons/) | Hard | Interval DP | 
| 121 | [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/) | Hard | String DP |

### Greedy (8 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 122 | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | Medium | Kadane's Algorithm |
| 123 | [Jump Game](https://leetcode.com/problems/jump-game/) | Medium | Greedy |
| 124 | [Jump Game II](https://leetcode.com/problems/jump-game-ii/) | Medium | Greedy |
| 125 | [Gas Station](https://leetcode.com/problems/gas-station/) | Medium | Greedy | 
| 126 | [Hand of Straights](https://leetcode.com/problems/hand-of-straights/) | Medium | Greedy, Hash Map | 
| 127 | [Merge Triplets to Form Target Triplet](https://leetcode.com/problems/merge-triplets-to-form-target-triplet/) | Medium | Greedy | 
| 128 | [Partition Labels](https://leetcode.com/problems/partition-labels/) | Medium | Greedy | 
| 129 | [Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/) | Medium | Greedy | 

### Intervals (5 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 130 | [Insert Interval](https://leetcode.com/problems/insert-interval/) | Medium | Interval Merge |
| 131 | [Merge Intervals](https://leetcode.com/problems/merge-intervals/) | Medium | Interval Merge |
| 132 | [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) | Medium | Greedy | 
| 133 | [Meeting Rooms](https://leetcode.com/problems/meeting-rooms/) | Easy | Sorting |
| 134 | [Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/) | Medium | Heap, Greedy |

### Math & Geometry (8 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 135 | [Rotate Image](https://leetcode.com/problems/rotate-image/) | Medium | Matrix Manipulation |
| 136 | [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/) | Medium | Matrix Traversal |
| 137 | [Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/) | Medium | Matrix Modification | 
| 138 | [Happy Number](https://leetcode.com/problems/happy-number/) | Easy | Number Theory | Google, Apple |
| 139 | [Plus One](https://leetcode.com/problems/plus-one/) | Easy | Array Manipulation | Google, Apple |
| 140 | [Pow(x, n)](https://leetcode.com/problems/powx-n/) | Medium | Fast Exponentiation |
| 141 | [Multiply Strings](https://leetcode.com/problems/multiply-strings/) | Medium | String Math | 
| 142 | [Detect Squares](https://leetcode.com/problems/detect-squares/) | Medium | Geometry | 

### Bit Manipulation (7 Problems)

| No. | Problem | Difficulty | Pattern |
|-----|---------|------------|---------|
| 143 | [Single Number](https://leetcode.com/problems/single-number/) | Easy | XOR |
| 144 | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) | Easy | Bit Counting |
| 145 | [Counting Bits](https://leetcode.com/problems/counting-bits/) | Easy | DP + Bits |
| 146 | [Reverse Bits](https://leetcode.com/problems/reverse-bits/) | Easy | Bit Reversal | 
| 147 | [Missing Number](https://leetcode.com/problems/missing-number/) | Easy | XOR/Math |
| 148 | [Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/) | Medium | Bit Arithmetic | 
| 149 | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | Medium | Integer Manipulation |

---
