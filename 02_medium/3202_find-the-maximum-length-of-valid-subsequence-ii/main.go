package main

// LeetCode #3202: Find the Maximum Length of Valid Subsequence II
// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/
// Difficulty: Medium
// Time: O(n * k) | Space: O(k)

import "fmt"

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	ans := 0
	for _, v := range nums {
		cur := v % k
		for j := 0; j < k; j++ {
			need := (j - cur%k + k) % k
			dp[cur][j] = dp[need][j] + 1
			if dp[cur][j] > ans {
				ans = dp[cur][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 4, 2, 3, 1, 4}, 3)) // Expected: 4
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 2))     // Expected: 3
}
