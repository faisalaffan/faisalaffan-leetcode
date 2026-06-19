package main

// LeetCode #1478: Allocate Mailboxes
// https://leetcode.com/problems/allocate-mailboxes/
// Difficulty: Hard
//
// Approach: DP + Median Cost
// Sort houses first. dp[i][j] = min distance to place j mailboxes
// among first i houses (0-indexed).
// cost[l][r] = min total distance to serve houses[l..r] with 1 mailbox
// placed at the median (optimal for minimizing sum of absolute distances).
// Transition: dp[i][j] = min over p < i of dp[p][j-1] + cost[p+1][i].

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minDistance([]int{1, 4, 8, 10, 20}, 3))
	// Expected: 5

	fmt.Println(minDistance([]int{2, 3, 5, 12, 18}, 2))
	// Expected: 9

	fmt.Println(minDistance([]int{7, 4, 6, 1}, 1))
	// Expected: 8
}

func minDistance(houses []int, k int) int {
	sort.Ints(houses)
	n := len(houses)

	if k >= n {
		return 0
	}

	// precompute cost[i][j] = min dist for 1 mailbox serving houses[i..j]
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// Median index
			mid := i + (j-i)/2
			median := houses[mid]
			total := 0
			for t := i; t <= j; t++ {
				total += absInt(houses[t] - median)
			}
			cost[i][j] = total
		}
	}

	// dp[i][j] = min distance for first i+1 houses with j+1 mailboxes
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// base: 1 mailbox
	for i := 0; i < n; i++ {
		dp[i][0] = cost[0][i]
	}

	// fill dp
	for j := 1; j < k; j++ {
		for i := j; i < n; i++ {
			for p := j - 1; p < i; p++ {
				dp[i][j] = minInt(dp[i][j], dp[p][j-1]+cost[p+1][i])
			}
		}
	}

	return dp[n-1][k-1]
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
