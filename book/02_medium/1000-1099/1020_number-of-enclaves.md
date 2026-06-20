# 1020 — Number Of Enclaves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numEnclaves(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(m * n)  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1020: Number of Enclaves
// https://leetcode.com/problems/number-of-enclaves/
// Difficulty: Medium
//
// Approach: DFS from border cells, mark reachable land, then count remaining
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}})) // 3
	fmt.Println(numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}})) // 0
}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return
		}
		grid[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}
```
