package main

// LeetCode #3496: Maximize Score After Pair Deletions
// https://leetcode.com/problems/maximize-score-after-pair-deletions/
// Difficulty: Medium [Paid]
// Complexity: O(n^2) time, O(n^2) space

import "fmt"

func main() {
	// Test case 1
	nums := []int{1, 2, 3, 4}
	cost := []int{1, 2, 3, 4}
	fmt.Println("Test 1:", MaximizeScoreAfterPairDeletions(nums, cost))

	// Test case 2
	nums2 := []int{5, 1, 5, 1}
	cost2 := []int{10, 1, 10, 1}
	fmt.Println("Test 2:", MaximizeScoreAfterPairDeletions(nums2, cost2))

	// Test case 3
	nums3 := []int{1, 2}
	cost3 := []int{3, 4}
	fmt.Println("Test 3:", MaximizeScoreAfterPairDeletions(nums3, cost3))
}

func MaximizeScoreAfterPairDeletions(nums []int, cost []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	// dp[l][r] = max score for subarray nums[l..r]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: single element
	for i := 0; i < n; i++ {
		dp[i][i] = cost[i]
	}

	// Base case: pair
	for i := 0; i+1 < n; i++ {
		dp[i][i+1] = cost[i] + cost[i+1]
	}

	for length := 3; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			maxScore := 0
			for k := l; k < r; k++ {
				score := dp[l][k] + dp[k+1][r]
				if score > maxScore {
					maxScore = score
				}
			}
			dp[l][r] = maxScore
		}
	}

	return dp[0][n-1]
}
