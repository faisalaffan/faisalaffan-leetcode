package main

// LeetCode #2312: Selling Pieces of Wood
// https://leetcode.com/problems/selling-pieces-of-wood/
// Difficulty: Hard
//
// Approach: DP. dp[h][w] = max profit for a piece of size h x w.
// Initialize dp with 0, then apply given prices. For each piece, try all
// horizontal cuts (split into h1 x w and (h-h1) x w) and vertical cuts
// (split into h x w1 and h x (w-w1)). Take max.
// m <= 200, n <= 200, so O(m*n*(m+n)) is fine.

import "fmt"

func main() {
	// Example 1: m=3, n=5, prices=[[1,4,2],[2,2,7],[2,1,3]] => 19
	fmt.Println(sellingWood(3, 5, [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}}))
	// Example 2: m=4, n=6, prices=[[3,2,10],[1,4,2],[4,1,3]] => 32
	fmt.Println(sellingWood(4, 6, [][]int{{3, 2, 10}, {1, 4, 2}, {4, 1, 3}}))
	// Edge: no prices
	fmt.Println(sellingWood(2, 2, [][]int{}))
	// Edge: single price
	fmt.Println(sellingWood(2, 2, [][]int{{2, 2, 5}}))
	// Edge: 1x1
	fmt.Println(sellingWood(1, 1, [][]int{{1, 1, 10}}))
}

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}

	// Apply given prices
	for _, p := range prices {
		h, w, pr := p[0], p[1], p[2]
		if h <= m && w <= n {
			if int64(pr) > dp[h][w] {
				dp[h][w] = int64(pr)
			}
		}
	}

	// DP: try all cuts
	for h := 1; h <= m; h++ {
		for w := 1; w <= n; w++ {
			// Horizontal cuts: split into h1 x w and (h-h1) x w
			for h1 := 1; h1 <= h/2; h1++ {
				val := dp[h1][w] + dp[h-h1][w]
				if val > dp[h][w] {
					dp[h][w] = val
				}
			}
			// Vertical cuts: split into h x w1 and h x (w-w1)
			for w1 := 1; w1 <= w/2; w1++ {
				val := dp[h][w1] + dp[h][w-w1]
				if val > dp[h][w] {
					dp[h][w] = val
				}
			}
		}
	}
	return dp[m][n]
}
