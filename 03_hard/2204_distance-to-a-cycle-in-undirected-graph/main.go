package main

// LeetCode #2204: Distance to a Cycle in Undirected Graph
// https://leetcode.com/problems/distance-to-a-cycle-in-undirected-graph/
// Difficulty: Hard
//
// Topo remove leaves + BFS:
// 1. Compute degrees. Repeatedly remove leaf nodes (degree == 1), marking them
//    as removed (deg = 0).
// 2. Remaining nodes (deg > 0) are cycle nodes. BFS from all cycle nodes to
//    compute distances.

import (
	"fmt"
)

func main() {
	n := 7
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {0, 1}, {5, 2}, {6, 5}}
	// Expected: [1 0 0 0 0 1 2]
	fmt.Println(distanceToCycle(n, edges))

	// Simple triangle
	n2 := 3
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 0}}
	fmt.Println(distanceToCycle(n2, edges2))

	// Single cycle with extra leaves
	n3 := 5
	edges3 := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 4}}
	fmt.Println(distanceToCycle(n3, edges3))
}

func distanceToCycle(n int, edges [][]int) []int {
	adj := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		deg[u]++
		deg[v]++
	}

	// Remove leaves: set deg = 0 to mark removed.
	q := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			q = append(q, i)
			deg[i] = 0
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			if deg[v] > 0 {
				deg[v]--
				if deg[v] == 1 {
					q = append(q, v)
					deg[v] = 0
				}
			}
		}
	}

	// BFS from all cycle nodes (deg > 0).
	ans := make([]int, n)
	queue := make([]int, 0, n)
	inQueue := make([]bool, n)
	for i := 0; i < n; i++ {
		if deg[i] > 0 {
			queue = append(queue, i)
			inQueue[i] = true
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if !inQueue[v] {
				ans[v] = ans[u] + 1
				inQueue[v] = true
				queue = append(queue, v)
			}
		}
	}

	return ans
}
