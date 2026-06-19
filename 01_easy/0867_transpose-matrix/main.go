package main

// LeetCode #867: Transpose Matrix
// https://leetcode.com/problems/transpose-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // [[1,4,7],[2,5,8],[3,6,9]]
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))             // [[1,4],[2,5],[3,6]]
}

// transpose returns the transpose of a matrix.
// Time: O(m*n). Space: O(m*n).
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
