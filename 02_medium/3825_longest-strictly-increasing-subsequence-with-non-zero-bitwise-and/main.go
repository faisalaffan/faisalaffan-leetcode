package main

// LeetCode #3825: Longest Strictly Increasing Subsequence With Non-Zero Bitwise AND
// https://leetcode.com/problems/longest-strictly-increasing-subsequence-with-non-zero-bitwise-and/
// Difficulty: Medium
// Time: O(log M * N log N) | Space: O(N)
// Approach: For each bit position, filter nums that have that bit set,
// compute LIS on the filtered list. Take max across all bits.

import (
	"fmt"
	"sort"
)

func LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd(nums []int) int {
	lis := func(arr []int) int {
		tails := []int{}
		for _, x := range arr {
			j := sort.Search(len(tails), func(i int) bool { return tails[i] >= x })
			if j == len(tails) {
				tails = append(tails, x)
			} else {
				tails[j] = x
			}
		}
		return len(tails)
	}

	ans := 0
	// Check up to 31 bits (since nums[i] <= 1e9)
	for bit := 0; bit < 31; bit++ {
		arr := []int{}
		for _, x := range nums {
			if (x>>bit)&1 == 1 {
				arr = append(arr, x)
			}
		}
		if len(arr) > 0 {
			l := lis(arr)
			if l > ans {
				ans = l
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd([]int{5, 4, 7})) // Expected: 2

	// Example 2
	fmt.Println(LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd([]int{2, 3, 6})) // Expected: 3

	// Example 3
	fmt.Println(LongestStrictlyIncreasingSubsequenceWithNonZeroBitwiseAnd([]int{0, 1})) // Expected: 1
}
