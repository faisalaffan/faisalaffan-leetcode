# 3459 — Length Of Longest V Shaped Diagonal Segment

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func lenOfVDiagonal(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3459: Length of Longest V-Shaped Diagonal Segment
// https://leetcode.com/problems/length-of-longest-v-shaped-diagonal-segment/
// Difficulty: Hard
//
// Given a grid, find the longest diagonal segment that forms a V shape
// (decreasing then increasing values). A diagonal segment moves in one
// of the four diagonal directions (down-right, down-left, up-right, up-left)
// and must form a V pattern: decreasing for some steps then increasing.
//
// Approach: DP from each cell in all diagonal directions. Track length of
// decreasing and increasing runs from each cell.

import "fmt"

func main() {
	// Example 1
	fmt.Println(lenOfVDiagonal([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// Example 2
	fmt.Println(lenOfVDiagonal([][]int{{1, 1}, {1, 1}}))
	// Example 3: single cell
	fmt.Println(lenOfVDiagonal([][]int{{5}}))
	// Edge: 2x3 grid
	fmt.Println(lenOfVDiagonal([][]int{{1, 2}, {3, 4}, {5, 6}}))
}

func lenOfVDiagonal(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// dpDec[i][j][dir] = longest decreasing diagonal starting at (i,j) in direction dir
	// dpInc[i][j][dir] = longest increasing diagonal starting at (i,j) in direction dir
	// dir: 0=down-right, 1=down-left, 2=up-right, 3=up-left
  // Matriks 2D
	dpDec := make([][][]int, m)
  // Matriks 2D
	dpInc := make([][][]int, m)
  // Range loop
	for i := range dpDec {
		dpDec[i] = make([][]int, n)
		dpInc[i] = make([][]int, n)
		for j := range dpDec[i] {
			dpDec[i][j] = []int{0, 0, 0, 0}
			dpInc[i][j] = []int{0, 0, 0, 0}
		}
	}

	dirs := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	ans := 1

	// Process cells in reverse diagonal order so dependencies are computed
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			for d, dir := range dirs {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					if grid[ni][nj] == grid[i][j]-1 {
						dpDec[i][j][d] = dpDec[ni][nj][d] + 1
					} else {
						dpDec[i][j][d] = 0
					}
					if grid[ni][nj] == grid[i][j]+1 {
						dpInc[i][j][d] = dpInc[ni][nj][d] + 1
					} else {
						dpInc[i][j][d] = 0
					}
				}
			}
		}
	}

	// For each cell, try all pairs of opposite directions
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// V-shape: decreasing in one direction, then increasing in opposite
			for d := 0; d < 4; d++ {
				opp := 3 - d // opposite direction
				dec := dpDec[i][j][d]
				inc := dpInc[i][j][opp]
				if dec > 0 && inc > 0 {
					length := dec + inc + 1
					if length > ans {
						ans = length
					}
				}
				// Just decreasing or just increasing also counts as a longer line
				if dec+1 > ans {
					ans = dec + 1
				}
				if inc+1 > ans {
					ans = inc + 1
				}
			}
		}
	}

	return ans
}
```
