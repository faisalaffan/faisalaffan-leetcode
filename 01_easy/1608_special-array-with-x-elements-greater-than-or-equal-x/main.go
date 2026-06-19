package main

// LeetCode #1608: Special Array With X Elements Greater Than or Equal X
// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func SpecialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if nums[i] >= i+1 {
			if i == len(nums)-1 || nums[i+1] < i+1 {
				return i + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(SpecialArray([]int{3, 5}))
	fmt.Println(SpecialArray([]int{0, 0}))
	fmt.Println(SpecialArray([]int{0, 4, 3, 0, 4}))
}
