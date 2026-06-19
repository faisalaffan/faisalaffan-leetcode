package main

// LeetCode #1959: Minimum Total Space Wasted With K Resizing Operations
// https://leetcode.com/problems/minimum-total-space-wasted-with-k-resizing-operations/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSpaceWastedKResizing([]int{10, 20, 30}, 1))
	fmt.Println(MinSpaceWastedKResizing([]int{10, 20, 15, 30, 20}, 2))
}

const INF = 1 << 30

// Time: O(k * n^2), Space: O(k * n)
func MinSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	k++ // k resizes = k+1 segments

	// g[i][j] = wasted space for segment nums[i..j]
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		s, mx := 0, 0
		for j := i; j < n; j++ {
			s += nums[j]
			if nums[j] > mx {
				mx = nums[j]
			}
			g[i][j] = mx*(j-i+1) - s
		}
	}

	// dp[i][j] = min waste for first i elements with j segments
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for h := 0; h < i; h++ {
				val := dp[h][j-1] + g[h][i-1]
				if val < dp[i][j] {
					dp[i][j] = val
				}
			}
		}
	}
	return dp[n][k]
}
