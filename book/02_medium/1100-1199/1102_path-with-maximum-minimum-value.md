# 1102 — Path With Maximum Minimum Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumMinimumPath(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(m * n * log(m * n))  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

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
  // Matriks 2D
	visited := make([][]bool, m)
  // Range loop
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
