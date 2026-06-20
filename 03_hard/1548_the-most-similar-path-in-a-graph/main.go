package main

// LeetCode #1548: The Most Similar Path in a Graph
// https://leetcode.com/problems/the-most-similar-path-in-a-graph/
// Difficulty: Hard [Paid]
//
// Given n cities connected by roads, each with a name, find a path of
// length equal to targetPath that minimizes the edit distance (number
// of mismatches) between city names and targetPath names.
//
// Approach: DP[i][v] = min edit distance for first i steps ending at city v.
// Reconstruct path by backtracking through DP.

import "fmt"

func main() {
	// Example 1
	fmt.Println(mostSimilar(5, [][]int{{0, 2}, {0, 3}, {1, 2}, {1, 3}, {1, 4}, {2, 4}},
		[]string{"ATL", "PEK", "LAX", "DXB", "HND"},
		[]string{"ATL", "DXB", "HND", "LAX"}))
	// Example 2
	fmt.Println(mostSimilar(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
		[]string{"BOM", "BOM", "MAA", "BOM"},
		[]string{"BOM", "MAA", "BOM"}))
	// Edge: single node
	fmt.Println(mostSimilar(1, [][]int{},
		[]string{"A"},
		[]string{"A", "B", "A"}))
}

func mostSimilar(n int, roads [][]int, names []string, targetPath []string) []int {
	m := len(targetPath)
	adj := make([][]int, n)
	for _, r := range roads {
		u, v := r[0], r[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// If no edges, handle separately
	if n == 1 {
		path := make([]int, m)
		for i := range path {
			path[i] = 0
		}
		return path
	}

	// dp[i][v] = min edit distance for first i steps (0-indexed) ending at v
	dp := make([][]int, m)
	prev := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		prev[i] = make([]int, n)
		for v := 0; v < n; v++ {
			cost := 0
			if names[v] != targetPath[i] {
				cost = 1
			}
			if i == 0 {
				dp[i][v] = cost
				prev[i][v] = -1
			} else {
				best := m + 1
				bestPrev := -1
				for _, u := range adj[v] {
					if dp[i-1][u] < best {
						best = dp[i-1][u]
						bestPrev = u
					}
				}
				dp[i][v] = best + cost
				prev[i][v] = bestPrev
			}
		}
	}

	// Find best ending city
	end := 0
	best := dp[m-1][0]
	for v := 1; v < n; v++ {
		if dp[m-1][v] < best {
			best = dp[m-1][v]
			end = v
		}
	}

	// Reconstruct path
	path := make([]int, m)
	path[m-1] = end
	for i := m - 1; i > 0; i-- {
		path[i-1] = prev[i][path[i]]
	}
	return path
}
