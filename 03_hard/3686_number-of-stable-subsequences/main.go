package main

// LeetCode #3686: Number of Stable Subsequences
// https://leetcode.com/problems/number-of-stable-subsequences/
// Difficulty: Hard
//
// Count subsequences that do NOT contain three consecutive elements with
// the same parity (all odd or all even).
//
// Approach: DP tracking count of subsequences ending in even/odd with
// 1 or 2 consecutive same-parity elements.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countStableSubsequences([]int{1, 3, 5}))
	// Example 2
	fmt.Println(countStableSubsequences([]int{2, 3, 4, 2}))
	// Edge: single element
	fmt.Println(countStableSubsequences([]int{5}))
	// Edge: all odd
	fmt.Println(countStableSubsequences([]int{1, 1, 1}))
}

const MOD = 1000000007

func countStableSubsequences(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// dp[parity][consecutive] = count of subsequences ending with given parity
	// parity: 0=even, 1=odd
	// consecutive: 1 or 2 (number of consecutive same-parity elements at end)
	dp := [2][3]int{}

	result := 0

	for _, v := range nums {
		p := v & 1 // parity: 0=even, 1=odd
		otherP := 1 - p

		// New subsequences ending with this element
		// Option 1: single element
		newSingle := 1

		// Option 2: append to subsequence ending with opposite parity
		// (breaks the consecutive count, resets to 1)
		newAfterOther := (dp[otherP][1] + dp[otherP][2]) % MOD

		// Option 3: append to subsequence ending with same parity with count 1
		// (now consecutive count becomes 2)
		newAfterSame1 := dp[p][1]

		// Can't append to subsequence ending with same parity with count 2
		// (that would make 3 consecutive same parity - not allowed)

		totalNew := (newSingle + newAfterOther + newAfterSame1) % MOD

		// Update dp: shift consecutive=2 to be replaced by new consecutive=2
		dp[p][2] = (dp[p][2] + newAfterSame1) % MOD
		dp[p][1] = (dp[p][1] + newSingle + newAfterOther) % MOD

		result = (result + totalNew) % MOD
	}

	return result
}
