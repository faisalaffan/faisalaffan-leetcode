package main

// LeetCode #3592: Inverse Coin Change
// https://leetcode.com/problems/inverse-coin-change/
// Difficulty: Medium
// Complexity: O(amount * n) time, O(amount) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", InverseCoinChange([]int{1, 2, 5}, 11))
	// Test case 2
	fmt.Println("Test 2:", InverseCoinChange([]int{2}, 3))
	// Test case 3
	fmt.Println("Test 3:", InverseCoinChange([]int{1}, 0))
}

func InverseCoinChange(coins []int, amount int) int {
	// Minimum number of coins to make amount
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
