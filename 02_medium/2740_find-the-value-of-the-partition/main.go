package main

// LeetCode #2740: Find the Value of the Partition
// https://leetcode.com/problems/find-the-value-of-the-partition/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func FindTheValueOfThePartition(nums []int) int {
	sort.Ints(nums)
	minDiff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if diff := nums[i] - nums[i-1]; diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func main() {
	fmt.Println(FindTheValueOfThePartition([]int{1, 3, 2, 4}))
	fmt.Println(FindTheValueOfThePartition([]int{100, 1, 10}))
}
