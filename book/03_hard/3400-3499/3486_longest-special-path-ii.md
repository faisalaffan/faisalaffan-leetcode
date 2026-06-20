# 3486 — Longest Special Path Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSpecialPath(edges [][]int, nums []int) []int
```

> **💡 Hint:** DFS with hash set tracking edge weights used along the path.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS, Backtracking

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3486: Longest Special Path II
// https://leetcode.com/problems/longest-special-path-ii/
// Difficulty: Hard
//
// Given a tree with weighted edges and node values, find the longest special
// path where no edge weight appears more than once (unique weights).
//
// Approach: DFS with hash set tracking edge weights used along the path.
// Backtrack to explore all paths.

import "fmt"

func main() {
	// Example 1
	fmt.Println(longestSpecialPath([][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 2}}, []int{1, 2, 3, 4}))
	// Example 2: simple chain
	fmt.Println(longestSpecialPath([][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 4}}, []int{1, 1, 1, 1}))
	// Edge: single node
	fmt.Println(longestSpecialPath([][]int{}, []int{5}))
	// Edge: two nodes
	fmt.Println(longestSpecialPath([][]int{{0, 1, 10}}, []int{1, 2}))
}

func longestSpecialPath(edges [][]int, nums []int) []int {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return []int{0, 0}
	}

  // Membuat matriks/slice 2D untuk DP
	g := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, w})
	}

	maxLen := 0
	minNodes := 0

	// DFS from each node as start (since tree is small enough for brute force)
	var dfs func(u, parent int, used map[int]bool, pathLen int, nodeCount int)
	dfs = func(u, parent int, used map[int]bool, pathLen int, nodeCount int) {
		// Update answer
		if pathLen > maxLen || (pathLen == maxLen && nodeCount < minNodes) {
			maxLen = pathLen
			minNodes = nodeCount
		}

		for _, edge := range g[u] {
			v, w := edge[0], edge[1]
			if v == parent || used[w] {
				continue
			}
			used[w] = true
			dfs(v, u, used, pathLen+w, nodeCount+1)
			delete(used, w)
		}
	}

	// For each starting node, do DFS
	for start := 0; start < n; start++ {
  // Membuat map (HashMap) — pencarian O(1)
		used := make(map[int]bool)
		dfs(start, -1, used, 0, 1)
	}

	return []int{maxLen, minNodes}
}
```
