# 0289 — Game Of Life

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func gameOfLife(board [][]int) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #289: Game of Life
// https://leetcode.com/problems/game-of-life/
// Difficulty: Medium
// Time: O(m*n), Space: O(1)

import "fmt"

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	countLive := func(r, c int) int {
		count := 0
		dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && (board[nr][nc] == 1 || board[nr][nc] == 2) {
				count++
			}
		}
		return count
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			live := countLive(r, c)
			if board[r][c] == 1 && (live < 2 || live > 3) {
				board[r][c] = 2
			}
			if board[r][c] == 0 && live == 3 {
				board[r][c] = -1
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 2 {
				board[r][c] = 0
			} else if board[r][c] == -1 {
				board[r][c] = 1
			}
		}
	}
}

func main() {
	board1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board1)
	fmt.Println(board1)

	board2 := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board2)
	fmt.Println(board2)
}
```
