package main

// LeetCode #311: Sparse Matrix Multiplication
// https://leetcode.com/problems/sparse-matrix-multiplication/
// Difficulty: Medium [Paid]
// Time: O(m*n*k), Space: O(m*k) optimized for sparse matrices

import "fmt"

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k, n := len(mat1), len(mat1[0]), len(mat2[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for kk := 0; kk < k; kk++ {
			if mat1[i][kk] != 0 {
				for j := 0; j < n; j++ {
					if mat2[kk][j] != 0 {
						result[i][j] += mat1[i][kk] * mat2[kk][j]
					}
				}
			}
		}
	}

	return result
}

func main() {
	mat1 := [][]int{{1, 0, 0}, {-1, 0, 3}}
	mat2 := [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	fmt.Println(multiply(mat1, mat2))

	mat1 = [][]int{{0}}
	mat2 = [][]int{{0}}
	fmt.Println(multiply(mat1, mat2))

	mat1 = [][]int{{1, -5}}
	mat2 = [][]int{{12}, {-1}}
	fmt.Println(multiply(mat1, mat2))
}
