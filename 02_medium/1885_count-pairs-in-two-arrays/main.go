package main

// LeetCode #1885: Count Pairs in Two Arrays
// https://leetcode.com/problems/count-pairs-in-two-arrays/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CountPairs([]int{1, 3, 4}, []int{1, 3, 4}))
	fmt.Println(CountPairs([]int{1, 2, 3}, []int{2, 3, 4}))
	fmt.Println(CountPairs([]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}))
}

// Time: O(n log n), Space: O(n)
func CountPairs(nums1 []int, nums2 []int) int {
	n := len(nums1)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = nums1[i] - nums2[i]
	}
	sort.Ints(diff)

	left, right := 0, n-1
	count := 0
	for left < right {
		if diff[left]+diff[right] > 0 {
			count += right - left
			right--
		} else {
			left++
		}
	}
	return count
}
