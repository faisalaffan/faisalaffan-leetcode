package main

// LeetCode #1708: Largest Subarray Length K
// https://leetcode.com/problems/largest-subarray-length-k/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(k) (for output)
func LargestSubarray(nums []int, k int) []int {
	bestIdx := 0
	for i := 1; i <= len(nums)-k; i++ {
		if nums[i] > nums[bestIdx] {
			bestIdx = i
		}
	}
	return nums[bestIdx : bestIdx+k]
}

func main() {
	fmt.Println(LargestSubarray([]int{1, 4, 5, 2, 3}, 3))
	fmt.Println(LargestSubarray([]int{1, 4, 5, 2, 3}, 4))
	fmt.Println(LargestSubarray([]int{1, 2, 3, 4, 5}, 2))
}
