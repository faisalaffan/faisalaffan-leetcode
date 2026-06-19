package main

// LeetCode #2826: Sorting Three Groups
// https://leetcode.com/problems/sorting-three-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func SortingThreeGroups(nums []int) int {
	n := len(nums)
	// dp[i][j] = min operations to make first i+1 elements sorted with last element == j+1
	dp := make([][3]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 0; j < 3; j++ {
			change := 0
			if nums[i-1] != j+1 {
				change = 1
			}
			dp[i][j] = dp[i-1][j] + change
			if j > 0 && dp[i][j-1] < dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[n][2]
}

func main() {
	fmt.Println(SortingThreeGroups([]int{2, 1, 3, 2, 1}))
	fmt.Println(SortingThreeGroups([]int{1, 2, 3, 1, 2, 3}))
}
