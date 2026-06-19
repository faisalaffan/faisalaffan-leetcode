package main

// LeetCode #3098: Find the Sum of Subsequence Powers
// https://leetcode.com/problems/find-the-sum-of-subsequence-powers/
// Difficulty: Hard
// Time: O(n^2 * k + n * D) where D = number of distinct differences
// Space: O(n * k)

import (
	"fmt"
	"sort"
)

const MOD = 1_000_000_007

func sumOfPowers(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	// Collect all distinct adjacent/any differences
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

	// Count k-length subsequences with min_abs_diff >= threshold
	// Returns count of subsequences of length k where min diff >= threshold
	countMinDiffGe := func(threshold int) int {
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, k+1)
		}

		runningSum := make([]int, k+1)
		ptr := 0

		for i := 0; i < n; i++ {
			// Advance ptr while the difference from ptr to i is >= threshold
			// Since nums is sorted, when nums[i]-nums[ptr] >= threshold, then
			// for all future i' > i, the same ptr also qualifies
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

	// Answer = sum over all subsequences of (min_abs_diff)
	// = sum_{d} d * (count(min_diff >= d) - count(min_diff >= d+1))
	// = sum_{d} count(min_diff >= d)

	answer := 0
	prevCount := 0
	for _, d := range diffs {
		currCount := countMinDiffGe(d)
		add := (d % MOD) * ((currCount - prevCount + MOD) % MOD) % MOD
		answer = (answer + add) % MOD
		prevCount = currCount
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
