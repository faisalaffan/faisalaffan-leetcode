# 0909 — Snakes And Ladders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func SnakesAndLadders(board [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #909: Snakes and Ladders
// https://leetcode.com/problems/snakes-and-ladders/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SnakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}))
	fmt.Println(SnakesAndLadders([][]int{{-1, -1}, {-1, 3}}))
}

// Time: O(n^2) | Space: O(n^2)
func SnakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n

	// Convert board position to (row, col)
	posToCoord := func(pos int) (int, int) {
		row := (pos - 1) / n
		col := (pos - 1) % n
		if row%2 == 1 {
			col = n - 1 - col
		}
		return n - 1 - row, col
	}

  // Alokasi slice
	dist := make([]int, target+1)
  // Range loop
	for i := range dist {
		dist[i] = -1
	}
	dist[1] = 0

	queue := []int{1}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			return dist[curr]
		}

		for next := curr + 1; next <= curr+6 && next <= target; next++ {
			r, c := posToCoord(next)
			dest := next
			if board[r][c] != -1 {
				dest = board[r][c]
			}
			if dist[dest] == -1 {
				dist[dest] = dist[curr] + 1
				queue = append(queue, dest)
			}
		}
	}

	return -1
}
```
