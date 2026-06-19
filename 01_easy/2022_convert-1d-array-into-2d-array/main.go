package main

// LeetCode #2022: Convert 1D Array Into 2D Array
// https://leetcode.com/problems/convert-1d-array-into-2d-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3, 4}, 2, 2)) // [[1 2] [3 4]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3}, 1, 3))    // [[1 2 3]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2}, 1, 1))       // []
}

// Time: O(m*n), Space: O(m*n)
func ConvertOneDArrayIntoTwoDArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = original[i*n+j]
		}
	}
	return result
}
