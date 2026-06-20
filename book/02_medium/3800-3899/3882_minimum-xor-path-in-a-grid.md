# 3882 — Minimum Xor Path In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MinimumXorPathInAGrid(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(M * N * 2^B)  |  **Ruang:** O(N * 2^B) where B = 11 (grid values < 1024)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3882: Minimum XOR Path in a Grid
// https://leetcode.com/problems/minimum-xor-path-in-a-grid/
// Difficulty: Medium
// Time: O(M * N * 2^B) | Space: O(N * 2^B) where B = 11 (grid values < 1024)
// Approach: DP tracking reachable XOR values at each cell. Only right/down moves.

import (
	"fmt"
	"math"
)

func MinimumXorPathInAGrid(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return -1
	}
	n := len(grid[0])
  // Edge case: input kosong
	if n == 0 {
		return -1
	}

	maxXor := 2048 // 2^11 since grid[i][j] <= 1023

	// Use bitset (boolean array) for each cell
  // Matriks 2D
	dp := make([][]bool, n)
	for j := 0; j < n; j++ {
		dp[j] = make([]bool, maxXor)
	}

	dp[0][grid[0][0]] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cur := make([]bool, maxXor)
			if i > 0 {
				for x := 0; x < maxXor; x++ {
					if dp[j][x] {
						cur[x^grid[i][j]] = true
					}
				}
			}
			if j > 0 {
				for x := 0; x < maxXor; x++ {
					if dp[j-1][x] {
						cur[x^grid[i][j]] = true
					}
				}
			}
			dp[j] = cur
		}
	}

	ans := math.MaxInt32
	for x := 0; x < maxXor; x++ {
		if dp[n-1][x] && x < ans {
			ans = x
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumXorPathInAGrid([][]int{{1, 2}, {3, 4}})) // Expected: 6

	// Example 2
	fmt.Println(MinimumXorPathInAGrid([][]int{{6, 7}, {5, 8}})) // Expected: 9

	// Example 3
	fmt.Println(MinimumXorPathInAGrid([][]int{{2, 7, 5}})) // Expected: 0
}
```
