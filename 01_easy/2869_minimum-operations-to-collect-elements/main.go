package main

// LeetCode #2869: Minimum Operations to Collect Elements
// https://leetcode.com/problems/minimum-operations-to-collect-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(k)

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 5))
}

func MinimumOperationsToCollectElements(nums []int, k int) int {
	seen := make([]bool, k+1)
	collected := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= 1 && nums[i] <= k && !seen[nums[i]] {
			seen[nums[i]] = true
			collected++
		}
		if collected == k {
			return len(nums) - i
		}
	}
	return len(nums)
}
