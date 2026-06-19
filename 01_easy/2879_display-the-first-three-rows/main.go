package main

// LeetCode #2879: Display the First Three Rows
// https://leetcode.com/problems/display-the-first-three-rows/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we return the first 3 rows of the dataframe.

import "fmt"

func main() {
	// LeetCode name: selectFirstRows
	fmt.Println(DisplayTheFirstThreeRows([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}}))
	// [[1 15] [2 11] [3 11]]

	fmt.Println(DisplayTheFirstThreeRows([][]int{{1, 15}}))
	// [[1 15]]
}

// Time: O(n) where n = min(3, rows) | Space: O(1)
// LeetCode submission name: selectFirstRows
func DisplayTheFirstThreeRows(df [][]int) [][]int {
	if len(df) > 3 {
		return df[:3]
	}
	return df
}
