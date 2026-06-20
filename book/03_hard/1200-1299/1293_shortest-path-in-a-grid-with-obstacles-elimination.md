# 1293 — Shortest Path In A Grid With Obstacles Elimination

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func shortestPath(grid [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1293: Shortest Path in a Grid with Obstacles Elimination
// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1293. Shortest Path in a Grid with Obstacles Elimination")
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}
	fmt.Println("k=1:", shortestPath(grid, 1), "(expected 6)")

	grid2 := [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}
	fmt.Println("k=1:", shortestPath(grid2, 1), "(expected -1)")
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	// visited[r][c][e] = true if visited (r,c) with e eliminations used.
  // Matriks 2D
	visited := make([][][]bool, m)
  // Range loop
	for i := range visited {
		visited[i] = make([][]bool, n)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}

	type state struct {
		r, c, elim, dist int
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	queue := []state{{0, 0, 0, 0}}
	visited[0][0][0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.r == m-1 && cur.c == n-1 {
			return cur.dist
		}

		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			newElim := cur.elim
			if grid[nr][nc] == 1 {
				newElim++
			}
			if newElim > k {
				continue
			}

			if !visited[nr][nc][newElim] {
				visited[nr][nc][newElim] = true
				queue = append(queue, state{nr, nc, newElim, cur.dist + 1})
			}
		}
	}

	return -1
}
```
