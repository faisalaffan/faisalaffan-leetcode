package main

// LeetCode #3367: Maximize Sum of Weights after Edge Removals
// https://leetcode.com/problems/maximize-sum-of-weights-after-edge-removals/
// Difficulty: Hard
//
// Given a tree with n nodes and weighted edges, remove edges so each node has
// at most k connections. Maximize sum of remaining edge weights.
//
// Approach: Tree DP. For each node, compute dp[node][0] = max sum in subtree
// when node has no parent edge (can use up to k child edges), and dp[node][1]
// = max sum when node is connected to parent (can use up to k-1 child edges).
// For each child, compute the gain of keeping its edge vs cutting it.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maximizeSumOfWeights([][]int{{0, 1, 4}, {0, 2, 2}, {2, 3, 12}, {2, 4, 6}}, 2))
	// Example 2
	fmt.Println(maximizeSumOfWeights([][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 7}}, 2))
	// Example 3: single edge
	fmt.Println(maximizeSumOfWeights([][]int{{0, 1, 10}}, 1))
	// Edge: k = 0
	fmt.Println(maximizeSumOfWeights([][]int{{0, 1, 5}, {0, 2, 3}}, 0))
	// Edge: star tree
	fmt.Println(maximizeSumOfWeights([][]int{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}}, 2))
}

func maximizeSumOfWeights(edges [][]int, k int) int64 {
	n := len(edges) + 1
	g := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, w})
	}

	var dfs func(u, parent int) (int64, int64)
	dfs = func(u, parent int) (int64, int64) {
		var base int64
		var gains []int64

		for _, edge := range g[u] {
			v, w := edge[0], int64(edge[1])
			if v == parent {
				continue
			}
			dp0, dp1 := dfs(v, u)
			base += dp0
			gain := w + dp1 - dp0
			if gain > 0 {
				gains = append(gains, gain)
			}
		}

		sort.Slice(gains, func(i, j int) bool {
			return gains[i] > gains[j]
		})

		// dp0: can use up to k child edges (root or cut edge to parent)
		dp0 := base
		for i := 0; i < k && i < len(gains); i++ {
			dp0 += gains[i]
		}

		// dp1: can use up to k-1 child edges (connected to parent)
		dp1 := base
		for i := 0; i < k-1 && i < len(gains); i++ {
			dp1 += gains[i]
		}

		return dp0, dp1
	}

	ans, _ := dfs(0, -1)
	return ans
}
