package main

// LeetCode #3824: Minimum K to Reduce Array Within Limit
// https://leetcode.com/problems/minimum-k-to-reduce-array-within-limit/
// Difficulty: Medium
// Time: O(N log M) where M = max(nums) | Space: O(1)
// Approach: Binary search on k. Check if total operations <= k^2.

import "fmt"

func MinimumKToReduceArrayWithinLimit(nums []int) int {
	// Binary search on k
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	lo, hi := 1, maxVal
	ans := maxVal

	for lo <= hi {
		mid := (lo + hi) / 2
		ops := 0
		for _, v := range nums {
			ops += (v + mid - 1) / mid // ceil(v / mid)
		}
		if ops <= mid*mid {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{3, 7, 5})) // Expected: 3

	// Example 2
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{1})) // Expected: 1

	// Example 3
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{10, 10, 10}))
}
