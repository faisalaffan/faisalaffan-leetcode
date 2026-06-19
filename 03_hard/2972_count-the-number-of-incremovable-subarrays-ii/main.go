package main

// LeetCode #2972: Count the Number of Incremovable Subarrays II
// https://leetcode.com/problems/count-the-number-of-incremovable-subarrays-ii/
// Difficulty: Hard
//
// Count subarrays whose removal leaves the remaining array strictly increasing.
// Equivalent to: find all pairs (l, r) such that nums[0..l-1] and nums[r+1..n-1]
// together form a strictly increasing sequence.
//
// Approach:
// 1. Find the longest strictly increasing suffix starting at position j.
// 2. For each prefix position i, extend the suffix pointer j to maintain
//    that nums[i] < nums[j] (connecting prefix and suffix).
// 3. For each valid i, count subarrays from i+1 to j-1 (any j'-1 where j' >= j).

import (
	"fmt"
	"math"
)

func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)

	// Find the first position of the strictly increasing suffix
	j := n - 1
	for j > 0 && nums[j-1] < nums[j] {
		j--
	}

	// If the whole array is already strictly increasing
	if j == 0 {
		return int64(n * (n + 1) / 2)
	}

	// Subarrays ending at n-1 that can be removed:
	// any prefix subarray [0, k] where k >= j-1
	ans := int64(n - j + 1)

	// Try all possible left boundaries
	prev := math.MinInt
	for i, x := range nums {
		if x <= prev {
			break
		}
		prev = x
		// Move j rightward while nums[j] <= x (breaks the increasing condition)
		for j < n && nums[j] <= x {
			j++
		}
		// Any subarray starting at i+1 and ending at j-1 or later can be removed
		ans += int64(n - j + 1)
	}
	return ans
}

func main() {
	// Example: [1,2,3,4] -> 10 (all subarrays)
	fmt.Println(incremovableSubarrayCount([]int{1, 2, 3, 4}))

	// Reverse order
	fmt.Println(incremovableSubarrayCount([]int{6, 5, 4, 3}))

	// Mixed
	fmt.Println(incremovableSubarrayCount([]int{1, 3, 2, 4}))
	fmt.Println(incremovableSubarrayCount([]int{1, 2, 1, 2}))
}
