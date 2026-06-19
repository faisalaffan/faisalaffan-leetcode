package main

// LeetCode #279: Perfect Squares
// https://leetcode.com/problems/perfect-squares/
// Difficulty: Medium
// Time: O(n * sqrt(n)), Space: O(n)

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			sq := j * j
			if 1+dp[i-sq] < dp[i] {
				dp[i] = 1 + dp[i-sq]
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(1))
}
