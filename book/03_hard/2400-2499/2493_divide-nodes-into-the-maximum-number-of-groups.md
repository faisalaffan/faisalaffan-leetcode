# 2493 — Divide Nodes Into The Maximum Number Of Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func magnificentSets(n int, edges [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2493: Divide Nodes Into the Maximum Number of Groups
// https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/
// Difficulty: Hard
//
// For each connected component:
//  1. Check if it's bipartite (BFS coloring). If not, return -1.
//  2. For each node in the component, BFS to find max distance.
//  3. Answer = sum of max distances across all components.

import "fmt"

func main() {
	// Example 1: n=6, edges=[[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]] => 4
	fmt.Println(magnificentSets(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	// Example 2: n=3, edges=[[1,2],[2,3],[3,1]] => -1 (triangle not bipartite)
	fmt.Println(magnificentSets(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Edge: single node
	fmt.Println(magnificentSets(1, [][]int{}))
	// Edge: two nodes
	fmt.Println(magnificentSets(2, [][]int{{1, 2}}))
	// Edge: star graph
	fmt.Println(magnificentSets(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}}))
}

func magnificentSets(n int, edges [][]int) int {
	// Build adjacency list (1-indexed)
  // Matriks 2D
	adj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []int{}
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	visited := make([]bool, n+1)
	total := 0

	for i := 1; i <= n; i++ {
		if !visited[i] {
			// Collect all nodes in this component (bfsCollect marks them visited)
			component := bfsCollect(i, adj, visited)

			// Check bipartite
  // HashMap: O(1) lookup
			color := make(map[int]int)
			if !isBipartite(component[0], adj, color) {
				return -1
			}

			// Find max depth in this component
			maxDepth := 0
			for _, node := range component {
				depth := bfsDepth(node, adj)
				if depth > maxDepth {
					maxDepth = depth
				}
			}
			total += maxDepth
		}
	}

	return total
}

func bfsCollect(start int, adj [][]int, visited []bool) []int {
	queue := []int{start}
	visited[start] = true
	result := []int{}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)
		for _, v := range adj[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return result
}

func isBipartite(start int, adj [][]int, color map[int]int) bool {
	queue := []int{start}
	color[start] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if c, ok := color[v]; ok {
				if c == color[u] {
					return false
				}
			} else {
				color[v] = color[u] ^ 1
				queue = append(queue, v)
			}
		}
	}
	return true
}

func bfsDepth(start int, adj [][]int) int {
  // HashMap: O(1) lookup
	dist := make(map[int]int)
	queue := []int{start}
	dist[start] = 1
	maxDist := 1

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if _, ok := dist[v]; !ok {
				dist[v] = dist[u] + 1
				if dist[v] > maxDist {
					maxDist = dist[v]
				}
				queue = append(queue, v)
			}
		}
	}
	return maxDist
}
```
