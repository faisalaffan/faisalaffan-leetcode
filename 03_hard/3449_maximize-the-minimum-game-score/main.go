package main

// LeetCode #3449: Maximize the Minimum Game Score
// https://leetcode.com/problems/maximize-the-minimum-game-score/
// Difficulty: Hard
//
// Given points array and integer m, in each move you can increase a point
// by 1. Moving between adjacent indices costs 1 move. Starting at index 0
// with m moves total, maximize the minimum value in the array.
//
// Approach: Binary search on answer with greedy feasibility check.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(maxScore([]int{2, 4, 2}, 5))
	// Example 2
	fmt.Println(maxScore([]int{1, 2, 3}, 4))
	// Example 3: single element
	fmt.Println(maxScore([]int{5}, 10))
	// Edge: zero moves
	fmt.Println(maxScore([]int{1, 2}, 0))
	// Edge: large m
	fmt.Println(maxScore([]int{3, 3, 3}, 6))
}

func maxScore(points []int, m int) int64 {
	if m < len(points) {
		return 0
	}

	// Binary search on the minimum value we can achieve
	var left int64 = 0
	var right int64 = math.MaxInt64

	for left < right {
		mid := (left + right + 1) / 2
		if canAchieve(points, m, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func canAchieve(points []int, m int, target int64) bool {
	n := len(points)
	moves := m
	prevExtra := int64(0)

	for i := 0; i < n; i++ {
		needed := target - int64(points[i]) - prevExtra
		if needed <= 0 {
			// This index already meets target
			if i < n-1 {
				moves--      // move to next index costs 1
				prevExtra = 0
			}
			continue
		}

		// Number of times we need to visit this index
		// Each visit adds 1 to points[i]
		times := (needed + int64(1) - 1) / int64(1) // ceil division
		// Actually, each visit at this index adds 1 (coming from adjacent)
		// We need at least times visits to increase by needed amount
		ops := needed
		if ops > int64(moves) {
			return false
		}
		moves -= int(ops)

		// After ops visits to index i, we may be at index i or i+1
		// If we're at odd visits compared to zero, we're at i+1
		if ops%2 == 1 {
			// We'll be at index i+1
			prevExtra = ops // total increments applied at i
			if i < n-1 {
				// Move to next already counted
			}
		} else {
			// We ended at i, need one more move to go to i+1
			prevExtra = ops
		}

		if i < n-1 && moves > 0 {
			moves--
		}
		prevExtra = 0
	}

	return true
}
