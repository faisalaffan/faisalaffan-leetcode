# 2077 — Paths In Maze That Lead To Same Room

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfPaths(corridors [][]int, n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * deg^2)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2077: Paths in Maze That Lead to Same Room
// https://leetcode.com/problems/paths-in-maze-that-lead-to-same-room/
// Difficulty: Medium [Paid]
// Time: O(n * deg^2) | Space: O(n + m)

import "fmt"

func numberOfPaths(corridors [][]int, n int) int {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n+1)
	for _, c := range corridors {
		u, v := c[0], c[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// For each pair of neighbors of a node, check if they are also connected
	// Use adjacency set for O(1) lookup
  // Alokasi slice integer
	adjSet := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		adjSet[i] = make(map[int]bool)
		for _, v := range adj[i] {
			adjSet[i][v] = true
		}
	}

	count := 0
	for u := 1; u <= n; u++ {
		for _, v := range adj[u] {
			if v > u { // Count each pair once
				for _, w := range adj[v] {
					if w > v && adjSet[u][w] {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfPaths([][]int{{1, 2}, {5, 1}, {1, 3}, {2, 4}, {4, 5}, {2, 3}}, 5))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfPaths([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}, 4))
	// Expected: 0 (no triangle)

	// Test case 3
	fmt.Println("Test 3:", numberOfPaths([][]int{{1, 2}, {2, 3}, {3, 1}, {1, 4}}, 4))
	// Expected: 1
}
```
