package main

// LeetCode #628: Maximum Product of Three Numbers
// https://leetcode.com/problems/maximum-product-of-three-numbers/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func MaximumProductOfThreeNumbers(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// The max product is either the three largest numbers
	// or two smallest (negative) numbers times the largest.
	p1 := nums[n-1] * nums[n-2] * nums[n-3]
	p2 := nums[0] * nums[1] * nums[n-1]
	if p1 > p2 {
		return p1
	}
	return p2
}

func main() {
	fmt.Println(MaximumProductOfThreeNumbers([]int{1, 2, 3}))
	fmt.Println(MaximumProductOfThreeNumbers([]int{1, 2, 3, 4}))
	fmt.Println(MaximumProductOfThreeNumbers([]int{-1, -2, -3}))
}
