package main

// LeetCode #219: Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[n]; ok && i-j <= k {
			return true
		}
		seen[n] = i
	}
	return false
}

func main() {
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
