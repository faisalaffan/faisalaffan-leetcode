package main

// LeetCode #1246: Palindrome Removal
// https://leetcode.com/problems/palindrome-removal/
// Difficulty: Hard [Paid]
//
// Given an integer array arr, in one move you can select a palindromic
// contiguous subarray and remove it. Find the minimum number of moves to
// remove all elements from the array.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumMoves([]int{1, 2})) // 2

	// Example 2
	fmt.Println(minimumMoves([]int{1, 3, 4, 1, 5})) // 3
	// Explanation: Remove [3,4,5] individually (3 moves), then [1,1] in 1 move. Total 3.
	// Actually: remove 4, remove 3, remaining [1, 5, 1] palindrome -> 1 more = 3.
	// Or: remove 4 (1), remove 5 (1), [1,3,1] palindrome (1). Total 3.

	// All same
	fmt.Println(minimumMoves([]int{1, 1, 1})) // 1

	// Single
	fmt.Println(minimumMoves([]int{5})) // 1

	// Two different
	fmt.Println(minimumMoves([]int{1, 3})) // 2

	// Two same
	fmt.Println(minimumMoves([]int{2, 2})) // 1

	// Example from problem
	fmt.Println(minimumMoves([]int{1, 2, 3, 1})) // 2

	// Larger non-trivial
	fmt.Println(minimumMoves([]int{1, 2, 3, 4, 5, 1})) // ?
}

// minimumMoves returns the minimum number of moves to remove the whole array.
//
// DP recurrence:
// dp[i][j] = minimum moves to remove arr[i..j] (inclusive)
//
// Base: dp[i][i] = 1 (single element is always a palindrome)
// dp[i][i+1] = 1 if arr[i] == arr[i+1] else 2
//
// For longer subarrays:
//   1. Split: dp[i][j] = min(dp[i][k] + dp[k+1][j]) for k in [i, j-1]
//   2. If arr[i] == arr[j]: dp[i][j] = min(dp[i][j], max(1, dp[i+1][j-1]))
//      When the ends match, they can be removed together with the inner
//      subarray — the inner removal takes dp[i+1][j-1] moves, and the
//      matching ends are consumed in the last move (or form their own
//      palindrome if the inner is empty).
func minimumMoves(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	// Base for length 2
	for i := 0; i < n-1; i++ {
		if arr[i] == arr[i+1] {
			dp[i][i+1] = 1
		} else {
			dp[i][i+1] = 2
		}
	}

	// Process by increasing length
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1

			// Start with a large number
			dp[i][j] = n

			// Split into two subproblems
			for k := i; k < j; k++ {
				candidate := dp[i][k] + dp[k+1][j]
				if candidate < dp[i][j] {
					dp[i][j] = candidate
				}
			}

			// If ends match, they can be removed together
			if arr[i] == arr[j] {
				inner := 0
				if i+1 < j {
					inner = dp[i+1][j-1]
				}
				candidate := max(1, inner)
				if candidate < dp[i][j] {
					dp[i][j] = candidate
				}
			}
		}
	}

	return dp[0][n-1]
}
