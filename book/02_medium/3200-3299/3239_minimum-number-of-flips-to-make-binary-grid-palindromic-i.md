# 3239 — Minimum Number Of Flips To Make Binary Grid Palindromic I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minFlips(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3239: Minimum Number of Flips to Make Binary Grid Palindromic I
// https://leetcode.com/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rowFlips := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			if grid[i][j] != grid[i][n-1-j] {
				rowFlips++
			}
		}
	}

	colFlips := 0
	for j := 0; j < n; j++ {
		for i := 0; i < m/2; i++ {
			if grid[i][j] != grid[m-1-i][j] {
				colFlips++
			}
		}
	}

	return min(rowFlips, colFlips)
}

func main() {
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}})) // Expected: 2
	fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))          // Expected: 1
}
```
