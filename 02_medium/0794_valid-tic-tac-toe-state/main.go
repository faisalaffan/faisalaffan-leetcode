package main

// LeetCode #794: Valid Tic-Tac-Toe State
// https://leetcode.com/problems/valid-tic-tac-toe-state/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}))
	fmt.Println(validTicTacToe([]string{"XOX", " X ", "   "}))
	fmt.Println(validTicTacToe([]string{"XOX", "O O", "XOX"}))
}

func validTicTacToe(board []string) bool {
	xCount, oCount := 0, 0
	for _, row := range board {
		for _, c := range row {
			if c == 'X' {
				xCount++
			} else if c == 'O' {
				oCount++
			}
		}
	}

	if xCount != oCount && xCount != oCount+1 {
		return false
	}

	xWin := isWinner(board, 'X')
	oWin := isWinner(board, 'O')

	if xWin && oWin {
		return false
	}
	if xWin && xCount != oCount+1 {
		return false
	}
	if oWin && xCount != oCount {
		return false
	}

	return true
}

func isWinner(board []string, player byte) bool {
	// Rows and columns
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	// Diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}
