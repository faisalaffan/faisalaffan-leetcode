# 0037 — Sudoku Solver

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan papan Sudoku 9x9. Isi sel kosong dengan angka 1-9 sehingga setiap baris, kolom, dan sub-box 3x3 berisi 1-9 tanpa pengulangan.

**Cara berpikir:** Backtracking — coba angka 1-9, jika buntu mundur dan coba yang lain.

**Fungsi Solusi:** `func solveSudoku(board [][]byte) `

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #37: Sudoku Solver
// https://leetcode.com/problems/sudoku-solver/
// Difficulty: Hard

import "fmt"

// solveSudoku solves a 9x9 Sudoku board in-place using backtracking.
//
// Complexity: O(9^(n)) time where n is number of empty cells, O(n) space for recursion
func solveSudoku(board [][]byte) {
	// Pre-compute rows, cols, boxes for O(1) constraint checking
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				boxIdx := (i/3)*3 + j/3
				rows[i][num] = true
				cols[j][num] = true
				boxes[boxIdx][num] = true
			}
		}
	}

	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					boxIdx := (i/3)*3 + j/3
					for num := byte(0); num < 9; num++ {
						if !rows[i][num] && !cols[j][num] && !boxes[boxIdx][num] {
							board[i][j] = num + '1'
							rows[i][num] = true
							cols[j][num] = true
							boxes[boxIdx][num] = true

							if backtrack() {
								return true
							}

							board[i][j] = '.'
							rows[i][num] = false
							cols[j][num] = false
							boxes[boxIdx][num] = false
						}
					}
					return false
				}
			}
		}
		return true
	}

	backtrack()
}

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i > 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j > 0 {
				fmt.Print("| ")
			}
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}

func main() {
	// Test case from LeetCode
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Input:")
	printBoard(board)

	solveSudoku(board)

	fmt.Println("\nSolution:")
	printBoard(board)
}
```
