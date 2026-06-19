package main

// LeetCode #915: Partition Array into Disjoint Intervals
// https://leetcode.com/problems/partition-array-into-disjoint-intervals/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{5, 0, 3, 8, 6}))
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{1, 1, 1, 0, 6, 12}))
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{1, 1}))
}

// Time: O(n) | Space: O(1)
func PartitionArrayIntoDisjointIntervals(nums []int) int {
	leftMax, curMax, idx := nums[0], nums[0], 0

	for i, v := range nums {
		if v > curMax {
			curMax = v
		}
		if v < leftMax {
			leftMax = curMax
			idx = i
		}
	}

	return idx + 1
}
