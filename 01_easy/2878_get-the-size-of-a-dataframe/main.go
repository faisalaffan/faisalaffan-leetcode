package main

// LeetCode #2878: Get the Size of a DataFrame
// https://leetcode.com/problems/get-the-size-of-a-dataframe/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we return the dimensions [rows, cols].

import "fmt"

func main() {
	// LeetCode name: getDataframeSize
	fmt.Println(GetTheSizeOfADataframe([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}})) // [4 2]
	fmt.Println(GetTheSizeOfADataframe([][]int{{5, 25}}))                             // [1 2]
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: getDataframeSize
func GetTheSizeOfADataframe(df [][]int) []int {
	if len(df) == 0 {
		return []int{0, 0}
	}
	return []int{len(df), len(df[0])}
}
