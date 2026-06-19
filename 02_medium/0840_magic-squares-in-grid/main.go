package main

// LeetCode #840: Magic Squares In Grid
// https://leetcode.com/problems/magic-squares-in-grid/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MagicSquaresInGrid([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
	fmt.Println(MagicSquaresInGrid([][]int{{8}}))
	fmt.Println(MagicSquaresInGrid([][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}))
}

// Time: O(m * n) | Space: O(1)
func MagicSquaresInGrid(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m < 3 || n < 3 {
		return 0
	}

	ans := 0
	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if isMagic(grid, i, j) {
				ans++
			}
		}
	}
	return ans
}

func isMagic(grid [][]int, r, c int) bool {
	seen := [16]bool{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := grid[r+i][c+j]
			if v < 1 || v > 9 || seen[v] {
				return false
			}
			seen[v] = true
		}
	}

	sum := grid[r][c] + grid[r][c+1] + grid[r][c+2]
	for i := 0; i < 3; i++ {
		if grid[r+i][c]+grid[r+i][c+1]+grid[r+i][c+2] != sum {
			return false
		}
		if grid[r][c+i]+grid[r+1][c+i]+grid[r+2][c+i] != sum {
			return false
		}
	}
	if grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2] != sum {
		return false
	}
	if grid[r][c+2]+grid[r+1][c+1]+grid[r+2][c] != sum {
		return false
	}

	return true
}
