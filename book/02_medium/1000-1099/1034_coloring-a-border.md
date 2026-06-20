# 1034 — Coloring A Border

## Deskripsi

**Soal:** [1034. Coloring A Border](https://leetcode.com/problems/coloring-a-border/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** DFS to find connected component, then color border cells

## Solusi Go

```go
package main

// LeetCode #1034: Coloring A Border
// https://leetcode.com/problems/coloring-a-border/
// Difficulty: Medium
//
// Approach: DFS to find connected component, then color border cells
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{{1, 1}, {1, 2}}, 0, 0, 3)) // [[3,3],[3,2]]
	fmt.Println(colorBorder([][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3)) // [[1,3,3],[2,3,3]]
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	origColor := grid[row][col]
	if origColor == color {
		return grid
	}

  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		visited[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] != origColor {
				grid[r][c] = color
			} else if !visited[nr][nc] {
				dfs(nr, nc)
			}
		}
	}

	dfs(row, col)
	return grid
}
```
