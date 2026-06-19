package main

// LeetCode #491: Non-decreasing Subsequences
// https://leetcode.com/problems/non-decreasing-subsequences/
// Difficulty: Medium
// Time: O(2^n * n) worst case
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(NonDecreasingSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(NonDecreasingSubsequences([]int{4, 4, 3, 2, 1}))
}

func NonDecreasingSubsequences(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) >= 2 {
			cp := make([]int, len(path))
			copy(cp, path)
			result = append(result, cp)
		}
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				backtrack(i+1, append(path, nums[i]))
			}
		}
	}
	backtrack(0, []int{})
	return result
}
