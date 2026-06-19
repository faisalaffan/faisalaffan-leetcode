package main

// LeetCode #3531: Count Covered Buildings
// https://leetcode.com/problems/count-covered-buildings/
// Difficulty: Medium
// Complexity: O(n log n) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	b := [][]int{{1, 5}, {2, 6}, {8, 10}}
	fmt.Println("Test 1:", CountCoveredBuildings(b))
	// Test case 2
	b2 := [][]int{{1, 3}, {4, 6}}
	fmt.Println("Test 2:", CountCoveredBuildings(b2))
	// Test case 3
	b3 := [][]int{{1, 10}, {2, 5}, {3, 8}}
	fmt.Println("Test 3:", CountCoveredBuildings(b3))
}

func CountCoveredBuildings(buildings [][]int) int {
	if len(buildings) == 0 {
		return 0
	}
	// Sort by start, then by end descending
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i][0] != buildings[j][0] {
			return buildings[i][0] < buildings[j][0]
		}
		return buildings[i][1] > buildings[j][1]
	})
	count := 0
	maxEnd := 0
	for _, b := range buildings {
		if b[1] <= maxEnd {
			count++
		} else {
			maxEnd = b[1]
		}
	}
	return count
}
