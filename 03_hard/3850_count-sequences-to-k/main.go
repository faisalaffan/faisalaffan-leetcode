package main

// LeetCode #3850: Count Sequences to K
// https://leetcode.com/problems/count-sequences-to-k/
// Difficulty: Hard
//
// Count number of sequences (ordered subsequences) that sum to
// exactly k. Elements can be used at most once.
//
// Approach: 0/1 knapsack DP. dp[s] = number of ways to achieve sum
// s. For each num, update dp in reverse.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSequences([]int{1, 2, 3}, 4))
	// Example 2
	fmt.Println(countSequences([]int{1, 1, 1}, 2))
	// Edge: empty
	fmt.Println(countSequences([]int{}, 5))
	// Edge: unreachable
	fmt.Println(countSequences([]int{10, 20}, 5))
}

const mod = 1000000007

func countSequences(nums []int, k int64) int {
	if k == 0 {
		return 1
	}
	dp := make([]int, k+1)
	dp[0] = 1
	for _, num := range nums {
		for s := k; s >= int64(num); s-- {
			dp[s] = (dp[s] + dp[s-int64(num)]) % mod
		}
	}
	return dp[k]
}
