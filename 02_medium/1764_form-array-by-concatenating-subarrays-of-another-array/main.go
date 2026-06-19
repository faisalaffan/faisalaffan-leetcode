package main

// LeetCode #1764: Form Array by Concatenating Subarrays of Another Array
// https://leetcode.com/problems/form-array-by-concatenating-subarrays-of-another-array/
// Difficulty: Medium
// Time: O(n * m), Space: O(1)

import "fmt"

func canChoose(groups [][]int, nums []int) bool {
	idx := 0
	for _, group := range groups {
		found := false
		for idx <= len(nums)-len(group) {
			if matches(nums, group, idx) {
				idx += len(group)
				found = true
				break
			}
			idx++
		}
		if !found {
			return false
		}
	}
	return true
}

func matches(nums []int, group []int, start int) bool {
	for i, v := range group {
		if nums[start+i] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canChoose([][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0})) // Expected: true
	fmt.Println(canChoose([][]int{{10, -2}, {1, 2, 3, 4}}, []int{1, 2, 3, 4, 10, -2})) // Expected: false
	fmt.Println(canChoose([][]int{{1, 2, 3}, {3, 4}}, []int{7, 7, 1, 2, 3, 4, 7, 7})) // Expected: false
}
