package main

// LeetCode #2018: Check if Word Can Be Placed In Crossword
// https://leetcode.com/problems/check-if-word-can-be-placed-in-crossword/
// Difficulty: Medium
// Time: O(m*n*len(word)) | Space: O(1)

import "fmt"

func placeWordInCrossword(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	w := len(word)
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				continue
			}
			for _, d := range dirs {
				// Check if this is a valid starting position:
				// either at edge or previous cell is blocked
				pi, pj := i-d[0], j-d[1]
				if pi >= 0 && pi < m && pj >= 0 && pj < n && board[pi][pj] != '#' {
					continue
				}

				// Try to place the word
				ok := true
				for k := 0; k < w; k++ {
					ci, cj := i+k*d[0], j+k*d[1]
					if ci < 0 || ci >= m || cj < 0 || cj >= n {
						ok = false
						break
					}
					if board[ci][cj] == '#' || (board[ci][cj] != ' ' && board[ci][cj] != word[k]) {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}

				// Check that cell after word is blocked or out of bounds
				ni, nj := i+w*d[0], j+w*d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && board[ni][nj] != '#' {
					continue
				}

				return true
			}
		}
	}
	return false
}

func main() {
	// Test case 1
	board1 := [][]byte{
		{'#', ' ', '#'},
		{' ', ' ', '#'},
		{'#', 'c', ' '},
	}
	fmt.Println("Test 1:", placeWordInCrossword(board1, "abc"))
	// Expected: true

	// Test case 2
	board2 := [][]byte{
		{' ', '#', 'a'},
		{' ', '#', 'c'},
		{' ', '#', 'a'},
	}
	fmt.Println("Test 2:", placeWordInCrossword(board2, "ac"))
	// Expected: false

	// Test case 3
	board3 := [][]byte{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}
	fmt.Println("Test 3:", placeWordInCrossword(board3, "hello"))
	// Expected: false (board too small)
}
