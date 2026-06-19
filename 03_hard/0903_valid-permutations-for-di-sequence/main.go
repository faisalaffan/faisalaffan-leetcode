package main

// LeetCode #903: Valid Permutations for DI Sequence
// https://leetcode.com/problems/valid-permutations-for-di-sequence/
// Difficulty: Hard
// DP with prefix sums. dp[i][j] = number of perm of length i+1 ending with
// value j (0-indexed). Use prefix sums O(n^2).

import "fmt"

func numPermsDISequence(s string) int {
	n := len(s)
	mod := int(1e9 + 7)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// length 1: perm [0] has 1 ending at 0
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		// compute prefix sum of dp[i-1]
		prefix := make([]int, n+1)
		prefix[0] = dp[i-1][0]
		for k := 1; k <= n; k++ {
			prefix[k] = (prefix[k-1] + dp[i-1][k]) % mod
		}

		for j := 0; j <= i; j++ {
			if s[i-1] == 'D' {
				// sum of dp[i-1][k] for k >= j
				sum := prefix[i-1]
				if j > 0 {
					sum = (sum - prefix[j-1] + mod) % mod
				}
				dp[i][j] = sum
			} else { // 'I'
				// sum of dp[i-1][k] for k < j
				if j > 0 {
					dp[i][j] = prefix[j-1]
				}
			}
		}
	}

	ans := 0
	for j := 0; j <= n; j++ {
		ans = (ans + dp[n][j]) % mod
	}
	return ans
}

func main() {
	// Example: "DID" -> 5
	fmt.Println(numPermsDISequence("DID"))   // Expected: 5
	fmt.Println(numPermsDISequence("I"))     // Expected: 1
	fmt.Println(numPermsDISequence("D"))     // Expected: 1
	fmt.Println(numPermsDISequence("ID"))    // Expected: 2
	fmt.Println(numPermsDISequence("DI"))    // Expected: 2
	fmt.Println(numPermsDISequence("DDI"))   // Expected: 3
}
