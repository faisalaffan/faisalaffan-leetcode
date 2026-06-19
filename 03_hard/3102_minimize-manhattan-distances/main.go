package main

// LeetCode #3102: Minimize Manhattan Distances
// https://leetcode.com/problems/minimize-manhattan-distances/
// Difficulty: Hard
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func minimumDistance(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return 0
	}

	// Track top-2 max and min for u = x+y and v = x-y
	// We need values AND indices
	max1U, max2U := math.MinInt32, math.MinInt32
	min1U, min2U := math.MaxInt32, math.MaxInt32
	max1V, max2V := math.MinInt32, math.MinInt32
	min1V, min2V := math.MaxInt32, math.MaxInt32
	idxMax1U, idxMin1U := -1, -1
	idxMax1V, idxMin1V := -1, -1

	for i, p := range points {
		u := p[0] + p[1]
		v := p[0] - p[1]

		// u max
		if u > max1U {
			max2U = max1U
			max1U = u
			idxMax1U = i
		} else if u > max2U {
			max2U = u
		}

		// u min
		if u < min1U {
			min2U = min1U
			min1U = u
			idxMin1U = i
		} else if u < min2U {
			min2U = u
		}

		// v max
		if v > max1V {
			max2V = max1V
			max1V = v
			idxMax1V = i
		} else if v > max2V {
			max2V = v
		}

		// v min
		if v < min1V {
			min2V = min1V
			min1V = v
			idxMin1V = i
		} else if v < min2V {
			min2V = v
		}
	}

	// Try removing each of the 4 candidate points (extreme ones)
	candidates := make(map[int]bool)
	candidates[idxMax1U] = true
	candidates[idxMin1U] = true
	candidates[idxMax1V] = true
	candidates[idxMin1V] = true

	result := math.MaxInt32
	for idx := range candidates {
		// Compute max_u without idx
		maxU := max1U
		if idx == idxMax1U {
			maxU = max2U
		}
		minU := min1U
		if idx == idxMin1U {
			minU = min2U
		}
		maxV := max1V
		if idx == idxMax1V {
			maxV = max2V
		}
		minV := min1V
		if idx == idxMin1V {
			minV = min2V
		}
		dist := max(maxU-minU, maxV-minV)
		if dist < result {
			result = dist
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumDistance([][]int{{3, 10}, {5, 15}, {1, 5}, {2, 2}, {4, 4}}))
	// Expected: 9 (remove point [5,15], remaining max distance = 9)

	// Test case 2 - from LeetCode
	fmt.Println("Test 2:", minimumDistance([][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", minimumDistance([][]int{{1, 1}, {1, 1}, {1, 1}}))
	// Expected: 0
}
