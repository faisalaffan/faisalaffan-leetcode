# 0867 — Transpose Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func transpose(matrix [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n). Space: O(m*n).  |  **Ruang:** O(m*n).


## 💻 Solusi Go

```go
package main

// LeetCode #867: Transpose Matrix
// https://leetcode.com/problems/transpose-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // [[1,4,7],[2,5,8],[3,6,9]]
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))             // [[1,4],[2,5],[3,6]]
}

// transpose returns the transpose of a matrix.
// Time: O(m*n). Space: O(m*n).
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
  // Matriks 2D
	result := make([][]int, n)
  // Range loop
	for i := range result {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
```
