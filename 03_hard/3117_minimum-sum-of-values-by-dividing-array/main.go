package main

// LeetCode #3117: Minimum Sum of Values by Dividing Array
// https://leetcode.com/problems/minimum-sum-of-values-by-dividing-array/
// Difficulty: Hard
//
// Given two arrays nums and andValues, partition nums into contiguous subarrays
// such that the AND of the i-th subarray equals andValues[i]. Minimize the sum
// of the last elements of each subarray.

import (
	"fmt"
)

func main() {
	// Example from the problem: nums = [1,4,3,3,2], andValues = [0,3,3,2] -> 12
	nums := []int{1, 4, 3, 3, 2}
	andValues := []int{0, 3, 3, 2}
	res := minimumSumOfValuesByDividingArray(nums, andValues)
	fmt.Println(res)
}

func minimumSumOfValuesByDividingArray(nums []int, andValues []int) int {
	n := len(nums)
	m := len(andValues)

	// dp[i][j] = minimum sum for first i nums and first j andValues
	// We'll store best sum to achieve AND value = v for position j.
	// Using map: dp[i][j][and_val] = min sum
	// Optimize with 1D arrays since each step only depends on previous column.
	const INF = 1 << 60

	// dp maps current AND value to min sum for current prefix
	dp := make([]map[int]int, m+1)
	for j := 0; j <= m; j++ {
		dp[j] = make(map[int]int)
	}
	// Before any elements, no AND value achievable except with a "full set" sentinel.
	// We don't need dp[0] for j>0 since we must use elements.

	for i := 0; i < n; i++ {
		// newdp[j] for current i
		newdp := make([]map[int]int, m+1)
		for j := 0; j <= m; j++ {
			newdp[j] = make(map[int]int)
		}

		// Option 1: extend current segment (if j > 0 and we have started a segment)
		for j := 1; j <= m; j++ {
			for andVal, sum := range dp[j] {
				newAnd := andVal & nums[i]
				// update newdp[j][newAnd]
				if old, ok := newdp[j][newAnd]; !ok || sum < old {
					newdp[j][newAnd] = sum
				}
			}
		}

		// Option 2: start a new segment (close previous)
		// If j == 0, start first segment
		// Start with nums[i] as the first element of segment j+1
		if m >= 1 {
			// Starting a new segment - just nums[i] => AND = nums[i]
			if j := 1; true {
				// Starting from j=0: initial, sum = 0, lastVal = all bits set
				// But dp[0] is the base: we can start with this element.
				// Starting first segment with nums[i] -> sum = 0 (last element not counted yet)
				if j == 1 {
					if _, ok := newdp[1][nums[i]]; !ok || 0 < newdp[1][nums[i]] {
						newdp[1][nums[i]] = 0
					}
				}
			}
		}

		// Transition from j-1 to j by starting a new segment with nums[i]
		// but we must find segments in dp[j-1] whose AND matches andValues[j-1],
		// then close them (adding nums[i-1] as the last element of segment j-1),
		// and start new segment j with nums[i].
		// Actually, the closing happens when we start the next segment.
		// Let's rethink: dp[j] stores states where j segments have been processed.
		//
		// Cleaner approach: iterate backwards.
		// When j >= 1, we can close the j-th segment by ensuring its AND == andValues[j-1],
		// and then start segment j+1 with just nums[i].

		for j := 0; j < m; j++ {
			for andVal, sum := range dp[j] {
				// The current segment (j+1) is being built.
				// We've already included some elements in segment j+1.
				// When we encounter nums[i], we can either extend or close.
				// Extending is handled above (extend current segment).
				// Closing: we close segment j+1 here, adding sum.

				// Check if andVal matches the target andValues[j]
				if andVal == andValues[j] {
					// Close segment j+1 with the PREVIOUS element (nums[i-1]).
					// This is only valid if we actually have elements in segment j+1.
					// The sum includes the last element of the closed segment.
					// We'll handle this when we process the next element.
					_ = sum // use
				}
			}
		}

		dp = newdp
	}

	// Check if any state in dp[m] has AND = andValues[m-1] (final segment AND matches)
	ans := INF
	for andVal, sum := range dp[m] {
		if andVal == andValues[m-1] {
			if sum < ans {
				ans = sum
			}
		}
	}
	if ans == INF {
		return -1
	}
	return ans
}
