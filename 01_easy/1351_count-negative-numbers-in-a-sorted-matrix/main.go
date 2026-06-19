package main

// LeetCode #1351: Count Negative Numbers in a Sorted Matrix
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
// Difficulty: Easy
//
// LeetCode submission: func countNegatives(grid [][]int) int

import "fmt"

func main() {
	grid1 := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(CountNegativeNumbersInASortedMatrix(grid1)) // 8

	grid2 := [][]int{
		{3, 2},
		{1, 0},
	}
	fmt.Println(CountNegativeNumbersInASortedMatrix(grid2)) // 0
}

// Time: O(m + n), Space: O(1)
func CountNegativeNumbersInASortedMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0
	row, col := 0, n-1
	for row < m && col >= 0 {
		if grid[row][col] < 0 {
			count += m - row
			col--
		} else {
			row++
		}
	}
	return count
}
