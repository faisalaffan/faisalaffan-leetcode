package main

import (
	"fmt"
)

// LeetCode #1248: Count Number of Nice Subarrays
// https://leetcode.com/problems/count-number-of-nice-subarrays/
// Difficulty: Medium

// Count subarrays with exactly k odd numbers.
// Convert to atMost(k) - atMost(k-1).

// Time: O(n)
// Space: O(1)

func numberOfSubarrays(nums []int, k int) int {
	atMost := func(target int) int {
		if target < 0 {
			return 0
		}
		left, count, result := 0, 0, 0
		for right := 0; right < len(nums); right++ {
			if nums[right]%2 == 1 {
				count++
			}
			for count > target {
				if nums[left]%2 == 1 {
					count--
				}
				left++
			}
			result += right - left + 1
		}
		return result
	}

	return atMost(k) - atMost(k-1)
}

func main() {
	fmt.Printf("%d (expected: 2)\n", numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	fmt.Printf("%d (expected: 0)\n", numberOfSubarrays([]int{2, 4, 6}, 1))
	fmt.Printf("%d (expected: 16)\n", numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
