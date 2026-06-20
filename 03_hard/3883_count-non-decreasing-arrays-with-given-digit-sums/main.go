package main

// LeetCode #3883: Count Non Decreasing Arrays With Given Digit Sums
// https://leetcode.com/problems/count-non-decreasing-arrays-with-given-digit-sums/
// Difficulty: Hard
//
// Count non-decreasing arrays where each element's digit sum equals
// the corresponding value in digitSum array.
//
// Approach: For each digit sum value, precompute all numbers <= limit
// with that digit sum. Use DP: dp[i][j] = ways for first i positions
// where nums[i] = candidates[i][j]. Transition by prefix sums for
// monotonicity constraint.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countNonDecreasingArrays([]int{1, 2, 1}))
	// Example 2
	fmt.Println(countNonDecreasingArrays([]int{2, 2, 2}))
	// Edge: single element
	fmt.Println(countNonDecreasingArrays([]int{5}))
	// Edge: empty
	fmt.Println(countNonDecreasingArrays([]int{}))
}

const mod = 1000000007
const maxVal = 1000

func countNonDecreasingArrays(digitSum []int) int {
	if len(digitSum) == 0 {
		return 0
	}

	// Precompute numbers grouped by digit sum
	bySum := make([][]int, 55) // max digit sum for numbers <= 1000
	for v := 0; v <= maxVal; v++ {
		s := digitSumOf(v)
		if s < len(bySum) {
			bySum[s] = append(bySum[s], v)
		}
	}

	// cand[i] = candidates for position i
	cand := make([][]int, len(digitSum))
	for i, s := range digitSum {
		if s >= len(bySum) || len(bySum[s]) == 0 {
			return 0
		}
		cand[i] = bySum[s]
	}

	// DP: dp[j] = ways ending with candidates[i][j]
	dp := make([]int, len(cand[0]))
	for j := range dp {
		dp[j] = 1
	}

	for i := 1; i < len(digitSum); i++ {
		ndp := make([]int, len(cand[i]))
		// Build prefix sums of dp
		prefix := make([]int, len(dp)+1)
		for j := 0; j < len(dp); j++ {
			prefix[j+1] = (prefix[j] + dp[j]) % mod
		}
		// For each candidate at position i, add all dp[j] with cand[i-1][j] <= cand[i][k]
		ptr := 0
		for k, val := range cand[i] {
			for ptr < len(cand[i-1]) && cand[i-1][ptr] <= val {
				ptr++
			}
			ndp[k] = prefix[ptr]
		}
		dp = ndp
	}

	ans := 0
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return ans
}

func digitSumOf(v int) int {
	s := 0
	for v > 0 {
		s += v % 10
		v /= 10
	}
	return s
}
