package main

// LeetCode #3033: Modify the Matrix
// https://leetcode.com/problems/modify-the-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: modifiedMatrix
	fmt.Println(ModifyTheMatrix([][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}}))
	// [[1 2 9] [4 8 6] [7 8 9]]
}

// Time: O(n * m) | Space: O(1) (modifies in place)
// LeetCode submission name: modifiedMatrix
func ModifyTheMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	m, n := len(matrix), len(matrix[0])

	// Find max in each column
	colMax := make([]int, n)
	for j := 0; j < n; j++ {
		maxVal := -1
		for i := 0; i < m; i++ {
			if matrix[i][j] > maxVal {
				maxVal = matrix[i][j]
			}
		}
		colMax[j] = maxVal
	}

	// Replace -1 values
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = colMax[j]
			}
		}
	}
	return matrix
}
