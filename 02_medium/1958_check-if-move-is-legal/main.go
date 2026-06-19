package main

// LeetCode #1958: Check if Move is Legal
// https://leetcode.com/problems/check-if-move-is-legal/
// Difficulty: Medium

import "fmt"

var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func main() {
	board := [][]byte{
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'}}
	fmt.Println(CheckMove(board, 4, 3, 'B'))

	board2 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'B', '.', '.', 'W', '.', '.', '.'},
		{'.', '.', 'W', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', 'W', '.', '.', 'W', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(CheckMove(board2, 4, 1, 'B'))
}

// Time: O(1) since board is always 8x8, Space: O(1)
func CheckMove(board [][]byte, rMove int, cMove int, color byte) bool {
	opponent := byte('B')
	if color == 'B' {
		opponent = 'W'
	}

	for _, d := range dirs {
		dr, dc := d[0], d[1]
		r, c := rMove+dr, cMove+dc
		steps := 0

		for r >= 0 && r < 8 && c >= 0 && c < 8 && board[r][c] == opponent {
			r += dr
			c += dc
			steps++
		}

		if steps > 0 && r >= 0 && r < 8 && c >= 0 && c < 8 && board[r][c] == color {
			return true
		}
	}
	return false
}
