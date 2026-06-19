package main

// LeetCode #48: Rotate Image
// https://leetcode.com/problems/rotate-image/
// Difficulty: Medium

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	// Transpose
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
	}
}

func main() {
	// Test case 1
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m1)
	fmt.Println(m1) // [[7 4 1] [8 5 2] [9 6 3]]

	// Test case 2
	m2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(m2)
	fmt.Println(m2) // [[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]
}

// Time: O(n^2) | Space: O(1)
