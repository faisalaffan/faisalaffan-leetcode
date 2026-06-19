package main

// LeetCode #217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
