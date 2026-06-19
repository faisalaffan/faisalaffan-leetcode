package main

// LeetCode #2969: Minimum Number of Coins for Fruits II
// https://leetcode.com/problems/minimum-number-of-coins-for-fruits-ii/
// Difficulty: Hard [Paid]

import "fmt"

func minimumCoins(prices []int) int {
	n := len(prices)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = int(1e9)
	}
	dp[n] = 0
	dq := make([]int, 0, n)
	dq = append(dq, n)
	for i := n - 1; i >= 0; i-- {
		for len(dq) > 0 && dq[0] > i+i+1 {
			dq = dq[1:]
		}
		dp[i] = prices[i] + dp[dq[0]]
		for len(dq) > 0 && dp[dq[len(dq)-1]] >= dp[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	return dp[0]
}

func main() {
	fmt.Println(minimumCoins([]int{3, 1, 2}))
	fmt.Println(minimumCoins([]int{1, 10, 1, 1}))
}
