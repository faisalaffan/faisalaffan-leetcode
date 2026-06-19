package main

// LeetCode #561: Array Partition
// https://leetcode.com/problems/array-partition/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func ArrayPartition(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(ArrayPartition([]int{1, 4, 3, 2}))
	fmt.Println(ArrayPartition([]int{6, 2, 6, 5, 1, 2}))
}
