package main

// LeetCode #799: Champagne Tower
// https://leetcode.com/problems/champagne-tower/
// Difficulty: Medium
// Time: O(query_row^2)
// Space: O(query_row)

import "fmt"

func main() {
	fmt.Println(champagneTower(1, 1, 1))
	fmt.Println(champagneTower(2, 1, 1))
	fmt.Println(champagneTower(100000009, 33, 17))
}

func champagneTower(poured int, queryRow int, queryGlass int) float64 {
	dp := make([]float64, queryRow+1)
	dp[0] = float64(poured)

	for row := 0; row < queryRow; row++ {
		next := make([]float64, queryRow+2)
		for col := 0; col <= row; col++ {
			if dp[col] > 1.0 {
				excess := (dp[col] - 1.0) / 2.0
				next[col] += excess
				next[col+1] += excess
			}
		}
		dp = next
	}

	if dp[queryGlass] > 1.0 {
		return 1.0
	}
	return dp[queryGlass]
}
