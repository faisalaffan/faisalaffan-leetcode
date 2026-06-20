# 0797 — All Paths From Source To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func allPathsSourceTarget(graph [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(2^n * n)  
**Kompleksitas Ruang:** O(2^n * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #797: All Paths From Source to Target
// https://leetcode.com/problems/all-paths-from-source-to-target/
// Difficulty: Medium
// Time: O(2^n * n)
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
  // Alokasi slice integer
	path := make([]int, 0)
	path = append(path, 0)

	var dfs func(node int)
	dfs = func(node int) {
		if node == len(graph)-1 {
  // Alokasi slice integer
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for _, neighbor := range graph[node] {
			path = append(path, neighbor)
			dfs(neighbor)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return result
}
```
