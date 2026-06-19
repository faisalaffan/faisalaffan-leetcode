package main

// LeetCode #3779: Minimum Number of Operations to Have Distinct Elements
// https://leetcode.com/problems/minimum-number-of-operations-to-have-distinct-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumNumberOfOperationsToHaveDistinctElements(nums []int) int {
	seen := make(map[int]bool)
	for i := len(nums) - 1; i >= 0; i-- {
		if seen[nums[i]] {
			return (i + 3) / 3
		}
		seen[nums[i]] = true
	}
	return 0
}

func main() {
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{3, 8, 3, 6, 5, 8}))
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{2, 2}))
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{4, 3, 5, 1, 2}))
}
