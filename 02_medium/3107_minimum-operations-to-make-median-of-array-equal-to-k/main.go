package main

// LeetCode #3107: Minimum Operations to Make Median of Array Equal to K
// https://leetcode.com/problems/minimum-operations-to-make-median-of-array-equal-to-k/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	n := len(nums)
	mid := n / 2
	var ans int64

	ans += abs(int64(nums[mid] - k))
	nums[mid] = k

	for i := mid - 1; i >= 0 && nums[i] > k; i-- {
		ans += abs(int64(nums[i] - k))
		nums[i] = k
	}

	for i := mid + 1; i < n && nums[i] < k; i++ {
		ans += abs(int64(nums[i] - k))
		nums[i] = k
	}

	return ans
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 4)) // Expected: 2
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 7)) // Expected: 3
	fmt.Println(minOperationsToMakeMedianK([]int{1, 2, 3, 4, 5, 6}, 4)) // Expected: 0
}
