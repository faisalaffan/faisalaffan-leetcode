package main

// LeetCode #2765: Longest Alternating Subarray
// https://leetcode.com/problems/longest-alternating-subarray/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(LongestAlternatingSubarray([]int{2, 3, 4, 3, 4}))
	fmt.Println(LongestAlternatingSubarray([]int{4, 5, 6}))
}

func LongestAlternatingSubarray(nums []int) int {
	n := len(nums)
	ans := -1
	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			continue
		}
		length := 2
		expected := -1 // next diff should be -1
		for j := i + 2; j < n; j++ {
			diff := nums[j] - nums[j-1]
			if diff != expected {
				break
			}
			length++
			expected = -expected
		}
		if length > ans {
			ans = length
		}
	}
	return ans
}
