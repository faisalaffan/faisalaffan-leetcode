# 1572 — Matrix Diagonal Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func diagonalSum(mat [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1572: Matrix Diagonal Sum
// https://leetcode.com/problems/matrix-diagonal-sum/
// Difficulty: Easy
//
// LeetCode submission: func diagonalSum(mat [][]int) int

import "fmt"

func main() {
	mat1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(MatrixDiagonalSum(mat1)) // 25

	mat2 := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	fmt.Println(MatrixDiagonalSum(mat2)) // 8
}

// Time: O(n), Space: O(1)
func MatrixDiagonalSum(mat [][]int) int {
	n := len(mat)
	sum := 0
	for i := 0; i < n; i++ {
		sum += mat[i][i]
		sum += mat[i][n-1-i]
	}
	if n%2 == 1 {
		mid := n / 2
		sum -= mat[mid][mid]
	}
	return sum
}
```
