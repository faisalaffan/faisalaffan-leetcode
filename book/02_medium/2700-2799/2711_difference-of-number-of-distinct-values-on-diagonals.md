# 2711 — Difference Of Number Of Distinct Values On Diagonals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func differenceOfDistinctValues(grid [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(m*n*(m+n))  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2711: Difference of Number of Distinct Values on Diagonals
// https://leetcode.com/problems/difference-of-number-of-distinct-values-on-diagonals/
// Difficulty: Medium
// Time: O(m*n*(m+n)) | Space: O(1)

import "fmt"

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
  // Matriks 2D
	ans := make([][]int, m)
  // Range loop
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Count distinct values above-left diagonal
  // HashMap: O(1) lookup
			aboveLeft := make(map[int]bool)
			r, c := i-1, j-1
			for r >= 0 && c >= 0 {
				aboveLeft[grid[r][c]] = true
				r--
				c--
			}

			// Count distinct values below-right diagonal
  // HashMap: O(1) lookup
			belowRight := make(map[int]bool)
			r, c = i+1, j+1
			for r < m && c < n {
				belowRight[grid[r][c]] = true
				r++
				c++
			}

			diff := len(aboveLeft) - len(belowRight)
			if diff < 0 {
				diff = -diff
			}
			ans[i][j] = diff
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", differenceOfDistinctValues([][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}))
	// Expected: [[1,1,0],[1,0,1],[0,1,1]]

	// Test case 2
	fmt.Println("Test 2:", differenceOfDistinctValues([][]int{{1}}))
	// Expected: [[0]]

	// Test case 3
	fmt.Println("Test 3:", differenceOfDistinctValues([][]int{{1, 2}, {3, 4}}))
	// Expected: [[1,0],[0,1]]
}
```
