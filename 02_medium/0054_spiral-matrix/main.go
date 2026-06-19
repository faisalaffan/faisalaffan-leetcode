package main

// LeetCode #54: Spiral Matrix
// https://leetcode.com/problems/spiral-matrix/
// Difficulty: Medium

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	top, bottom, left, right := 0, m-1, 0, n-1

	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// [1 2 3 6 9 8 7 4 5]

	// Test case 2
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	// [1 2 3 4 8 12 11 10 9 5 6 7]
}

// Time: O(m*n) | Space: O(1) (excluding output)
