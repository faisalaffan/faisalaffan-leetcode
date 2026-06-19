package main

// LeetCode #2555: Maximize Win From Two Segments
// https://leetcode.com/problems/maximize-win-from-two-segments/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	// dp[i] = max prizes we can win with one segment ending at or before position i
	dp := make([]int, n+1)
	ans := 0

	left := 0
	for right := 0; right < n; right++ {
		for prizePositions[right]-prizePositions[left] > k {
			left++
		}
		// Current segment [left, right] covers `right-left+1` prizes
		curr := right - left + 1
		// Best with one segment up to position before left
		dp[right+1] = max(dp[right], curr)
		// Combine with best segment before current segment
		ans = max(ans, curr+dp[left])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximizeWin([]int{1, 1, 2, 2, 3, 3, 5}, 2))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", maximizeWin([]int{1, 2, 3, 4}, 0))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", maximizeWin([]int{1, 2, 3, 4, 5, 6}, 2))
	// Expected: 4
}
