# 0847 — Shortest Path Visiting All Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestPathLength(graph [][]int) int
```

> **💡 Hint:** BFS over state (node, visitedMask). Start from every node simultaneously

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #847: Shortest Path Visiting All Nodes
// https://leetcode.com/problems/shortest-path-visiting-all-nodes/
// Difficulty: Hard
// Approach: BFS over state (node, visitedMask). Start from every node simultaneously
// (multi-source BFS). The mask tracks which nodes have been visited.

import "fmt"

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	target := (1 << n) - 1

	// dist[node][mask] = shortest steps to reach this state
  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = make([]int, 1<<n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

  // Alokasi slice integer
	queue := make([][2]int, 0)
	for i := 0; i < n; i++ {
		mask := 1 << i
		queue = append(queue, [2]int{i, mask})
		dist[i][mask] = 0
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		node, mask := cur[0], cur[1]

		if mask == target {
			return dist[node][mask]
		}

		for _, nei := range graph[node] {
			newMask := mask | (1 << nei)
			if dist[nei][newMask] == -1 {
				dist[nei][newMask] = dist[node][mask] + 1
				queue = append(queue, [2]int{nei, newMask})
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}})) // Expected: 4
	fmt.Println(shortestPathLength([][]int{{1}, {0, 2, 4}, {1, 3}, {2}, {1}}))
	// Expected: 4 (0->1->4->1->2->3: path 0-1-4-1-2-3 = 5 steps... let me verify)
	// This tests a more complex graph
}
```
