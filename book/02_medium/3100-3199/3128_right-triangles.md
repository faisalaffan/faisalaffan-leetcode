# 3128 — Right Triangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numberOfRightTriangles(grid [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n)  |  **Ruang:** O(m + n)


## 💻 Solusi Go

```go
package main

// LeetCode #3128: Right Triangles
// https://leetcode.com/problems/right-triangles/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m + n)

import "fmt"

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Alokasi slice
	rowSum := make([]int, m)
  // Alokasi slice
	colSum := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowSum[i]++
				colSum[j]++
			}
		}
	}

	var ans int64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans += int64(rowSum[i]-1) * int64(colSum[j]-1)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfRightTriangles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}})) // Expected: 2
	fmt.Println(numberOfRightTriangles([][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}})) // Expected: 2
}
```
