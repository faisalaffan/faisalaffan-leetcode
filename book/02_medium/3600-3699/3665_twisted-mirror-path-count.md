# 3665 — Twisted Mirror Path Count

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func twistedMirrorPathCount(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3665: Twisted Mirror Path Count
// https://leetcode.com/problems/twisted-mirror-path-count/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func twistedMirrorPathCount(grid [][]int) int {
	mod := int(1e9 + 7)
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Matriks 2D
	dp := make([][]int, m)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				continue
			}
			cur := dp[i][j]

			// Try moving right to (i, j+1)
			if j+1 < n {
				if grid[i][j+1] == 1 {
					// Mirror at (i, j+1): reflect down to (i+1, j+1)
					if i+1 < m {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i][j+1] = (dp[i][j+1] + cur) % mod
				}
			}

			// Try moving down to (i+1, j)
			if i+1 < m {
				if grid[i+1][j] == 1 {
					// Mirror at (i+1, j): reflect right to (i+1, j+1)
					if j+1 < n {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i+1][j] = (dp[i+1][j] + cur) % mod
				}
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 0}, {0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1}, {1, 0}}))
}
```
