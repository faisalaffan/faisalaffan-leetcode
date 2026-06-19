package main

// LeetCode #1030: Matrix Cells in Distance Order
// https://leetcode.com/problems/matrix-cells-in-distance-order/
// Difficulty: Easy
// Time: O(R*C) | Space: O(R*C)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0)) // [[0,0],[0,1]]
	fmt.Println(allCellsDistOrder(2, 2, 0, 1)) // [[0,1],[0,0],[1,1],[1,0]]
	fmt.Println(allCellsDistOrder(2, 3, 1, 2)) // [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
}

// LeetCode submission: allCellsDistOrder
func allCellsDistOrder(rows, cols, rCenter, cCenter int) [][]int {
	ans := make([][]int, 0, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ans = append(ans, []int{r, c})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		di := abs(ans[i][0]-rCenter) + abs(ans[i][1]-cCenter)
		dj := abs(ans[j][0]-rCenter) + abs(ans[j][1]-cCenter)
		return di < dj
	})
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
