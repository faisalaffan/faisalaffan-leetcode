# 0566 — Reshape The Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ReshapeTheMatrix(mat [][]int, r, c int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n), Space: O(m*n)  |  **Ruang:** O(m*n)


## 💻 Solusi Go

```go
package main

// LeetCode #566: Reshape the Matrix
// https://leetcode.com/problems/reshape-the-matrix/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(m*n)
func ReshapeTheMatrix(mat [][]int, r, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
  // Matriks 2D
	result := make([][]int, r)
  // Range loop
	for i := range result {
		result[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		result[i/c][i%c] = mat[i/n][i%n]
	}
	return result
}

func main() {
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 2, 4))
}
```
