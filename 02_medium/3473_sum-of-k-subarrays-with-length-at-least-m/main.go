package main

// LeetCode #3473: Sum of K Subarrays With Length at Least M
// https://leetcode.com/problems/sum-of-k-subarrays-with-length-at-least-m/
// Difficulty: Medium
// Complexity: O(n*k) time, O(n*k) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", SumOfKSubarraysWithLengthAtLeastM([]int{1, 2, 3, 4}, 2, 2))
	// Test case 2
	fmt.Println("Test 2:", SumOfKSubarraysWithLengthAtLeastM([]int{5, 1, 2, 3}, 2, 1))
	// Test case 3
	fmt.Println("Test 3:", SumOfKSubarraysWithLengthAtLeastM([]int{-1, -2, -3}, 1, 2))
}

func SumOfKSubarraysWithLengthAtLeastM(nums []int, k int, m int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// dp[i][j] = max sum using exactly j subarrays considering first i elements
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1 << 60
		}
	}
	dp[0][0] = 0

	for j := 1; j <= k; j++ {
		best := -1 << 60
		for i := m * j; i <= n; i++ {
			// consider ending subarray at i-1
			start := i - m
			if dp[start][j-1] > best {
				best = dp[start][j-1]
			}
			for l := m; l <= i; l++ {
				if dp[i-l][j-1] != -1<<60 {
					sum := prefix[i] - prefix[i-l]
					if dp[i-l][j-1]+sum > dp[i][j] {
						dp[i][j] = dp[i-l][j-1] + sum
					}
				}
			}
		}
	}

	if dp[n][k] == -1<<60 {
		return 0
	}
	return dp[n][k]
}
