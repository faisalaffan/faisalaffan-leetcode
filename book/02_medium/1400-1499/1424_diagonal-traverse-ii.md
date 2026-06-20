# 1424 — Diagonal Traverse Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findDiagonalOrder(nums [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n) where m = number of rows, n = avg columns  |  **Ruang:** O(m*n) for result


## 💻 Solusi Go

```go
package main

// LeetCode #1424: Diagonal Traverse II
// https://leetcode.com/problems/diagonal-traverse-ii/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// [1,4,2,7,5,3,8,6,9]

	// Test case 2
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3, 4, 5},
		{6, 7},
		{8},
		{9, 10, 11},
		{12, 13, 14, 15, 16},
	}))
	// [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]

	// Test case 3
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}}))
	// [1,2,3]
}

// Time: O(m*n) where m = number of rows, n = avg columns
// Space: O(m*n) for result
func findDiagonalOrder(nums [][]int) []int {
	// Group elements by (i+j) diagonal index
	// For each diagonal, elements appear in reverse row order
  // Matriks 2D
	diagonals := make([][]int, 0)

	for i, row := range nums {
		for j := range row {
			idx := i + j
			if idx >= len(diagonals) {
				diagonals = append(diagonals, []int{})
			}
			// Prepend to maintain reverse row order within diagonal
			diagonals[idx] = append(diagonals[idx], nums[i][j])
		}
	}

  // Alokasi slice
	result := make([]int, 0)
	for _, d := range diagonals {
		// Reverse the diagonal (since we prepended, it's in reverse)
		for i := len(d) - 1; i >= 0; i-- {
			result = append(result, d[i])
		}
	}

	return result
}
```
