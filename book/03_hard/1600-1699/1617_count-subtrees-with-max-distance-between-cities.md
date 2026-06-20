# 1617 — Count Subtrees With Max Distance Between Cities

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func countSubgraphsForEachDiameter(n int, edges [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS, Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1617: Count Subtrees With Max Distance Between Cities
// https://leetcode.com/problems/count-subtrees-with-max-distance-between-cities/
// Difficulty: Hard

import "fmt"

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	// Floyd-Warshall for all-pairs shortest paths
	INF := 1000
  // Matriks 2D
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = INF
		}
		dist[i][i] = 0
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		dist[u][v] = 1
		dist[v][u] = 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

  // Alokasi slice
	result := make([]int, n-1) // diameters 1..n-1

	// Enumerate all non-empty subsets
	for mask := 1; mask < (1 << n); mask++ {
		// Collect nodes in this subset
  // Alokasi slice
		nodes := make([]int, 0, n)
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				nodes = append(nodes, i)
			}
		}
		if len(nodes) < 2 {
			continue
		}

		// Check connectivity: BFS on induced subgraph
  // HashMap: O(1) lookup
		visited := make(map[int]bool)
		queue := []int{nodes[0]}
		visited[nodes[0]] = true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, nb := range nodes {
				if !visited[nb] && dist[cur][nb] == 1 {
					visited[nb] = true
					queue = append(queue, nb)
				}
			}
		}
		if len(visited) != len(nodes) {
			continue // not connected
		}

		// Compute diameter (max distance between any two nodes in subset)
		maxDist := 0
  // Linear scan O(n)
		for i := 0; i < len(nodes); i++ {
			for j := i + 1; j < len(nodes); j++ {
				if dist[nodes[i]][nodes[j]] > maxDist {
					maxDist = dist[nodes[i]][nodes[j]]
				}
			}
		}

		if maxDist >= 1 && maxDist <= n-1 {
			result[maxDist-1]++
		}
	}

	return result
}

func main() {
	// Test case 1: n=4, edges=[[1,2],[2,3],[2,4]] -> [3,4,0]
	n := 4
	edges := [][]int{{1, 2}, {2, 3}, {2, 4}}
	result := countSubgraphsForEachDiameter(n, edges)
	fmt.Printf("n=%d edges=%v -> %v (expected [3,4,0])\n", n, edges, result)

	// Test case 2: n=2, edges=[[1,2]] -> [1]
	n2 := 2
	edges2 := [][]int{{1, 2}}
	result2 := countSubgraphsForEachDiameter(n2, edges2)
	fmt.Printf("n=%d edges=%v -> %v (expected [1])\n", n2, edges2, result2)

	// Test case 3: n=3, edges=[[1,2],[2,3]] -> [2,1]
	n3 := 3
	edges3 := [][]int{{1, 2}, {2, 3}}
	result3 := countSubgraphsForEachDiameter(n3, edges3)
	fmt.Printf("n=%d edges=%v -> %v (expected [2,1])\n", n3, edges3, result3)
}
```
