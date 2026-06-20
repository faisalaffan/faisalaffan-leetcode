package main

// LeetCode #3729: Count Distinct Subarrays Divisible by K in Sorted Array
// https://leetcode.com/problems/count-distinct-subarrays-divisible-by-k-in-sorted-array/
// Difficulty: Hard
//
// Count distinct subarrays (by value sequence) whose sum is divisible by k.
// Array is sorted, so equal values are contiguous. Subarrays are distinct
// when their sequences of values differ.
//
// Approach: Prefix sum + hash set of subarray content for dedup.

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(numGoodSubarrays([]int{1, 2, 3}, 3))
	// Example 2
	fmt.Println(numGoodSubarrays([]int{1, 1, 2}, 2))
	// Edge: single element
	fmt.Println(numGoodSubarrays([]int{5}, 5))
	// Edge: all same
	fmt.Println(numGoodSubarrays([]int{2, 2, 2}, 2))
}

func numGoodSubarrays(nums []int, k int) int64 {
	n := len(nums)
	pref := make([]int, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + v
	}

	seen := make(map[string]bool)
	var ans int64

	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := i; j < n; j++ {
			if j > i {
				sb.WriteByte(',')
			}
			sb.WriteString(strconv.Itoa(nums[j]))
			if (pref[j+1]-pref[i])%k == 0 {
				key := sb.String()
				if !seen[key] {
					seen[key] = true
					ans++
				}
			}
		}
	}

	return ans
}
