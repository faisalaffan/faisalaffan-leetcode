package main

import "fmt"

// LeetCode #446: Arithmetic Slices II - Subsequence
// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/
// Difficulty: Hard
//
// DP with per-index difference maps. dp[i][diff] = number of arithmetic
// subsequences ending at i with common difference diff. For each pair (j,i),
// diff = nums[i]-nums[j]; the subsequences ending at i with that diff are
// dp[j][diff] (extend existing) + 1 (new pair). Sum all valid subsequences
// with length >= 3 (i.e. when dp[j][diff] >= 1).

func main() {
	// Example 1: [2,4,6,8,10] => 7
	fmt.Println("n=7:", numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
	// Example 2: [7,7,7,7,7] => 16
	fmt.Println("n=16:", numberOfArithmeticSlices([]int{7, 7, 7, 7, 7}))
	// Example 3: [0,2000000000,-294967296] => 0
	fmt.Println("n=0:", numberOfArithmeticSlices([]int{0, 2000000000, -294967296}))
	// Edge: short array
	fmt.Println("n=0:", numberOfArithmeticSlices([]int{1, 2}))
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	// dp[i] is a map from difference -> count of subsequences ending at i
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	var total int
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			// count of subsequences ending at j with this diff
			count := dp[j][diff]
			// We add count to dp[i][diff] (extending existing subsequences)
			dp[i][diff] += count
			// Also add the pair (j,i) as a potential start of a new subsequence
			dp[i][diff]++

			// If count >= 1, then extending gives a valid subsequence of length >= 3
			if count >= 1 {
				total += count
			}
		}
	}
	return total
}
