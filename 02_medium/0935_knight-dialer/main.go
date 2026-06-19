package main

// LeetCode #935: Knight Dialer
// https://leetcode.com/problems/knight-dialer/
// Difficulty: Medium

import "fmt"

const mod = 1_000_000_007

// Time: O(n) | Space: O(1)
func knightDialer(n int) int {
	dp := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp = [10]int{
			(dp[4] + dp[6]) % mod,
			(dp[6] + dp[8]) % mod,
			(dp[7] + dp[9]) % mod,
			(dp[4] + dp[8]) % mod,
			(dp[0] + dp[3] + dp[9]) % mod,
			0,
			(dp[0] + dp[1] + dp[7]) % mod,
			(dp[2] + dp[6]) % mod,
			(dp[1] + dp[3]) % mod,
			(dp[2] + dp[4]) % mod,
		}
	}
	sum := 0
	for _, v := range dp {
		sum = (sum + v) % mod
	}
	return sum
}

func main() {
	fmt.Println(knightDialer(1))
	fmt.Println(knightDialer(2))
	fmt.Println(knightDialer(3131))
}
