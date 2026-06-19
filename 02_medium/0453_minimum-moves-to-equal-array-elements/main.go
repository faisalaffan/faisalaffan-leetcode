package main

// LeetCode #453: Minimum Moves to Equal Array Elements
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func minMoves(nums []int) int {
	minVal := math.MaxInt32
	sum := 0
	for _, n := range nums {
		sum += n
		if n < minVal {
			minVal = n
		}
	}
	return sum - minVal*len(nums)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves([]int{1, 2, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minMoves([]int{1, 1, 1}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minMoves([]int{1, 1000000000}))
	// Expected: 999999999
}
