package main

// LeetCode #265: Paint House II
// https://leetcode.com/problems/paint-house-ii/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"math"
)

func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	n := len(costs)
	k := len(costs[0])

	// Find min and second min for the first house
	prevMin1, prevMin2 := math.MaxInt32, math.MaxInt32
	prevMinIdx := -1

	for j := 0; j < k; j++ {
		cost := costs[0][j]
		if cost < prevMin1 {
			prevMin2 = prevMin1
			prevMin1 = cost
			prevMinIdx = j
		} else if cost < prevMin2 {
			prevMin2 = cost
		}
	}

	for i := 1; i < n; i++ {
		curMin1, curMin2 := math.MaxInt32, math.MaxInt32
		curMinIdx := -1

		for j := 0; j < k; j++ {
			var cost int
			if j == prevMinIdx {
				cost = costs[i][j] + prevMin2
			} else {
				cost = costs[i][j] + prevMin1
			}

			if cost < curMin1 {
				curMin2 = curMin1
				curMin1 = cost
				curMinIdx = j
			} else if cost < curMin2 {
				curMin2 = cost
			}
		}

		prevMin1, prevMin2 = curMin1, curMin2
		prevMinIdx = curMinIdx
	}

	return prevMin1
}

func main() {
	fmt.Println(minCostII([][]int{{1, 5, 3}, {2, 9, 4}}))
}
