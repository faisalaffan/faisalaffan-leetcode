package main

// LeetCode #976: Largest Perimeter Triangle
// https://leetcode.com/problems/largest-perimeter-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))          // 5
	fmt.Println(largestPerimeter([]int{1, 2, 1}))          // 0
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))       // 8
}

// largestPerimeter finds the largest perimeter of a triangle from the given side lengths.
// Time: O(n log n). Space: O(log n).
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}
