# 3249 — Count The Number Of Good Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func countGoodNodes(edges [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3249: Count the Number of Good Nodes
// https://leetcode.com/problems/count-the-number-of-good-nodes/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	ans := 0

	var dfs func(u, parent int) int
	dfs = func(u, parent int) int {
		size := 1
		childSize := -1
		good := true

		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			sz := dfs(v, u)
			if childSize == -1 {
				childSize = sz
			} else if sz != childSize {
				good = false
			}
			size += sz
		}

		if good {
			ans++
		}
		return size
	}

	dfs(0, -1)
	return ans
}

func main() {
	fmt.Println(countGoodNodes([][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}})) // Expected: 7
	fmt.Println(countGoodNodes([][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}))        // Expected: 6
}
```
