package main

// LeetCode #1252: Cells with Odd Values in a Matrix
// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
// Difficulty: Easy
// Time: O(n + m + k) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}})) // 6
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}})) // 0
}

// LeetCode submission: oddCells
func oddCells(m, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, idx := range indices {
		rows[idx[0]]++
		cols[idx[1]]++
	}
	oddRows, oddCols := 0, 0
	for _, v := range rows {
		if v%2 == 1 {
			oddRows++
		}
	}
	for _, v := range cols {
		if v%2 == 1 {
			oddCols++
		}
	}
	return oddRows*(n-oddCols) + (m-oddRows)*oddCols
}
