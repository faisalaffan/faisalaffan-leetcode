# 2290 — Minimum Obstacle Removal To Reach Corner

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minimumObstacles(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Dijkstra

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Dijkstra** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2290: Minimum Obstacle Removal to Reach Corner
// https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/
// Difficulty: Hard
//
// Approach: 0-1 BFS (Dijkstra with deque since edge weights are 0 or 1).
// Treat grid cells as nodes. Moving from (r,c) to (nr,nc) costs grid[nr][nc]
// (0 if empty, 1 if obstacle). Use deque: push front for cost 0, push back for cost 1.

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Example 1: [[0,1,1],[1,1,0],[1,1,0]] => 2
	fmt.Println(minimumObstacles([][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}))
	// Example 2: [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]] => 0
	fmt.Println(minimumObstacles([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}))
	// Edge: single cell
	fmt.Println(minimumObstacles([][]int{{0}}))
	fmt.Println(minimumObstacles([][]int{{1}}))
	// Edge: 1xN row, no obstacles
	fmt.Println(minimumObstacles([][]int{{0, 0, 0, 0}}))
}

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
  // Matriks 2D
	dist := make([][]int, m)
  // Range loop
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	deque := list.New()
	dist[0][0] = grid[0][0]
	deque.PushFront([2]int{0, 0})

	for deque.Len() > 0 {
		front := deque.Remove(deque.Front()).([2]int)
		r, c := front[0], front[1]

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			cost := grid[nr][nc]
			nd := dist[r][c] + cost
			if nd < dist[nr][nc] {
				dist[nr][nc] = nd
				if cost == 0 {
					deque.PushFront([2]int{nr, nc})
				} else {
					deque.PushBack([2]int{nr, nc})
				}
			}
		}
	}
	return dist[m-1][n-1]
}
```
