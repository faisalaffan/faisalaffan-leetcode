package main

// LeetCode #3077: Maximum Strength of K Disjoint Subarrays
// https://leetcode.com/problems/maximum-strength-of-k-disjoint-subarrays/
// Difficulty: Hard
// Time: O(n*k) | Space: O(k)

import (
	"fmt"
	"math"
)

func MaximumStrength(nums []int, k int) int64 {
	n := len(nums)

	// dp0[j] = max strength with j subarrays, NOT using current element
	// dp1[j] = max strength with j subarrays, ENDING at current element
	dp0 := make([]int64, k+1)
	dp1 := make([]int64, k+1)

	// Weight for the j-th subarray (1-indexed): (-1)^(j+1) * (k-j+1)
	weight := func(j int) int64 {
		w := int64(k - j + 1)
		if j%2 == 0 {
			return -w
		}
		return w
	}

	// Initialize with -inf
	negInf := int64(math.MinInt64 / 2)
	for j := 0; j <= k; j++ {
		dp0[j] = negInf
		dp1[j] = negInf
	}
	dp0[0] = 0

	for i := 0; i < n; i++ {
		ndp0 := make([]int64, k+1)
		ndp1 := make([]int64, k+1)
		for j := 0; j <= k; j++ {
			ndp0[j] = negInf
			ndp1[j] = negInf
		}

		for j := 0; j <= k; j++ {
			// Not using nums[i]: carry forward best state without nums[i]
			ndp0[j] = max(ndp0[j], dp0[j])
			ndp0[j] = max(ndp0[j], dp1[j])

			if j > 0 {
				w := weight(j)
				// Start new subarray at nums[i]
				bestPrev := max(dp0[j-1], dp1[j-1])
				ndp1[j] = max(ndp1[j], bestPrev+w*int64(nums[i]))

				// Extend current subarray to include nums[i]
				ndp1[j] = max(ndp1[j], dp1[j]+w*int64(nums[i]))
			}
		}

		dp0, dp1 = ndp0, ndp1
	}

	return max(dp0[k], dp1[k])
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumStrength([]int{1, 2, 3, -1, 2}, 3))
	// Expected: 22

	// Test case 2
	fmt.Println("Test 2:", MaximumStrength([]int{12, -2, -2, -2, -2}, 5))
	// Expected: 64

	// Test case 3
	fmt.Println("Test 3:", MaximumStrength([]int{-1, -2, -3}, 1))
	// Expected: -1
}
