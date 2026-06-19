package main

// LeetCode #837: New 21 Game
// https://leetcode.com/problems/new-21-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NewTwoOneGame(10, 1, 10))
	fmt.Println(NewTwoOneGame(6, 1, 10))
	fmt.Println(NewTwoOneGame(21, 17, 10))
}

// Time: O(n) | Space: O(n)
func NewTwoOneGame(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k-1+maxPts {
		return 1.0
	}

	dp := make([]float64, n+1)
	dp[0] = 1.0
	windowSum := 1.0
	var ans float64

	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)
		if i < k {
			windowSum += dp[i]
		} else {
			ans += dp[i]
		}
		if i >= maxPts {
			windowSum -= dp[i-maxPts]
		}
	}

	return ans
}
