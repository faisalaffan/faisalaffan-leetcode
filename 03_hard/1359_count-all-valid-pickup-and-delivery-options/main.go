package main

// LeetCode #1359: Count All Valid Pickup and Delivery Options
// https://leetcode.com/problems/count-all-valid-pickup-and-delivery-options/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

func countOrders(n int) int {
	// dp[i] = ways to schedule i orders
	// Recurrence: dp[i] = dp[i-1] * C(2i, 2)
	// Explanation: For the i-th order, we have 2i positions.
	// Pick 2 positions for Pi and Di: C(2i, 2) = 2i*(2i-1)/2
	// In exactly half of those, Pi comes before Di.
	// So new_ways = C(2i, 2) = i*(2i-1)
	// dp[i] = dp[i-1] * i * (2*i - 1) % mod

	dp := 1
	for i := 2; i <= n; i++ {
		dp = dp * i % mod * (2*i - 1) % mod
	}
	return dp
}

func main() {
	// Example 1
	fmt.Println(countOrders(1))
	// Expected: 1

	// Example 2
	fmt.Println(countOrders(2))
	// Expected: 6

	// Example 3
	fmt.Println(countOrders(3))
	// Expected: 90
}
