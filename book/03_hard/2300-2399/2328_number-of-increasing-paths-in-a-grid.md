# 2328 — Number Of Increasing Paths In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countPaths(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2328: Number of Increasing Paths in a Grid
// https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/
// Difficulty: Hard
//
// Approach: DFS + memoization. For each cell, count number of strictly
// increasing paths starting from that cell (including length-1 path).
// Use memoization: memo[r][c] = count of increasing paths starting at (r,c).
// Order doesn't matter since paths must be strictly increasing (no cycles).
// Result is sum of all memo values mod 1e9+7.

import "fmt"

func main() {
	// Example 1: [[1,1],[3,4]] => 8
	fmt.Println(countPaths([][]int{{1, 1}, {3, 4}}))
	// Example 2: [[1],[2]] => 3
	fmt.Println(countPaths([][]int{{1}, {2}}))
	// Edge: single cell
	fmt.Println(countPaths([][]int{{5}}))
	// Edge: all decreasing
	fmt.Println(countPaths([][]int{{5, 4}, {3, 2}}))
	// Edge: 3x3 all equal
	fmt.Println(countPaths([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

const mod = 1_000_000_007

func countPaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
  // Matriks 2D
	memo := make([][]int, m)
  // Range loop
	for i := range memo {
		memo[i] = make([]int, n)
		// -1 means uncomputed
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memo[r][c] != -1 {
			return memo[r][c]
		}
		// Each path starts at this cell (count = 1 for the cell itself)
		count := 1
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			if grid[nr][nc] <= grid[r][c] {
				continue
			}
			count = (count + dfs(nr, nc)) % mod
		}
		memo[r][c] = count
		return count
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = (ans + dfs(i, j)) % mod
		}
	}
	return ans
}
```
