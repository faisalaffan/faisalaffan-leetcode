package main

// LeetCode #3098: Find the Sum of Subsequence Powers
// https://leetcode.com/problems/find-the-sum-of-subsequence-powers/
// Difficulty: Hard
// Time: O(n^2 * k + D * n * k) | Space: O(n * k)

import (
	"fmt"
	"sort"
)

const MOD = 1_000_000_007

func sumOfPowers(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	// Collect distinct pairwise differences
	diffSet := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diffSet[nums[j]-nums[i]] = true
		}
	}

	diffs := make([]int, 0, len(diffSet))
	for d := range diffSet {
		diffs = append(diffs, d)
	}
	sort.Ints(diffs)

	// f(threshold) = count of k-length subsequences whose min_abs_diff >= threshold
	countGe := func(threshold int) int {
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, k+1)
		}

		runningSum := make([]int, k+1)
		ptr := 0

		for i := 0; i < n; i++ {
			// Two-pointer: advance ptr while difference from ptr to i >= threshold
			for ptr < i && nums[i]-nums[ptr] >= threshold {
				for j := 1; j <= k; j++ {
					runningSum[j] = (runningSum[j] + dp[ptr][j]) % MOD
				}
				ptr++
			}

			dp[i][1] = 1
			for j := 2; j <= k && j <= i+1; j++ {
				dp[i][j] = runningSum[j-1]
			}
		}

		total := 0
		for i := 0; i < n; i++ {
			total = (total + dp[i][k]) % MOD
		}
		return total
	}

	// Answer = sum_{d=1}^{maxDiff} countGe(d)
	// Since countGe is piecewise constant between consecutive distinct diffs,
	// we use: answer += countGe(d_i) * (d_i - d_{i-1})
	answer := 0
	prevDiff := 0
	for _, d := range diffs {
		cnt := countGe(d)
		add := cnt * (d - prevDiff) % MOD
		answer = (answer + add) % MOD
		prevDiff = d
	}

	return answer
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sumOfPowers([]int{1, 2, 3, 4}, 3))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", sumOfPowers([]int{2, 2}, 2))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", sumOfPowers([]int{4, 3, -1}, 2))
	// Expected: 10
}
