package main

import (
	"fmt"
)

// LeetCode #1144: Decrease Elements To Make Array Zigzag
// https://leetcode.com/problems/decrease-elements-to-make-array-zigzag/
// Difficulty: Medium

// We can make array zigzag in 2 patterns:
// A[0] < A[1] > A[2] < A[3] > ...  (even-indexed are valleys)
// A[0] > A[1] < A[2] > A[3] < ...  (odd-indexed are valleys)
// Compute min total decreases for each pattern.

// Time: O(n)
// Space: O(1)

func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	movesEven := 0 // even indices are valleys (less than neighbors)
	movesOdd := 0  // odd indices are valleys

	for i, v := range nums {
		left := 1001
		if i > 0 {
			left = nums[i-1]
		}
		right := 1001
		if i < n-1 {
			right = nums[i+1]
		}
		neighbor := left
		if right < neighbor {
			neighbor = right
		}
		if i%2 == 0 {
			// even index: should be smaller than neighbors (valley)
			if v >= neighbor {
				movesEven += v - (neighbor - 1)
			}
		} else {
			// odd index: should be smaller than neighbors (valley)
			if v >= neighbor {
				movesOdd += v - (neighbor - 1)
			}
		}
	}

	// For pattern 2: even indices should be peaks, odd indices valleys
	movesEven2 := 0
	movesOdd2 := 0
	for i, v := range nums {
		left := 1001
		if i > 0 {
			left = nums[i-1]
		}
		right := 1001
		if i < n-1 {
			right = nums[i+1]
		}
		neighbor := left
		if right < neighbor {
			neighbor = right
		}
		if i%2 == 0 {
			// even index: should be larger than neighbors (peak)
			if v >= neighbor {
				movesOdd2 += v - (neighbor - 1)
			}
		} else {
			// odd index: should be larger than neighbors (peak)
			if v >= neighbor {
				movesEven2 += v - (neighbor - 1)
			}
		}
	}

	result := movesEven + movesOdd
	if movesEven2+movesOdd2 < result {
		result = movesEven2 + movesOdd2
	}
	return result
}

func main() {
	fmt.Printf("%d (expected: 2)\n", movesToMakeZigzag([]int{1, 2, 3}))
	fmt.Printf("%d (expected: 4)\n", movesToMakeZigzag([]int{9, 6, 1, 6, 2}))
	fmt.Printf("%d (expected: 0)\n", movesToMakeZigzag([]int{1, 3, 2}))
}
