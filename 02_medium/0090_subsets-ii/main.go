package main

// LeetCode #90: Subsets II
// https://leetcode.com/problems/subsets-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	start := 0

	for i := 0; i < len(nums); i++ {
		n := len(result)
		begin := 0
		if i > 0 && nums[i] == nums[i-1] {
			begin = start
		}
		start = n
		for j := begin; j < n; j++ {
			newSubset := make([]int, len(result[j])+1)
			copy(newSubset, result[j])
			newSubset[len(result[j])] = nums[i]
			result = append(result, newSubset)
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	// [[] [1] [2] [1 2] [2 2] [1 2 2]]

	// Test case 2
	fmt.Println(subsetsWithDup([]int{0})) // [[] [0]]
}

// Time: O(n * 2^n) | Space: O(n * 2^n)
