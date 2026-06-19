package main

// LeetCode #2883: Drop Missing Data
// https://leetcode.com/problems/drop-missing-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows where the name column is empty.

import "fmt"

func main() {
	// LeetCode name: dropMissingData
	fmt.Println(DropMissingData([][]string{{"1", "Alice", "15"}, {"2", "", "11"}, {"3", "Bob", "12"}}))
	// [[1 Alice 15] [3 Bob 12]]

	fmt.Println(DropMissingData([][]string{{"1", "", "10"}}))
	// []
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropMissingData
func DropMissingData(df [][]string) [][]string {
	result := [][]string{}
	for _, row := range df {
		if row[1] != "" {
			result = append(result, row)
		}
	}
	return result
}
