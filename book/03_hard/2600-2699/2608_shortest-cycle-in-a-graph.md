# 2608 — Shortest Cycle In A Graph

## Deskripsi

**Soal:** [2608. Shortest Cycle In A Graph](https://leetcode.com/problems/shortest-cycle-in-a-graph/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func findShortestCycle(n int, edges [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2608: Shortest Cycle in a Graph
// https://leetcode.com/problems/shortest-cycle-in-a-graph/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

// findShortestCycle returns the length of the shortest cycle in an undirected graph.
// BFS from each unvisited node. For each node, BFS tracks parent to avoid going back.
// When we encounter a visited neighbor that is not parent, we found a cycle.
//
// Complexity: O(n * (n+m)) time, O(n+m) space
func findShortestCycle(n int, edges [][]int) int {
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := math.MaxInt32

	for start := 0; start < n; start++ {
  // Membuat slice untuk menyimpan hasil
		dist := make([]int, n)
  // Iterasi seluruh elemen
		for i := range dist {
			dist[i] = -1
		}
  // Membuat slice untuk menyimpan hasil
		parent := make([]int, n)
  // Iterasi seluruh elemen
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

func main() {
	// Example 1: n=7, edges=[[0,1],[1,2],[2,0],[3,4],[4,5],[5,6],[6,3]] -> 3
	fmt.Println("Test 1: ->", findShortestCycle(7, [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}}))
	// Example 2: n=4, edges=[[0,1],[0,2]] -> -1
	fmt.Println("Test 2: ->", findShortestCycle(4, [][]int{{0, 1}, {0, 2}}))
	// Example 3: n=5, edges=[[0,1],[1,2],[2,3],[3,1]] -> 3
	fmt.Println("Test 3: ->", findShortestCycle(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}}))
	// Simple triangle
	fmt.Println("Test 4: triangle ->", findShortestCycle(3, [][]int{{0, 1}, {1, 2}, {2, 0}}))
	// No edges
	fmt.Println("Test 5: no edges ->", findShortestCycle(3, [][]int{}))
	// 4-cycle
	fmt.Println("Test 6: square ->", findShortestCycle(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}))
}
```
