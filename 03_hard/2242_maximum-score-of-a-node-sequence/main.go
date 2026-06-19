package main

// LeetCode #2242: Maximum Score of a Node Sequence
// https://leetcode.com/problems/maximum-score-of-a-node-sequence/
// Difficulty: Hard
//
// Given a graph with n nodes (0..n-1) with node scores, and edges (u, v).
// Find the maximum score of a node sequence [a, b, c, d] where
// edges exist between (a,b), (b,c), (c,d) and all four nodes are distinct.

import (
	"fmt"
	"sort"
)

// maximumScore returns maximum score of a valid 4-node sequence.
func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)

	// adjacency list of neighbors sorted by score descending (keep up to 3 best)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// For each node, keep its top 3 neighbors by score (to limit search space)
	top3 := make([][]int, n)
	for i := 0; i < n; i++ {
		neighbors := adj[i]
		sort.Slice(neighbors, func(a, b int) bool {
			return scores[neighbors[a]] > scores[neighbors[b]]
		})
		if len(neighbors) > 3 {
			neighbors = neighbors[:3]
		}
		top3[i] = neighbors
	}

	maxScore := -1

	// For each edge (b, c), try extending to a and d via top neighbors
	for _, e := range edges {
		b, c := e[0], e[1]

		// For node b, look at its top3 (excluding c)
		for _, a := range top3[b] {
			if a == c {
				continue
			}
			// For node c, look at its top3 (excluding b and a)
			for _, d := range top3[c] {
				if d == b || d == a {
					continue
				}
				score := scores[a] + scores[b] + scores[c] + scores[d]
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}

	return maxScore
}

func main() {
	// Example 1
	scores1 := []int{5, 2, 9, 8, 4}
	edges1 := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
	fmt.Println(maximumScore(scores1, edges1)) // Expected: 24

	// Example 2
	scores2 := []int{9, 10, 11, 12, 13, 14}
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(maximumScore(scores2, edges2)) // Expected: 46
}
