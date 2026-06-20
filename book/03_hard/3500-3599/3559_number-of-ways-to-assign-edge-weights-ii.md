# 3559 — Number Of Ways To Assign Edge Weights Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func waysToAssignEdgeWeights(n int, edges [][]int, m int) int
```

> **💡 Hint:** Tree DP. For each node, compute ways to assign weights to

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3559: Number of Ways to Assign Edge Weights II
// https://leetcode.com/problems/number-of-ways-to-assign-edge-weights-ii/
// Difficulty: Hard
//
// Given a tree, count the number of ways to assign edge weights (from 1..m)
// such that the distance between each pair of nodes satisfies constraints.
//
// Approach: Tree DP. For each node, compute ways to assign weights to
// its edges while respecting distance constraints.

import "fmt"

func main() {
	// Example 1
	fmt.Println(waysToAssignEdgeWeights(3, [][]int{{0, 1}, {1, 2}}, 2))
	// Example 2
	fmt.Println(waysToAssignEdgeWeights(2, [][]int{{0, 1}}, 3))
	// Edge: single node
	fmt.Println(waysToAssignEdgeWeights(1, [][]int{}, 5))
}

const MOD = 1000000007

func waysToAssignEdgeWeights(n int, edges [][]int, m int) int {
	if n <= 1 {
		return 1
	}

  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// DFS to compute DP
	var dfs func(u, parent int) int64
	dfs = func(u, parent int) int64 {
		ways := int64(1)
		childCount := 0
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			childCount++
			childWays := dfs(v, u)
			// For each child edge, we can assign any of m weights
			ways = (ways * childWays) % MOD
		}
		// For this node's edge to its parent, we have m choices
		if parent != -1 {
			ways = (ways * int64(m)) % MOD
		}
		return ways
	}

	return int(dfs(0, -1))
}
```
