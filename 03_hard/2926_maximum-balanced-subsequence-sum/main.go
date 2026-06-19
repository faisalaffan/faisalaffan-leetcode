package main

// LeetCode #2926: Maximum Balanced Subsequence Sum
// https://leetcode.com/problems/maximum-balanced-subsequence-sum/
// Difficulty: Hard
//
// Approach: BIT (Fenwick tree) with coordinate compression.
// Transform: balanced condition nums[j] - nums[i] >= j - i
// is equivalent to nums[j] - j >= nums[i] - i.
// Define key[i] = nums[i] - i. We need a subsequence with non-decreasing keys.
// DP + BIT: dp[i] = nums[i] + max(dp[j]) for j < i with key[j] <= key[i].

import (
	"fmt"
	"math"
	"sort"
)

func maxBalancedSubsequenceSum(nums []int) int64 {
	n := len(nums)
	keys := make([]int, n)
	for i, v := range nums {
		keys[i] = v - i
	}

	// Coordinate compression
	sorted := make([]int, n)
	copy(sorted, keys)
	sort.Ints(sorted)
	m := 1
	for i := 1; i < n; i++ {
		if sorted[i] != sorted[m-1] {
			sorted[m] = sorted[i]
			m++
		}
	}
	sorted = sorted[:m]

	// BIT for prefix maximum
	bit := make([]int64, m+2)
	for i := range bit {
		bit[i] = math.MinInt64
	}

	query := func(pos int) int64 {
		res := int64(math.MinInt64)
		for pos > 0 {
			if bit[pos] > res {
				res = bit[pos]
			}
			pos -= pos & -pos
		}
		return res
	}
	update := func(pos int, val int64) {
		for pos <= m {
			if val > bit[pos] {
				bit[pos] = val
			}
			pos += pos & -pos
		}
	}

	var ans int64 = int64(nums[0])
	for i, v := range nums {
		pos := sort.SearchInts(sorted, keys[i]) + 1 // 1-indexed BIT
		best := query(pos)
		if best == math.MinInt64 {
			best = 0
		}
		cur := best + int64(v)
		if cur > ans {
			ans = cur
		}
		update(pos, cur)
	}
	return ans
}

func main() {
	// Example: [3,3,5,6] -> 14 (subsequence [3,5,6])
	fmt.Println(maxBalancedSubsequenceSum([]int{3, 3, 5, 6}))

	// All negative: pick the max single element
	fmt.Println(maxBalancedSubsequenceSum([]int{-5, -3, -1}))

	// Mixed
	fmt.Println(maxBalancedSubsequenceSum([]int{5, -10, 3}))
}
