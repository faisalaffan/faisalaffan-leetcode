package main

// LeetCode #1918: Kth Smallest Subarray Sum
// https://leetcode.com/problems/kth-smallest-subarray-sum/
// Difficulty: Medium [Paid]

import (
	"fmt"
)

func main() {
	fmt.Println(KthSmallestSubarraySum([]int{2, 1, 3}, 4))
	fmt.Println(KthSmallestSubarraySum([]int{3, 3, 3}, 4))
}

// Time: O(n log sum), Space: O(1)
func KthSmallestSubarraySum(nums []int, k int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	totalSum := prefix[n]

	// Binary search on sum value
	left, right := 0, totalSum
	for left < right {
		mid := left + (right-left)/2
		if countSubarraysLE(nums, prefix, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func countSubarraysLE(nums []int, prefix []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		// Find first j where prefix[j+1]-prefix[i] > target
		lo, hi := i, len(nums)-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			sum := prefix[mid+1] - prefix[i]
			if sum <= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		count += lo - i
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
