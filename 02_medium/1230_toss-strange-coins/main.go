package main

import (
	"fmt"
)

// LeetCode #1230: Toss Strange Coins
// https://leetcode.com/problems/toss-strange-coins/
// Difficulty: Medium [Paid]

// Probability that exactly target coins land heads.
// dp[j] = probability of j heads after processing i coins.

// Time: O(n * target)
// Space: O(target)

func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := make([]float64, target+1)
	dp[0] = 1.0

	for i := 0; i < n; i++ {
		for j := min(target, i+1); j >= 0; j-- {
			if j > 0 {
				dp[j] = dp[j-1]*prob[i] + dp[j]*(1-prob[i])
			} else {
				dp[0] = dp[0] * (1 - prob[i])
			}
		}
	}

	return dp[target]
}

func main() {
	fmt.Printf("%.5f (expected: 0.40000)\n",
		probabilityOfHeads([]float64{0.4}, 1))

	fmt.Printf("%.5f (expected: 0.40000)\n",
		probabilityOfHeads([]float64{0.5, 0.5, 0.5, 0.5, 0.5}, 0))
}
