package main

// LeetCode #566: Reshape the Matrix
// https://leetcode.com/problems/reshape-the-matrix/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(m*n)
func ReshapeTheMatrix(mat [][]int, r, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		result[i/c][i%c] = mat[i/n][i%n]
	}
	return result
}

func main() {
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 2, 4))
}
