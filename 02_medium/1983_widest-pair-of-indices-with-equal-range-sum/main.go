package main

// LeetCode #1983: Widest Pair of Indices With Equal Range Sum
// https://leetcode.com/problems/widest-pair-of-indices-with-equal-range-sum/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 2, 3}))
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{1, 1, 1}, []int{1, 1, 1}))
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{0, 1}, []int{1, 0}))
}

// Time: O(n), Space: O(n)
func WidestPairOfIndicesWithEqualRangeSum(nums1 []int, nums2 []int) int {
	first := make(map[int]int)
	first[0] = -1
	maxWidth := 0
	prefix1, prefix2 := 0, 0

	for i := 0; i < len(nums1); i++ {
		prefix1 += nums1[i]
		prefix2 += nums2[i]
		diff := prefix1 - prefix2

		if idx, ok := first[diff]; ok {
			if i-idx > maxWidth {
				maxWidth = i - idx
			}
		} else {
			first[diff] = i
		}
	}

	return maxWidth
}
