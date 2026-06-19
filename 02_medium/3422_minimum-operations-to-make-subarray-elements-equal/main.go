package main

// LeetCode #3422: Minimum Operations to Make Subarray Elements Equal
// https://leetcode.com/problems/minimum-operations-to-make-subarray-elements-equal/
// Difficulty: Medium [Paid]
// Time: O(n log k) Space: O(k)

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(nums []int, k int) int64 {
	n := len(nums)
	if k > n {
		return 0
	}

	var ans int64 = math.MaxInt64
	window := make([]int, k)

	for i := 0; i <= n-k; i++ {
		copy(window, nums[i:i+k])
		sort.Ints(window)
		median := window[k/2]

		var ops int64
		for _, v := range window {
			diff := v - median
			if diff < 0 {
				diff = -diff
			}
			ops += int64(diff)
		}
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{1, 4, 2, 6}, 3)) // 3
	fmt.Println(minOperations([]int{1, 2, 3, 4}, 2)) // 0
}
