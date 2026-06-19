package main

// LeetCode #51: N-Queens
// https://leetcode.com/problems/n-queens/
// Difficulty: Hard

import "fmt"

// solveNQueens returns all distinct solutions to the N-Queens puzzle.
// Each solution is represented as a board where 'Q' marks a queen and '.' marks empty.
//
// Complexity: O(n!) time, O(n) space for recursion/board (excluding output)
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = string(row)
	}

	cols := make([]bool, n)
	d1 := make([]bool, 2*n-1) // diagonal: row - col + n - 1
	d2 := make([]bool, 2*n-1) // anti-diagonal: row + col

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			solution := make([]string, n)
			copy(solution, board)
			result = append(result, solution)
			return
		}

		for col := 0; col < n; col++ {
			idx1 := row - col + n - 1
			idx2 := row + col
			if cols[col] || d1[idx1] || d2[idx2] {
				continue
			}

			// Place queen
			r := []byte(board[row])
			r[col] = 'Q'
			board[row] = string(r)
			cols[col] = true
			d1[idx1] = true
			d2[idx2] = true

			backtrack(row + 1)

			// Remove queen
			r[col] = '.'
			board[row] = string(r)
			cols[col] = false
			d1[idx1] = false
			d2[idx2] = false
		}
	}

	backtrack(0)
	return result
}

func main() {
	// Test case from LeetCode: n=4 -> 2 solutions
	solutions := solveNQueens(4)
	fmt.Printf("n=4 has %d solutions:\n", len(solutions))
	for i, sol := range solutions {
		fmt.Printf("Solution %d:\n", i+1)
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println()
	}

	// Test n=1
	solutions1 := solveNQueens(1)
	fmt.Printf("n=1 has %d solutions\n", len(solutions1))
}
