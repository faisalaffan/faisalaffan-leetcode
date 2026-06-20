# 1275 — Find Winner On A Tic Tac Toe Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func tictactoe(moves [][]int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1275: Find Winner on a Tic Tac Toe Game
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}})) // "A"
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}})) // "B"
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}})) // "Draw"
}

// LeetCode submission: tictactoe
func tictactoe(moves [][]int) string {
  // Matriks 2D
	board := make([][]byte, 3)
  // Range loop
	for i := range board {
		board[i] = make([]byte, 3)
	}
	for i, m := range moves {
		player := byte('A')
		if i%2 == 1 {
			player = 'B'
		}
		board[m[0]][m[1]] = player
	}
	// Check rows and cols
	for i := 0; i < 3; i++ {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return string(board[i][0])
		}
		if board[0][i] != 0 && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return string(board[0][i])
		}
	}
	// Check diagonals
	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return string(board[0][0])
	}
	if board[0][2] != 0 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return string(board[0][2])
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
```
