# 3887 — Incremental Even Weighted Cycle Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func incrementalEvenWeightedCycleQueries(n int, edges [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Union-Find

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Union-Find** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3887: Incremental Even-Weighted Cycle Queries
// https://leetcode.com/problems/incremental-even-weighted-cycle-queries/
// Difficulty: Hard
//
// Given a connected, undirected graph with n nodes and weighted
// edges, process incremental edge additions. After each addition,
// count the number of cycles in the graph where the sum of edge
// weights in the cycle is even.
//
// Approach: Maintain a DSU with parity tracking. When adding an
// edge (u, v, w), if u and v are already connected, the edge forms
// a cycle. Track parity of path from u to root.

import "fmt"

func main() {
	// Example 1
	fmt.Println(incrementalEvenWeightedCycleQueries(4, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 3, 3}, {0, 3, 2}}))
	// Example 2
	fmt.Println(incrementalEvenWeightedCycleQueries(3, [][]int{{0, 1, 1}, {1, 2, 2}, {0, 2, 3}}))
	// Edge: single edge
	fmt.Println(incrementalEvenWeightedCycleQueries(2, [][]int{{0, 1, 5}}))
}

func incrementalEvenWeightedCycleQueries(n int, edges [][]int) int {
  // Alokasi slice
	parent := make([]int, n)
  // Alokasi slice
	rank := make([]int, n)
  // Alokasi slice
	xorToRoot := make([]int, n) // parity of path weight to root

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			orig := parent[x]
			parent[x] = find(parent[x])
			xorToRoot[x] ^= xorToRoot[orig]
		}
		return parent[x]
	}

	union := func(u, v, w int) bool {
		pu, pv := find(u), find(v)
		xorUV := xorToRoot[u] ^ xorToRoot[v] ^ w

		if pu == pv {
			return xorUV == 0 // forms an even-weight cycle
		}

		if rank[pu] < rank[pv] {
			pu, pv = pv, pu
			u, v = v, u
			xorUV = xorToRoot[u] ^ xorToRoot[v] ^ w
		}
		parent[pv] = pu
		xorToRoot[pv] = xorUV
		if rank[pu] == rank[pv] {
			rank[pu]++
		}
		return false
	}

	evenCycles := 0
	for _, e := range edges {
		if union(e[0], e[1], e[2]) {
			evenCycles++
		}
	}
	return evenCycles
}
```
