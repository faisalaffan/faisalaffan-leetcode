package main

// LeetCode #3948: Lexicographically Maximum MEX Array
// https://leetcode.com/problems/lexicographically-maximum-mex-array/
// Difficulty: Hard
//
// Given array nums, construct a lexicographically maximum array
// result where result[i] is the MEX of a subsequence of nums
// ending at position i (or the MEX after certain operations).
//
// Approach: Track frequency of each value. For each i, find the
// MEX by checking the smallest non-negative integer not in the
// current multiset.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maximumMEXArray([]int{0, 1, 2, 3}))
	// Example 2
	fmt.Println(maximumMEXArray([]int{0, 0, 1, 1}))
	// Edge: empty
	fmt.Println(maximumMEXArray([]int{}))
}

func maximumMEXArray(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	freq := make(map[int]int)
	result := make([]int, n)
	mex := 0

	for i, v := range nums {
		freq[v]++

		// Update mex: find smallest non-negative not in freq
		for freq[mex] > 0 {
			mex++
		}
		result[i] = mex
	}

	return result
}
