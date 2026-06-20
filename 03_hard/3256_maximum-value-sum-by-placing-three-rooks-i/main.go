package main

// LeetCode #3256: Maximum Value Sum by Placing Three Rooks I
// https://leetcode.com/problems/maximum-value-sum-by-placing-three-rooks-i/
// Difficulty: Hard
//
// Place three rooks on an m x n board such that no two rooks share the same
// row or column. Maximize the sum of their cell values.
//
// Approach: For each row, keep only the top 3 (value, column) pairs.
// Enumerate all triplets of distinct rows (r1, r2, r3) and all 3x3x3 column
// combinations, checking for distinct columns.
//
// Time: O(m^3 * 27) = O(m^3), but m ≤ 100 so acceptable for Part I.
// Space: O(m)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: 3x3 board
	board := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(maximumValueSum(board))

	// Example 2: 4x4
	board2 := [][]int{{10, 20, 30, 40}, {50, 60, 70, 80}, {90, 100, 110, 120}, {130, 140, 150, 160}}
	fmt.Println(maximumValueSum(board2))

	// Example 3: 3x4
	board3 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(maximumValueSum(board3))

	// Example 4: 5x3
	board4 := [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}, {-10, -11, -12}, {-13, -14, -15}}
	fmt.Println(maximumValueSum(board4))

	// Example 5: 2x5 (only 2 rows, can't place 3 rooks)
	board5 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	fmt.Println(maximumValueSum(board5))
}

type cell struct {
	val int
	col int
}

func maximumValueSum(board [][]int) int64 {
	m := len(board)
	if m < 3 {
		return 0
	}
	n := len(board[0])
	if n < 3 {
		return 0
	}

	// For each row, find top 3 (value, column) pairs
	rowTop := make([][]cell, m)
	for i := 0; i < m; i++ {
		row := make([]cell, n)
		for j := 0; j < n; j++ {
			row[j] = cell{board[i][j], j}
		}
		sort.Slice(row, func(a, b int) bool {
			return row[a].val > row[b].val
		})
		rowTop[i] = row[:3]
	}

	var ans int64 = -1 << 60

	for r1 := 0; r1 < m; r1++ {
		for r2 := r1 + 1; r2 < m; r2++ {
			for r3 := r2 + 1; r3 < m; r3++ {
				for _, c1 := range rowTop[r1] {
					for _, c2 := range rowTop[r2] {
						if c2.col == c1.col {
							continue
						}
						for _, c3 := range rowTop[r3] {
							if c3.col == c1.col || c3.col == c2.col {
								continue
							}
							sum := int64(c1.val) + int64(c2.val) + int64(c3.val)
							if sum > ans {
								ans = sum
							}
						}
					}
				}
			}
		}
	}

	return ans
}
