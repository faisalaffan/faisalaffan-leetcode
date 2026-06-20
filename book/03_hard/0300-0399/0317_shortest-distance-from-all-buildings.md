# 0317 — Shortest Distance From All Buildings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func shortestDistance(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #317: Shortest Distance from All Buildings
// https://leetcode.com/problems/shortest-distance-from-all-buildings/
// Difficulty: Hard [Paid]

import "fmt"

func shortestDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	rows, cols := len(grid), len(grid[0])

	// totalDist[r][c] = sum of distances from all reachable buildings
  // Matriks 2D
	totalDist := make([][]int, rows)
	// reachable[r][c] = count of buildings that can reach this cell
  // Matriks 2D
	reachable := make([][]int, rows)
	for r := 0; r < rows; r++ {
		totalDist[r] = make([]int, cols)
		reachable[r] = make([]int, cols)
	}

	totalBuildings := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				totalBuildings++
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// BFS from each building
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 1 {
				continue
			}

  // Matriks 2D
			visited := make([][]bool, rows)
			for i := 0; i < rows; i++ {
				visited[i] = make([]bool, cols)
			}

			type cell struct{ r, c int }
			queue := []cell{{r, c}}
			visited[r][c] = true
			dist := 0

			for len(queue) > 0 {
				dist++
				nextQ := []cell{}
				for _, cur := range queue {
					for _, d := range dirs {
						nr, nc := cur.r+d[0], cur.c+d[1]
						if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
							continue
						}
						if visited[nr][nc] || grid[nr][nc] != 0 {
							continue
						}
						visited[nr][nc] = true
						totalDist[nr][nc] += dist
						reachable[nr][nc]++
						nextQ = append(nextQ, cell{nr, nc})
					}
				}
				queue = nextQ
			}
		}
	}

	ans := -1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 && reachable[r][c] == totalBuildings {
				if ans == -1 || totalDist[r][c] < ans {
					ans = totalDist[r][c]
				}
			}
		}
	}
	return ans
}

func main() {
	// Example 1
	grid1 := [][]int{
		{1, 0, 2, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println(shortestDistance(grid1))
	// 7

	// Example 2
	grid2 := [][]int{{1}}
	fmt.Println(shortestDistance(grid2))
	// -1

	// Example 3
	grid3 := [][]int{{1, 0}}
	fmt.Println(shortestDistance(grid3))
	// 1
}
```
