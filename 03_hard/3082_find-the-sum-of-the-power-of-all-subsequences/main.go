package main

// LeetCode #3082: Find the Sum of the Power of All Subsequences
// https://leetcode.com/problems/find-the-sum-of-the-power-of-all-subsequences/
// Difficulty: Hard
// Time: O(n * sum(nums)) | Space: O(sum(nums))

import "fmt"

const MOD = 1_000_000_007

func sumOfPower(nums []int, k int) int {
	// dp[s] = number of subsequences with sum exactly s
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	if k > totalSum {
		return 0
	}

	dp := make([]int, totalSum+1)
	dp[0] = 1

	for _, v := range nums {
		for s := totalSum; s >= v; s-- {
			dp[s] = (dp[s] + dp[s-v]) % MOD
		}
	}

	// Sum of dp[s] for s >= k gives the count of subsequences with sum >= k
	result := 0
	for s := k; s <= totalSum; s++ {
		result = (result + dp[s]) % MOD
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sumOfPower([]int{1, 2, 3}, 2))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", sumOfPower([]int{2, 3, 3, 4, 6, 7}, 12))
	// Expected: 61

	// Test case 3
	fmt.Println("Test 3:", sumOfPower([]int{1, 1, 1}, 2))
	// Expected: 7
}
