# 2257 — Count Unguarded Cells In The Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countUnguarded(m int, n int, guards [][]int, walls [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n)  |  **Ruang:** O(m * n)


## 💻 Solusi Go

```go
package main

// LeetCode #2257: Count Unguarded Cells in the Grid
// https://leetcode.com/problems/count-unguarded-cells-in-the-grid/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
  // Matriks 2D
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	// 1 = wall, 2 = guard

	for _, w := range walls {
		grid[w[0]][w[1]] = 1
	}
	for _, g := range guards {
		grid[g[0]][g[1]] = 2
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, g := range guards {
		gr, gc := g[0], g[1]
		for _, d := range dirs {
			r, c := gr+d[0], gc+d[1]
			for r >= 0 && r < m && c >= 0 && c < n && grid[r][c] != 1 && grid[r][c] != 2 {
				if grid[r][c] == 0 {
					grid[r][c] = 3 // guarded
				}
				r += d[0]
				c += d[1]
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
	// Expected: 7

	// Test case 2
	fmt.Println(countUnguarded(3, 3, [][]int{{1, 1}}, [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}))
	// Expected: 4
}
```
