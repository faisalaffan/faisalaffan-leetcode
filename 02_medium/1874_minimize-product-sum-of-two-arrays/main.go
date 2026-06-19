package main

// LeetCode #1874: Minimize Product Sum of Two Arrays
// https://leetcode.com/problems/minimize-product-sum-of-two-arrays/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinProductSum([]int{5, 3, 4, 2}, []int{4, 2, 2, 5}))
	fmt.Println(MinProductSum([]int{2, 1, 4, 5, 7}, []int{3, 2, 4, 8, 6}))
}

// Time: O(n log n), Space: O(1)
func MinProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i] * nums2[i]
	}
	return sum
}
