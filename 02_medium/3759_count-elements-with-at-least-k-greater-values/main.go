package main

// LeetCode #3759: Count Elements With at Least K Greater Values
// https://leetcode.com/problems/count-elements-with-at-least-k-greater-values/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countElementsWithAtLeastKGreaterValues(nums []int, k int) int {
	if k == 0 {
		return len(nums)
	}
	n := len(nums)
	sort.Ints(nums)
	left := n - k
	// Skip duplicates of threshold value
	for left-1 >= 0 && nums[left-1] == nums[left] {
		left--
	}
	return left
}

func main() {
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{3, 1, 2}, 1))
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{5, 5, 5}, 2))
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{1, 2, 3, 4, 5}, 2))
}
