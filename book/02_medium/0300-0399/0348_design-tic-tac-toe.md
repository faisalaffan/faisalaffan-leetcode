# 0348 — Design Tic Tac Toe

## Deskripsi

**Soal:** [0348. Design Tic Tac Toe](https://leetcode.com/problems/design-tic-tac-toe/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per move  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(n int) TicTacToe`

## Solusi Go

```go
package main

// LeetCode #348: Design Tic-Tac-Toe
// https://leetcode.com/problems/design-tic-tac-toe/
// Difficulty: Medium [Paid]
// Time: O(1) per move | Space: O(n)

import "fmt"

type TicTacToe struct {
	rows     []int
	cols     []int
	diag     int
	antiDiag int
	n        int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows: make([]int, n),
		cols: make([]int, n),
		diag: 0, antiDiag: 0, n: n,
	}
}

func (t *TicTacToe) Move(row int, col int, player int) int {
	// Player 1: +1, Player 2: -1
	val := 1
	if player == 2 {
		val = -1
	}

	t.rows[row] += val
	t.cols[col] += val

	if row == col {
		t.diag += val
	}
	if row+col == t.n-1 {
		t.antiDiag += val
	}

	// Check win
	if abs(t.rows[row]) == t.n || abs(t.cols[col]) == t.n ||
		abs(t.diag) == t.n || abs(t.antiDiag) == t.n {
		return player
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Test case 1
	toe := Constructor(3)
	fmt.Println("Move 1:", toe.Move(0, 0, 1)) // Player 1
	fmt.Println("Move 2:", toe.Move(0, 2, 2)) // Player 2
	fmt.Println("Move 3:", toe.Move(2, 2, 1)) // Player 1
	fmt.Println("Move 4:", toe.Move(1, 1, 2)) // Player 2
	fmt.Println("Move 5:", toe.Move(2, 0, 1)) // Player 1
	fmt.Println("Move 6:", toe.Move(1, 0, 2)) // Player 2
	fmt.Println("Move 7:", toe.Move(2, 1, 1)) // Player 1 - wins
	// Expected: 0, 0, 0, 0, 0, 0, 1

	fmt.Println()

	// Test case 2: Player 2 wins
	toe2 := Constructor(3)
	fmt.Println("Move 1:", toe2.Move(0, 0, 1))
	fmt.Println("Move 2:", toe2.Move(1, 0, 2))
	fmt.Println("Move 3:", toe2.Move(0, 2, 1))
	fmt.Println("Move 4:", toe2.Move(1, 1, 2))
	fmt.Println("Move 5:", toe2.Move(0, 1, 1))
	fmt.Println("Move 6:", toe2.Move(1, 2, 2)) // Player 2 wins
	// Expected: 0, 0, 0, 0, 0, 2
}
```
