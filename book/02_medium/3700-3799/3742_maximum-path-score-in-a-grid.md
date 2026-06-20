# 3742 — Maximum Path Score In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumPathScoreInAGrid(grid [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(m*n*k)  |  **Ruang:** O(m*n*k)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3742: Maximum Path Score in a Grid
// https://leetcode.com/problems/maximum-path-score-in-a-grid/
// Difficulty: Medium
// Time: O(m*n*k) | Space: O(m*n*k)

import "fmt"

func maximumPathScoreInAGrid(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	maxCost := k + 1
	if m+n+1 < maxCost {
		maxCost = m + n + 1
	}

	// dp[i][j][c] = max score at (i,j) with cost c, -1 = unreachable
  // Matriks 2D
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, maxCost)
			for c := 0; c < maxCost; c++ {
				dp[i][j][c] = -1
			}
		}
	}

	dp[0][0][0] = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for c := 0; c < maxCost; c++ {
				cur := dp[i][j][c]
				if cur == -1 {
					continue
				}
				// Move right
				if j+1 < n {
					add := 0
					if grid[i][j+1] > 0 {
						add = 1
					}
					if c+add < maxCost {
						val := cur + grid[i][j+1]
						if val > dp[i][j+1][c+add] {
							dp[i][j+1][c+add] = val
						}
					}
				}
				// Move down
				if i+1 < m {
					add := 0
					if grid[i+1][j] > 0 {
						add = 1
					}
					if c+add < maxCost {
						val := cur + grid[i+1][j]
						if val > dp[i+1][j][c+add] {
							dp[i+1][j][c+add] = val
						}
					}
				}
			}
		}
	}

	ans := -1
	for c := 0; c < maxCost; c++ {
		if dp[m-1][n-1][c] > ans {
			ans = dp[m-1][n-1][c]
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumPathScoreInAGrid([][]int{{0, 1, 2}, {1, 0, 1}, {2, 1, 0}}, 3))
	fmt.Println(maximumPathScoreInAGrid([][]int{{0, 0}, {0, 0}}, 1))
	fmt.Println(maximumPathScoreInAGrid([][]int{{0, 2}, {2, 0}}, 1))
}
```
