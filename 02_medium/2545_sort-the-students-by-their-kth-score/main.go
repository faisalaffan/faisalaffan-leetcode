package main

// LeetCode #2545: Sort the Students by Their Kth Score
// https://leetcode.com/problems/sort-the-students-by-their-kth-score/
// Difficulty: Medium
// Time: O(m log m) | Space: O(1) (excluding output)

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sortTheStudents([][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2))
	// Expected: [[7,5,11,2],[10,6,9,1],[4,8,3,15]]

	// Test case 2
	fmt.Println("Test 2:", sortTheStudents([][]int{{3, 4}, {5, 6}}, 0))
	// Expected: [[5,6],[3,4]]

	// Test case 3: single student
	fmt.Println("Test 3:", sortTheStudents([][]int{{1, 2, 3}}, 1))
	// Expected: [[1,2,3]]
}
