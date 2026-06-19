package main

// LeetCode #1580: Put Boxes Into the Warehouse II
// https://leetcode.com/problems/put-boxes-into-the-warehouse-ii/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxBoxesInWarehouseII([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(MaxBoxesInWarehouseII([]int{1, 2, 2, 3, 4}, []int{3, 4, 1, 2}))
	fmt.Println(MaxBoxesInWarehouseII([]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}))
}

func MaxBoxesInWarehouseII(boxes []int, warehouse []int) int {
	// Time: O(N log N + M), Space: O(1)
	// In warehouse II, boxes can enter from either left or right side.
	// We can think of it as: each position's max height is the min of
	// the prefix max from left and prefix max from right.
	sort.Ints(boxes)

	n := len(warehouse)
	// Preprocess: effective height at each position
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = warehouse[0]
	for i := 1; i < n; i++ {
		if warehouse[i] < leftMax[i-1] {
			leftMax[i] = warehouse[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	rightMax[n-1] = warehouse[n-1]
	for i := n - 2; i >= 0; i-- {
		if warehouse[i] < rightMax[i+1] {
			rightMax[i] = warehouse[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	// Effective height = max(leftMax, rightMax) since we can enter from either side
	effective := make([]int, n)
	for i := 0; i < n; i++ {
		if leftMax[i] > rightMax[i] {
			effective[i] = leftMax[i]
		} else {
			effective[i] = rightMax[i]
		}
	}

	// Greedily fit boxes
	boxIdx := 0
	for i := 0; i < n && boxIdx < len(boxes); i++ {
		if boxes[boxIdx] <= effective[i] {
			boxIdx++
		}
	}

	return boxIdx
}
