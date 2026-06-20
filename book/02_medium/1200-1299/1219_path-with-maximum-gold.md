# 1219 — Path With Maximum Gold

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func getMaximumGold(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Backtracking

**Waktu:** O(m*n * 4^(k)) where k = max cells with gold  |  **Ruang:** O(k) for recursion

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1219: Path with Maximum Gold
// https://leetcode.com/problems/path-with-maximum-gold/
// Difficulty: Medium

// DFS from each cell with gold, backtracking. Max gold collected.

// Time: O(m*n * 4^(k)) where k = max cells with gold
// Space: O(k) for recursion

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	maxGold := 0

	var dfs func(r, c, gold int)
	dfs = func(r, c, gold int) {
		val := grid[r][c]
		gold += val
		if gold > maxGold {
			maxGold = gold
		}

		grid[r][c] = 0 // mark visited
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] > 0 {
				dfs(nr, nc, gold)
			}
		}
		grid[r][c] = val // restore
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				dfs(i, j, 0)
			}
		}
	}

	return maxGold
}

func main() {
	fmt.Printf("%d (expected: 24)\n",
		getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))

	fmt.Printf("%d (expected: 28)\n",
		getMaximumGold([][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}))
}
```
