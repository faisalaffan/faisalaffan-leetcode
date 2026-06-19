package main

// LeetCode #3547: Maximum Sum of Edge Values in a Graph
// https://leetcode.com/problems/maximum-sum-of-edge-values-in-a-graph/
// Difficulty: Hard
//
// Given a graph with edge values, maximize the sum of selected edges such
// that each node has at most one selected incident edge (matching).
//
// Approach: Maximum weight matching in a general graph. For bipartite graphs,
// use Hungarian algorithm. For small n, use DP over subsets.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maximumSumOfEdgeValues(3, [][]int{{0, 1, 5}, {1, 2, 3}, {0, 2, 4}}))
	// Example 2
	fmt.Println(maximumSumOfEdgeValues(2, [][]int{{0, 1, 10}}))
	// Edge: single edge
	fmt.Println(maximumSumOfEdgeValues(2, [][]int{{0, 1, 7}}))
}

func maximumSumOfEdgeValues(n int, edges [][]int) int64 {
	// Build adjacency with weights
	type edge struct {
		v int
		w int64
	}
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], int64(e[2])
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}

	// DP over subsets for maximum weight matching
	m := 1 << n
	dp := make([]int64, m)
	for i := 1; i < m; i++ {
		dp[i] = -1
	}

	for mask := 0; mask < m; mask++ {
		if dp[mask] < 0 {
			continue
		}
		// Find first unpaired node
		u := 0
		for u < n && (mask&(1<<u)) != 0 {
			u++
		}
		if u >= n {
			continue
		}
		// Skip this node (leave it unpaired)
		newMask := mask | (1 << u)
		if dp[mask] > dp[newMask] {
			dp[newMask] = dp[mask]
		}
		// Pair u with any available v
		for _, e := range adj[u] {
			v := e.v
			if mask&(1<<v) == 0 {
				newMask2 := newMask | (1 << v)
				if dp[mask]+e.w > dp[newMask2] {
					dp[newMask2] = dp[mask] + e.w
				}
			}
		}
	}

	return dp[m-1]
}
