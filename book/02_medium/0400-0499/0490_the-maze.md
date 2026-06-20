# 0490 — The Maze

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func TheMaze(maze [][]int, start []int, destination []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(m * n)  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #490: The Maze
// https://leetcode.com/problems/the-maze/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{4, 4}))
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{3, 2}))
}

func TheMaze(maze [][]int, start []int, destination []int) bool {
	m, n := len(maze), len(maze[0])
  // Matriks 2D
	visited := make([][]bool, m)
  // Range loop
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{start}
	visited[start[0]][start[1]] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == destination[0] && cur[1] == destination[1] {
			return true
		}

		for _, d := range dirs {
			r, c := cur[0], cur[1]
			// Roll until hitting a wall
			for r+d[0] >= 0 && r+d[0] < m && c+d[1] >= 0 && c+d[1] < n && maze[r+d[0]][c+d[1]] == 0 {
				r += d[0]
				c += d[1]
			}
			if !visited[r][c] {
				visited[r][c] = true
				queue = append(queue, []int{r, c})
			}
		}
	}

	return false
}
```
