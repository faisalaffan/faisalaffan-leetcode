# 1102 — Path With Maximum Minimum Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumMinimumPath(grid [][]int) int
```

> **💡 Hint:** BFS with max-heap (priority queue). Always visit cell with largest value.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(m * n * log(m * n))  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1102: Path With Maximum Minimum Value
// https://leetcode.com/problems/path-with-maximum-minimum-value/
// Difficulty: Medium
//
// Approach: BFS with max-heap (priority queue). Always visit cell with largest value.
// Time: O(m * n * log(m * n))
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maximumMinimumPath([][]int{{5, 4, 5}, {1, 2, 6}, {7, 4, 6}})) // 4
	fmt.Println(maximumMinimumPath([][]int{{2, 0, 1}, {4, 3, 2}, {5, 4, 5}})) // 3
}

type cell struct {
	i, j, val int
}

func maximumMinimumPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Max-heap using slice
	heap := []cell{{0, 0, grid[0][0]}}
	visited[0][0] = true
	result := grid[0][0]

	for len(heap) > 0 {
		// Find max value in heap (simple linear scan for small heap)
		maxIdx := 0
		for k := 1; k < len(heap); k++ {
			if heap[k].val > heap[maxIdx].val {
				maxIdx = k
			}
		}
		cur := heap[maxIdx]
		heap[maxIdx] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]

		if cur.val < result {
			result = cur.val
		}

		if cur.i == m-1 && cur.j == n-1 {
			return result
		}

		for _, d := range dirs {
			ni, nj := cur.i+d[0], cur.j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] {
				visited[ni][nj] = true
				heap = append(heap, cell{ni, nj, grid[ni][nj]})
			}
		}
	}

	return result
}
```
