package main

// LeetCode #419: Battleships in a Board
// https://leetcode.com/problems/battleships-in-a-board/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func countBattleships(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				// Count only if it's the start of a ship (no X above or to the left)
				if (i == 0 || board[i-1][j] != 'X') && (j == 0 || board[i][j-1] != 'X') {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	b1 := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	fmt.Println("Test 1:", countBattleships(b1))
	// Expected: 2

	// Test case 2
	b2 := [][]byte{{'X'}}
	fmt.Println("Test 2:", countBattleships(b2))
	// Expected: 1

	// Test case 3: Empty
	b3 := [][]byte{{'.', '.'}, {'.', '.'}}
	fmt.Println("Test 3:", countBattleships(b3))
	// Expected: 0
}
