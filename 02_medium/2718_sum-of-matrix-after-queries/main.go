package main

// LeetCode #2718: Sum of Matrix After Queries
// https://leetcode.com/problems/sum-of-matrix-after-queries/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func SumOfMatrixAfterQueries(n int, queries [][]int) int64 {
	rowSet := make(map[int]bool)
	colSet := make(map[int]bool)
	rowSum := int64(0)
	colSum := int64(0)

	var result int64
	for i := len(queries) - 1; i >= 0; i-- {
		typ, idx, val := queries[i][0], queries[i][1], int64(queries[i][2])
		if typ == 0 { // row
			if rowSet[idx] {
				continue
			}
			rowSet[idx] = true
			rowVal := val * int64(n-len(colSet))
			result += rowVal - rowSum
		} else { // col
			if colSet[idx] {
				continue
			}
			colSet[idx] = true
			colVal := val * int64(n-len(rowSet))
			result += colVal - colSum
		}
	}
	return result
}

func main() {
	fmt.Println(SumOfMatrixAfterQueries(3, [][]int{{0, 0, 1}, {1, 2, 2}, {0, 2, 3}, {1, 0, 4}}))
	fmt.Println(SumOfMatrixAfterQueries(1, [][]int{{0, 0, 5}}))
}
