package main

// LeetCode #2378: Choose Edges to Maximize Score in a Tree
// https://leetcode.com/problems/choose-edges-to-maximize-score-in-a-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Maximum weight matching on a tree (no two edges share a node).
// DP: dp0[u] = max when u is free, dp1[u] = max when edge(u,parent) is taken.
// dp1[u] = sum(dp0[child])
// dp0[u] = max(not matching any child, matching one child v: w + dp1[v] + sum_{other} dp0[other])

import "fmt"

func main() {
	fmt.Println(maxScore([][]int{{0, 1, 5}, {1, 2, 3}, {0, 3, 2}})) // 5
	fmt.Println(maxScore([][]int{{0, 1, 10}, {0, 2, 20}}))           // 20
	fmt.Println(maxScore([][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 3}})) // 4
}

func maxScore(edges [][]int) int64 {
	n := len(edges) + 1
	graph := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], [2]int{v, w})
		graph[v] = append(graph[v], [2]int{u, w})
	}

	var dfs func(u, parent int) (int64, int64)
	dfs = func(u, parent int) (int64, int64) {
		var base int64
		bestDelta := int64(0)

		for _, nei := range graph[u] {
			v, w := nei[0], int64(nei[1])
			if v == parent {
				continue
			}
			child0, child1 := dfs(v, u)
			base += child0
			delta := w + child1 - child0
			if delta > bestDelta {
				bestDelta = delta
			}
		}

		dp1 := base
		dp0 := base + bestDelta
		return dp0, dp1
	}

	dp0, _ := dfs(0, -1)
	return dp0
}
