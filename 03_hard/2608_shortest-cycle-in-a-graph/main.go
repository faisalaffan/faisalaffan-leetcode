package main

// LeetCode #2608: Shortest Cycle in a Graph
// https://leetcode.com/problems/shortest-cycle-in-a-graph/
// Difficulty: Hard
//
// Approach: BFS from each unvisited node.
// For each node, BFS tracks parent to avoid going back.
// When we encounter a visited neighbor that is not parent,
// we found a cycle of length = dist[u] + dist[v] + 1.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1: n=7, edges=[[0,1],[1,2],[2,0],[3,4],[4,5],[5,6],[6,3]] -> 3
	fmt.Println(findShortestCycle(7, [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}}))
	// Example 2: n=4, edges=[[0,1],[0,2]] -> -1
	fmt.Println(findShortestCycle(4, [][]int{{0, 1}, {0, 2}}))
	// Example 3: n=5, edges=[[0,1],[1,2],[2,3],[3,1]] -> 3
	fmt.Println(findShortestCycle(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}}))
}

func findShortestCycle(n int, edges [][]int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := math.MaxInt32

	for start := 0; start < n; start++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		parent := make([]int, n)
		for i := range parent {
			parent[i] = -1
		}

		q := []int{start}
		dist[start] = 0

		for len(q) > 0 {
			u := q[0]
			q = q[1:]

			for _, v := range adj[u] {
				if v == parent[u] {
					continue
				}
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					parent[v] = u
					q = append(q, v)
				} else {
					// Cycle found
					cycleLen := dist[u] + dist[v] + 1
					if cycleLen < ans {
						ans = cycleLen
					}
				}
			}
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
