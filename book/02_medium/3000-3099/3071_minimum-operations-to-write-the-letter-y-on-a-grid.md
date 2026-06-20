# 3071 — Minimum Operations To Write The Letter Y On A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minimumOperationsToWriteY(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3071: Minimum Operations to Write the Letter Y on a Grid
// https://leetcode.com/problems/minimum-operations-to-write-the-letter-y-on-a-grid/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumOperationsToWriteY([][]int{{1, 2, 2}, {2, 2, 3}, {2, 3, 3}}))
	fmt.Println(minimumOperationsToWriteY([][]int{{0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}}))
}

func minimumOperationsToWriteY(grid [][]int) int {
	n := len(grid)
	// Find max value
	maxV := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > maxV {
				maxV = grid[i][j]
			}
		}
	}
	size := maxV + 1
  // Alokasi slice
	yCnt := make([]int, size)
  // Alokasi slice
	notYCnt := make([]int, size)
	center := n / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			isY := false
			if i == j && i <= center {
				isY = true
			} else if i+j == n-1 && i <= center {
				isY = true
			} else if j == center && i >= center {
				isY = true
			}
			if isY {
				yCnt[v]++
			} else {
				notYCnt[v]++
			}
		}
	}
	totalY := 0
	totalNotY := 0
	for _, c := range yCnt {
		totalY += c
	}
	for _, c := range notYCnt {
		totalNotY += c
	}
	ans := n * n
	for yv := 0; yv < size; yv++ {
		for nv := 0; nv < size; nv++ {
			if yv == nv {
				continue
			}
			ops := totalY - yCnt[yv] + totalNotY - notYCnt[nv]
			if ops < ans {
				ans = ops
			}
		}
	}
	return ans
}
```
