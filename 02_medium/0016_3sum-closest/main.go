package main

// LeetCode #16: 3Sum Closest
// https://leetcode.com/problems/3sum-closest/
// Difficulty: Medium

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}

	return closest
}

func main() {
	// Test case 1
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // 2

	// Test case 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1)) // 0

	// Test case 3
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100)) // 2
}

// Time: O(n^2) | Space: O(1)
