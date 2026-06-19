package main

// LeetCode #1337: The K Weakest Rows in a Matrix
// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
// Difficulty: Easy
//
// LeetCode submission: func kWeakestRows(mat [][]int, k int) []int

import (
	"fmt"
	"sort"
)

func main() {
	mat1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(TheKWeakestRowsInAMatrix(mat1, 3)) // [2 0 3]

	mat2 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(TheKWeakestRowsInAMatrix(mat2, 2)) // [0 2]
}

// Time: O(m * n + m log m), Space: O(m)
func TheKWeakestRowsInAMatrix(mat [][]int, k int) []int {
	strength := make([][2]int, len(mat))
	for i, row := range mat {
		s := 0
		for _, v := range row {
			if v == 0 {
				break
			}
			s++
		}
		strength[i] = [2]int{s, i}
	}

	sort.Slice(strength, func(i, j int) bool {
		if strength[i][0] != strength[j][0] {
			return strength[i][0] < strength[j][0]
		}
		return strength[i][1] < strength[j][1]
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = strength[i][1]
	}
	return res
}
