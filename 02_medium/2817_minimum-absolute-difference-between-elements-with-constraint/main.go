package main

// LeetCode #2817: Minimum Absolute Difference Between Elements With Constraint
// https://leetcode.com/problems/minimum-absolute-difference-between-elements-with-constraint/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"math"
	"sort"
)

func MinimumAbsoluteDifferenceBetweenElementsWithConstraint(nums []int, x int) int {
	n := len(nums)
	if x >= n {
		return -1
	}

	// For each index, we need to find nums[j] with j >= i+x closest to nums[i]
	// Process from right to left, maintaining a sorted set of values
	best := math.MaxInt32

	// Use a sorted slice
	sorted := make([]int, 0)
	for i := n - 1 - x; i >= 0; i-- {
		// Add nums[i+x] to the sorted set
		val := nums[i+x]
		pos := sort.SearchInts(sorted, val)
		sorted = append(sorted, 0)
		copy(sorted[pos+1:], sorted[pos:])
		sorted[pos] = val

		// Find closest to nums[i]
		pos2 := sort.SearchInts(sorted, nums[i])
		if pos2 < len(sorted) {
			if diff := sorted[pos2] - nums[i]; diff < best {
				best = diff
			}
		}
		if pos2 > 0 {
			if diff := nums[i] - sorted[pos2-1]; diff < best {
				best = diff
			}
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumAbsoluteDifferenceBetweenElementsWithConstraint([]int{4, 3, 2, 4}, 2))
	fmt.Println(MinimumAbsoluteDifferenceBetweenElementsWithConstraint([]int{5, 3, 2, 10, 15}, 1))
}
