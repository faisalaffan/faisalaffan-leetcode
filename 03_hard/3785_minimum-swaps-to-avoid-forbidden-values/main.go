package main

// LeetCode #3785: Minimum Swaps to Avoid Forbidden Values
// https://leetcode.com/problems/minimum-swaps-to-avoid-forbidden-values/
// Difficulty: Hard
//
// Given array nums and list of forbidden values, find min swaps
// so that every prefix position has a non-forbidden value.
//
// Approach: Count forbidden values in prefix. Each swap can fix
// at most 2 positions. Answer = ceil(badPrefix / 2).

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumSwaps([]int{1, 2, 3, 4}, []int{1, 4}))
	// Example 2
	fmt.Println(minimumSwaps([]int{2, 1, 3}, []int{1}))
	// Edge: no forbidden
	fmt.Println(minimumSwaps([]int{1, 2, 3}, []int{}))
	// Edge: all forbidden
	fmt.Println(minimumSwaps([]int{1, 1, 1}, []int{1}))
}

func minimumSwaps(nums []int, forbidden []int) int {
	forbid := make(map[int]bool)
	for _, v := range forbidden {
		forbid[v] = true
	}

	// Count forbidden values not in correct position
	bad := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if forbid[nums[i]] {
			bad++
		}
	}

	return (bad + 1) / 2
}
