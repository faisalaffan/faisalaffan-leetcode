# 1582 — Special Positions In A Binary Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numSpecial(mat [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n), Space: O(m + n)  |  **Ruang:** O(m + n)


## 💻 Solusi Go

```go
package main

// LeetCode #1582: Special Positions in a Binary Matrix
// https://leetcode.com/problems/special-positions-in-a-binary-matrix/
// Difficulty: Easy
//
// LeetCode submission: func numSpecial(mat [][]int) int

import "fmt"

func main() {
	mat1 := [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	fmt.Println(SpecialPositionsInABinaryMatrix(mat1)) // 1

	mat2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(SpecialPositionsInABinaryMatrix(mat2)) // 3
}

// Time: O(m * n), Space: O(m + n)
func SpecialPositionsInABinaryMatrix(mat [][]int) int {
	m, n := len(mat), len(mat[0])
  // Alokasi slice
	rows := make([]int, m)
  // Alokasi slice
	cols := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}
	return count
}
```
