package main

// LeetCode #887: Super Egg Drop
// https://leetcode.com/problems/super-egg-drop/
// Difficulty: Hard
//
// DP[k][m] = maximum number of floors that can be tested with k eggs and m moves.
// Recurrence: dp[k][m] = dp[k-1][m-1] + 1 + dp[k][m-1]
//   - If egg breaks: can test dp[k-1][m-1] floors below
//   - If egg survives: can test dp[k][m-1] floors above
//   - +1 for the current floor
//
// Keep incrementing m until dp[k][m] >= n.

import "fmt"

func superEggDrop(k int, n int) int {
	// dp[eggs][moves] = max floors testable
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1) // enough space; moves never exceeds n
	}

	m := 0
	for dp[k][m] < n {
		m++
		for eggs := 1; eggs <= k; eggs++ {
			dp[eggs][m] = dp[eggs-1][m-1] + 1 + dp[eggs][m-1]
		}
	}
	return m
}

func main() {
	// Example 1: k=1, n=2 -> 2
	fmt.Println("Test 1:", superEggDrop(1, 2)) // 2

	// Example 2: k=2, n=6 -> 3
	fmt.Println("Test 2:", superEggDrop(2, 6)) // 3

	// Example 3: k=3, n=14 -> 4
	fmt.Println("Test 3:", superEggDrop(3, 14)) // 4

	// Edge: k=2, n=1 -> 1
	fmt.Println("Test 4:", superEggDrop(2, 1)) // 1

	// Edge: k=100, n=10000 -> just verify it runs
	fmt.Println("Test 5:", superEggDrop(100, 10000))

	// k=2, n=100
	fmt.Println("Test 6:", superEggDrop(2, 100)) // 14 (since 14*15/2=105 >= 100)
}
