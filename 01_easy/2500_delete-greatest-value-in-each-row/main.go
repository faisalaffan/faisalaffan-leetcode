package main

// LeetCode #2500: Delete Greatest Value in Each Row
// https://leetcode.com/problems/delelete-greatest-value-in-each-row/
// Difficulty: Easy
// Time O(n * m log m) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(DeleteGreatestValueInEachRow([][]int{{1, 2, 4}, {3, 3, 1}})) // 8
	fmt.Println(DeleteGreatestValueInEachRow([][]int{{10}}))                 // 10
}

func DeleteGreatestValueInEachRow(grid [][]int) int {
	for i := range grid {
		sort.Ints(grid[i])
	}

	m := len(grid[0])
	sum := 0
	for col := m - 1; col >= 0; col-- {
		maxVal := 0
		for row := 0; row < len(grid); row++ {
			if grid[row][col] > maxVal {
				maxVal = grid[row][col]
			}
		}
		sum += maxVal
	}
	return sum
}
