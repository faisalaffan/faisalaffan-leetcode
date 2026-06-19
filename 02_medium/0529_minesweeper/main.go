package main

// LeetCode #529: Minesweeper
// https://leetcode.com/problems/minesweeper/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	result := UpdateBoard(board, []int{3, 0})
	for _, row := range result {
		fmt.Println(string(row))
	}
}

func UpdateBoard(board [][]byte, click []int) [][]byte {
	r, c := click[0], click[1]
	if board[r][c] == 'M' {
		board[r][c] = 'X'
		return board
	}
	reveal(board, r, c)
	return board
}

func reveal(board [][]byte, r, c int) {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != 'E' {
		return
	}
	m, n := len(board), len(board[0])
	mines := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			nr, nc := r+dr, c+dc
			if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == 'M' {
				mines++
			}
		}
	}
	if mines > 0 {
		board[r][c] = byte('0' + mines)
		return
	}
	board[r][c] = 'B'
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			reveal(board, r+dr, c+dc)
		}
	}
}
