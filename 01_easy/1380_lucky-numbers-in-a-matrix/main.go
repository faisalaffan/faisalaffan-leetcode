package main

// LeetCode #1380: Lucky Numbers in a Matrix
// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
// Difficulty: Easy
//
// LeetCode submission: func luckyNumbers(matrix [][]int) []int

import "fmt"

func main() {
	mat1 := [][]int{
		{3, 7, 8},
		{9, 11, 13},
		{15, 16, 17},
	}
	fmt.Println(LuckyNumbersInAMatrix(mat1)) // [15]

	mat2 := [][]int{
		{1, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 16, 17, 12},
	}
	fmt.Println(LuckyNumbersInAMatrix(mat2)) // [12]
}

// Time: O(m * n), Space: O(m + n)
func LuckyNumbersInAMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	rowMin := make([]int, m)
	for i := range rowMin {
		rowMin[i] = 1<<31 - 1
	}
	colMax := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := matrix[i][j]
			if v < rowMin[i] {
				rowMin[i] = v
			}
			if v > colMax[j] {
				colMax[j] = v
			}
		}
	}
	res := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == rowMin[i] && matrix[i][j] == colMax[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}
