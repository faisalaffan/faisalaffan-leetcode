# 0785 — Is Graph Bipartite

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func isBipartite(graph [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #785: Is Graph Bipartite?
// https://leetcode.com/problems/is-graph-bipartite/
// Difficulty: Medium
// Time: O(V + E)
// Space: O(V)

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
  // Alokasi slice integer
	color := make([]int, n)

	var dfs func(node int, c int) bool
	dfs = func(node int, c int) bool {
		if color[node] != 0 {
			return color[node] == c
		}
		color[node] = c

		for _, neighbor := range graph[node] {
			if !dfs(neighbor, -c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if color[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}
```
