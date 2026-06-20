# 2661 — First Completely Painted Row Or Column

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func firstCompleteIndex(arr []int, mat [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)


## 💻 Solusi Go

```go
package main

// LeetCode #2661: First Completely Painted Row or Column
// https://leetcode.com/problems/first-completely-painted-row-or-column/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
  // Alokasi slice
	pos := make([][2]int, m*n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pos[mat[i][j]] = [2]int{i, j}
		}
	}

  // Alokasi slice
	rowCount := make([]int, m)
  // Alokasi slice
	colCount := make([]int, n)

	for idx, v := range arr {
		r, c := pos[v][0], pos[v][1]
		rowCount[r]++
		colCount[c]++
		if rowCount[r] == n || colCount[c] == m {
			return idx
		}
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", firstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", firstCompleteIndex([]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", firstCompleteIndex([]int{1, 2}, [][]int{{1}, {2}}))
	// Expected: 0
}
```
