package main

// LeetCode #2965: Find Missing and Repeated Values
// https://leetcode.com/problems/find-missing-and-repeated-values/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findMissingAndRepeatedValues
	fmt.Println(FindMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}})) // [2, 4]
	fmt.Println(FindMissingAndRepeatedValues([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}})) // [9, 5]
}

// Time: O(n^2) | Space: O(n^2)
// LeetCode submission name: findMissingAndRepeatedValues
func FindMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	total := n * n
	seen := make(map[int]int)
	repeated := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := grid[i][j]
			seen[val]++
			if seen[val] == 2 {
				repeated = val
			}
		}
	}
	for i := 1; i <= total; i++ {
		if seen[i] == 0 {
			return []int{repeated, i}
		}
	}
	return []int{repeated, 0}
}
