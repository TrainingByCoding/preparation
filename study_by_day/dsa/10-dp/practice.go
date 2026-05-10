package main

import "fmt"

// ========================================
// 1. Climbing Stairs
// ========================================
// n stairs, climb 1 or 2 at a time. How many distinct ways?
// dp[i] = dp[i-1] + dp[i-2]

func climbStairs(n int) int {
	// TODO: dp, or just two variables (Fibonacci style)
	return 0
}

// ========================================
// 2. House Robber
// ========================================
// Can't rob adjacent houses. Max money.
// dp[i] = max(dp[i-1], dp[i-2] + nums[i])

func rob(nums []int) int {
	// TODO: track prev2, prev1; iterate updating them
	return 0
}

// ========================================
// 3. Longest Common Subsequence
// ========================================
// LCS of "abcde" and "ace" = 3 ("ace")
// dp[i][j] = if chars match: dp[i-1][j-1]+1, else max(dp[i-1][j], dp[i][j-1])

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// TODO: fill dp table
	return dp[m][n]
}

// ========================================
// 4. 0/1 Knapsack
// ========================================
// weights[], values[], capacity W
// dp[i][w] = max value using first i items with capacity w

func knapsack(weights, values []int, W int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	// TODO: for each item i, for each capacity w:
	//   if weights[i-1] <= w: dp[i][w] = max(dp[i-1][w], values[i-1] + dp[i-1][w-weights[i-1]])
	//   else: dp[i][w] = dp[i-1][w]
	return dp[n][W]
}

// ========================================
// 5. Coin Change
// ========================================
// Min coins to make amount. coins=[1,5,11], amount=15 → 3 (5+5+5)
// dp[i] = min coins for amount i

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	} // init to "infinity"
	// TODO: for each amount i, for each coin:
	//   if coin <= i: dp[i] = min(dp[i], 1 + dp[i-coin])
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println("=== dsa/10: Dynamic Programming ===\n")

	fmt.Println("Climb stairs n=5:", climbStairs(5))                            // 8
	fmt.Println("House robber:", rob([]int{2, 7, 9, 3, 1}))                     // 12
	fmt.Println("LCS abcde/ace:", longestCommonSubsequence("abcde", "ace"))     // 3
	fmt.Println("Knapsack:", knapsack([]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 5)) // 7
	fmt.Println("Coin change:", coinChange([]int{1, 5, 11}, 15))                // 3
}

/*
SOLUTIONS:

func climbStairs(n int) int {
	if n <= 2 { return n }
	a, b := 1, 2
	for i := 3; i <= n; i++ { a, b = b, a+b }
	return b
}

func rob(nums []int) int {
	prev2, prev1 := 0, 0
	for _, n := range nums { prev2, prev1 = prev1, max(prev1, prev2+n) }
	return prev1
}
func max(a, b int) int { if a > b { return a }; return b }

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp { dp[i] = make([]int, n+1) }
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] { dp[i][j] = dp[i-1][j-1]+1 } else if dp[i-1][j] > dp[i][j-1] { dp[i][j] = dp[i-1][j] } else { dp[i][j] = dp[i][j-1] }
		}
	}
	return dp[m][n]
}

func knapsack(weights, values []int, W int) int {
	n := len(weights); dp := make([][]int, n+1)
	for i := range dp { dp[i] = make([]int, W+1) }
	for i := 1; i <= n; i++ {
		for w := 0; w <= W; w++ {
			dp[i][w] = dp[i-1][w]
			if weights[i-1] <= w {
				v := values[i-1] + dp[i-1][w-weights[i-1]]
				if v > dp[i][w] { dp[i][w] = v }
			}
		}
	}
	return dp[n][W]
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ { dp[i] = amount+1 }
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i && 1+dp[i-c] < dp[i] { dp[i] = 1+dp[i-c] }
		}
	}
	if dp[amount] > amount { return -1 }; return dp[amount]
}
*/
