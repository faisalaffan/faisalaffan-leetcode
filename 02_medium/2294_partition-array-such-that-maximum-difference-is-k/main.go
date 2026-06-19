package main

// LeetCode #2294: Partition Array Such That Maximum Difference Is K
// https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	partitions := 1
	minVal := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-minVal > k {
			partitions++
			minVal = nums[i]
		}
	}
	return partitions
}

func main() {
	// Test case 1
	fmt.Println(partitionArray([]int{3, 6, 1, 2, 5}, 2))
	// Expected: 2

	// Test case 2
	fmt.Println(partitionArray([]int{1, 2, 3}, 1))
	// Expected: 2

	// Test case 3
	fmt.Println(partitionArray([]int{2, 2, 4, 5}, 0))
	// Expected: 3
}
