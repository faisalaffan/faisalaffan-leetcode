package main

// LeetCode #2811: Check if it is Possible to Split Array
// https://leetcode.com/problems/check-if-it-is-possible-to-split-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func CheckIfItIsPossibleToSplitArray(nums []int, m int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if nums[i]+nums[i+1] >= m {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{2, 3, 3, 2, 3}, 6))
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{1, 1}, 3))
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{1, 2, 1}, 4))
}
