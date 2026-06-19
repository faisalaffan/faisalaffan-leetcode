package main

// LeetCode #2971: Find Polygon With the Largest Perimeter
// https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1, 10}))
	fmt.Println(largestPerimeter([]int{1, 2, 3}))
}

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	s := int64(0)
	for _, x := range nums {
		s += int64(x)
	}
	for i := len(nums) - 1; i >= 2; i-- {
		if int64(nums[i]) < s-int64(nums[i]) {
			return s
		}
		s -= int64(nums[i])
	}
	return -1
}
