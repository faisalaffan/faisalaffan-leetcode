package main

// LeetCode #2946: Matrix Similarity After Cyclic Shifts
// https://leetcode.com/problems/matrix-similarity-after-cyclic-shifts/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: areSimilar
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 4)) // false
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2)) // true
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{2, 2}, {2, 2}}, 3)) // true
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: areSimilar
func MatrixSimilarityAfterCyclicShifts(mat [][]int, k int) bool {
	for _, row := range mat {
		m := len(row)
		if m == 0 {
			continue
		}
		shift := k % m
		if shift == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			// After left shift by k, row[(j + k) % m] moves to position j
			if row[j] != row[(j+shift)%m] {
				return false
			}
		}
	}
	return true
}
