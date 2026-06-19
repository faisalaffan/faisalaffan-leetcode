package main

// LeetCode #1402: Reducing Dishes
// https://leetcode.com/problems/reducing-dishes/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	// Sort descending
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})

	curr := 0
	maxVal := 0

	for _, s := range satisfaction {
		curr += s
		if curr > 0 {
			maxVal += curr
		}
	}

	return maxVal
}

func main() {
	// Example 1
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	// Expected: 14

	// Example 2
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	// Expected: 20

	// Example 3
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
	// Expected: 0
}
