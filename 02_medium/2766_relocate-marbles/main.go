package main

// LeetCode #2766: Relocate Marbles
// https://leetcode.com/problems/relocate-marbles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	positions := make(map[int]bool)
	for _, n := range nums {
		positions[n] = true
	}

	for i := 0; i < len(moveFrom); i++ {
		delete(positions, moveFrom[i])
		positions[moveTo[i]] = true
	}

	result := make([]int, 0, len(positions))
	for p := range positions {
		result = append(result, p)
	}
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(RelocateMarbles([]int{1, 2, 3}, []int{1}, []int{4}))
	fmt.Println(RelocateMarbles([]int{1, 1, 2, 2}, []int{1, 2}, []int{3, 4}))
}
