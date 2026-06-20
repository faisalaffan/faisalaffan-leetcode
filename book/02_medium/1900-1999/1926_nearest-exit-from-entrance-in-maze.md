# 1926 — Nearest Exit From Entrance In Maze

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func NearestExit(maze [][]byte, entrance []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(m*n), Space: O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1926: Nearest Exit from Entrance in Maze
// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
// Difficulty: Medium

import "fmt"

func main() {
	maze := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'}}
	fmt.Println(NearestExit(maze, []int{1, 2}))

	maze2 := [][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'}}
	fmt.Println(NearestExit(maze2, []int{1, 0}))
}

// Time: O(m*n), Space: O(m*n)
func NearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][2]int{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = '+' // mark as visited
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		steps++
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && maze[nr][nc] == '.' {
					if nr == 0 || nr == m-1 || nc == 0 || nc == n-1 {
						return steps
					}
					maze[nr][nc] = '+'
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
	}
	return -1
}
```
