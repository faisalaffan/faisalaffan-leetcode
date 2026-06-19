package main

import (
	"fmt"
)

// LeetCode #1272: Remove Interval
// https://leetcode.com/problems/remove-interval/
// Difficulty: Medium [Paid]

// Remove an interval from a set of non-overlapping intervals.
// Return the resulting intervals.

// Time: O(n)
// Space: O(n)

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	result := make([][]int, 0)
	rs, re := toBeRemoved[0], toBeRemoved[1]

	for _, interval := range intervals {
		s, e := interval[0], interval[1]
		// No overlap
		if e <= rs || s >= re {
			result = append(result, interval)
		} else {
			// Left part (if any)
			if s < rs {
				result = append(result, []int{s, rs})
			}
			// Right part (if any)
			if e > re {
				result = append(result, []int{re, e})
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [[0 1] [6 7]])\n",
		removeInterval([][]int{{0, 2}, {3, 4}, {5, 7}}, []int{1, 6}))

	fmt.Printf("%v (expected: [[3 4]])\n",
		removeInterval([][]int{{0, 5}}, []int{0, 3}))
}
