package main

// LeetCode #2713: Maximum Strictly Increasing Cells in a Matrix
// https://leetcode.com/problems/maximum-strictly-increasing-cells-in-a-matrix/
// Difficulty: Hard
//
// From any cell (r,c) jump to any cell in same row/col with strictly larger value.
// DP with row/col max tracking, processing cells in increasing value order.

import (
	"fmt"
	"sort"
)

func main() {
	mat := [][]int{{3, 1, 6}, {-9, 5, 7}}
	fmt.Println(MaximumStrictlyIncreasingCellsInAMatrix(mat))
}

func MaximumStrictlyIncreasingCellsInAMatrix(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	m, n := len(mat), len(mat[0])

	type cell struct{ r, c int }
	byVal := make(map[int][]cell)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			byVal[mat[i][j]] = append(byVal[mat[i][j]], cell{i, j})
		}
	}

	vals := make([]int, 0, len(byVal))
	for v := range byVal {
		vals = append(vals, v)
	}
	sort.Ints(vals)

	rowMax := make([]int, m)
	colMax := make([]int, n)
	result := 0

	for _, val := range vals {
		cells := byVal[val]
		tmp := make([]int, len(cells))
		for k, c := range cells {
			best := 1
			if rowMax[c.r]+1 > best {
				best = rowMax[c.r] + 1
			}
			if colMax[c.c]+1 > best {
				best = colMax[c.c] + 1
			}
			tmp[k] = best
			if best > result {
				result = best
			}
		}
		for k, c := range cells {
			if tmp[k] > rowMax[c.r] {
				rowMax[c.r] = tmp[k]
			}
			if tmp[k] > colMax[c.c] {
				colMax[c.c] = tmp[k]
			}
		}
	}
	return result
}
