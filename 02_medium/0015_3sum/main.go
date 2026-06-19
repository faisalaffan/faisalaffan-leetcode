package main

// LeetCode #15: 3Sum
// https://leetcode.com/problems/3sum/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]

	// Test case 2
	fmt.Println(threeSum([]int{0, 1, 1})) // []

	// Test case 3
	fmt.Println(threeSum([]int{0, 0, 0})) // [[0 0 0]]
}

// Time: O(n^2) | Space: O(1) (excluding output)
