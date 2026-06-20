# 1139 — Largest 1 Bordered Square

## Deskripsi

**Soal:** [1139. Largest 1 Bordered Square](https://leetcode.com/problems/largest-1-bordered-square/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n * min(m, n))  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** —

**Fungsi Solusi:** `func largest1BorderedSquare(grid [][]int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1139: Largest 1-Bordered Square
// https://leetcode.com/problems/largest-1-bordered-square/
// Difficulty: Medium

// Precompute horizontal/vertical consecutive 1s, then check each cell as
// bottom-right corner.

// Time: O(m * n * min(m, n))
// Space: O(m * n)

func largest1BorderedSquare(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
  // Membuat slice 2D untuk DP/tabel
	hor := make([][]int, m)
  // Membuat slice 2D untuk DP/tabel
	ver := make([][]int, m)
	for i := 0; i < m; i++ {
		hor[i] = make([]int, n)
		ver[i] = make([]int, n)
	}
	maxSide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if j == 0 {
					hor[i][j] = 1
				} else {
					hor[i][j] = hor[i][j-1] + 1
				}
				if i == 0 {
					ver[i][j] = 1
				} else {
					ver[i][j] = ver[i-1][j] + 1
				}
				minSide := hor[i][j]
				if ver[i][j] < minSide {
					minSide = ver[i][j]
				}
				for s := minSide; s > maxSide; s-- {
					if hor[i-s+1][j] >= s && ver[i][j-s+1] >= s {
						maxSide = s
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func main() {
	fmt.Printf("%d (expected: 9)\n", largest1BorderedSquare([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Printf("%d (expected: 1)\n", largest1BorderedSquare([][]int{{1, 1, 0, 0}}))
	fmt.Printf("%d (expected: 0)\n", largest1BorderedSquare([][]int{{0}}))
}
```
