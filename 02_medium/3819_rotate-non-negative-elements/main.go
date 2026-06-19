package main

// LeetCode #3819: Rotate Non Negative Elements
// https://leetcode.com/problems/rotate-non-negative-elements/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Collect non-negative elements, rotate left by k cyclically, place back.

import "fmt"

func RotateNonNegativeElements(nums []int, k int) []int {
	n := len(nums)
	type pair struct {
		idx int
		val int
	}
	var nonNeg []pair
	for i, v := range nums {
		if v >= 0 {
			nonNeg = append(nonNeg, pair{i, v})
		}
	}

	m := len(nonNeg)
	if m == 0 {
		res := make([]int, n)
		copy(res, nums)
		return res
	}

	result := make([]int, n)
	copy(result, nums)

	// For each position that had a non-negative, put the rotated value
	for i, p := range nonNeg {
		srcIdx := (i + k) % m
		result[p.idx] = nonNeg[srcIdx].val
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(RotateNonNegativeElements([]int{1, -2, 3, -4}, 3)) // Expected: [3 -2 1 -4]

	// Example 2
	fmt.Println(RotateNonNegativeElements([]int{-3, -2, 7}, 1)) // Expected: [-3 -2 7]

	// Example 3
	fmt.Println(RotateNonNegativeElements([]int{5, 4, -9, 6}, 2)) // Expected: [6 5 -9 4]
}
