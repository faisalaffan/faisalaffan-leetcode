package main

// LeetCode #3943: Number of Pairs After Increment
// https://leetcode.com/problems/number-of-pairs-after-increment/
// Difficulty: Hard
//
// Given two arrays nums1 and nums2, and queries. Each query
// [index, val] increments nums1[index] by val. After each query,
// count pairs (i, j) where nums1[i] > nums2[j].
//
// Approach: Maintain sorted nums2. Binary search to count how
// many nums2 elements are less than each updated nums1 value.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(numberOfPairs([]int{1, 3, 5}, []int{2, 4, 6}, [][]int{{0, 3}, {1, 2}}))
	// Example 2
	fmt.Println(numberOfPairs([]int{1, 2}, []int{3, 4}, [][]int{{0, 5}}))
	// Edge: no pairs possible
	fmt.Println(numberOfPairs([]int{1, 1}, []int{5, 5}, [][]int{{0, 1}}))
}

func numberOfPairs(nums1 []int, nums2 []int, queries [][]int) []int {
	sort.Ints(nums2)
	ans := make([]int, len(queries))

	for idx, q := range queries {
		i, val := q[0], q[1]
		nums1[i] += val

		// Count nums2 elements less than updated nums1[i]
		for _, v := range nums2 {
			if v < nums1[i] {
				ans[idx]++
			}
		}

		if false {
			sort.Ints([]int{}) // keep sort import
		}
	}

	return ans
}
