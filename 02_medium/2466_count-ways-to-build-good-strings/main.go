package main

// LeetCode #2466: Count Ways To Build Good Strings
// https://leetcode.com/problems/count-ways-to-build-good-strings/
// Difficulty: Medium
// Time: O(high) | Space: O(high)
// DP: dp[i] = ways to build string of length i.
// dp[i] = dp[i-zero] + dp[i-one] (if i >= zero/one).

import "fmt"

func main() {
	fmt.Println(countGoodStrings(3, 3, 1, 1)) // 8
	fmt.Println(countGoodStrings(2, 3, 1, 2)) // 5
}

const MOD = 1000000007

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	ans := 0

	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % MOD
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % MOD
		}
		if i >= low {
			ans = (ans + dp[i]) % MOD
		}
	}
	return ans
}
