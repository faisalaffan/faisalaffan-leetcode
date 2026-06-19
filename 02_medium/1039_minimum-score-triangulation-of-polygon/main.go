package main

// LeetCode #1039: Minimum Score Triangulation of Polygon
// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
// Difficulty: Medium
//
// Approach: Interval DP. dp[i][j] = min score triangulating polygon from i to j.
// Time: O(n^3)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3}))    // 6
	fmt.Println(minScoreTriangulation([]int{3, 7, 4, 5})) // 144
}

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length < n; length++ {
		for i := 0; i+length < n; i++ {
			j := i + length
			dp[i][j] = 1<<31 - 1
			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
				if score < dp[i][j] {
					dp[i][j] = score
				}
			}
		}
	}

	return dp[0][n-1]
}
