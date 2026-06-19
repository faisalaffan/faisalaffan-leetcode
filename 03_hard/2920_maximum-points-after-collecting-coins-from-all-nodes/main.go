package main

// LeetCode #2920: Maximum Points After Collecting Coins From All Nodes
// https://leetcode.com/problems/maximum-points-after-collecting-coins-from-all-nodes/
// Difficulty: Hard
//
// Approach: Tree DP with memoization.
// At each node, we can either:
//   1. Collect coins[i] - k points (penalty), children unaffected.
//   2. Collect floor(coins[i]/2) points, and all descendants' coins are halved.
// Since coins[i] <= 10^4, at most 14 halvings reduce everything to 0.
// DP[node][shifts] = max points from the subtree when coins have been halved `shifts` times.

import "fmt"

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	// dp[node][shifts] = max points from subtree at node with `shifts` halvings applied
	const maxShifts = 14
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, maxShifts)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(node, parent, shifts int) int
	dfs = func(node, parent, shifts int) int {
		if shifts >= maxShifts {
			return 0
		}
		if dp[node][shifts] != -1 {
			return dp[node][shifts]
		}

		// Option 1: collect with penalty k
		way1 := (coins[node] >> shifts) - k
		// Option 2: halve this node and the entire subtree
		way2 := coins[node] >> (shifts + 1)

		for _, child := range g[node] {
			if child == parent {
				continue
			}
			way1 += dfs(child, node, shifts)
			way2 += dfs(child, node, shifts+1)
		}

		if way1 > way2 {
			dp[node][shifts] = way1
		} else {
			dp[node][shifts] = way2
		}
		return dp[node][shifts]
	}

	return dfs(0, -1, 0)
}

func main() {
	// Example 1: edges=[[0,1],[1,2],[2,3]], coins=[10,10,3,3], k=2
	// Optimal: collect penalty at 0,1 and halve at 2,3 (or all penalty)
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}, {2, 3}}, []int{10, 10, 3, 3}, 2))

	// Example from LeetCode: k=5 -> 11
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}, {2, 3}}, []int{10, 10, 3, 3}, 5))

	// Example 2: star tree, k=0 -> 16
	fmt.Println(maximumPoints([][]int{{0, 1}, {0, 2}}, []int{8, 4, 4}, 0))
}
