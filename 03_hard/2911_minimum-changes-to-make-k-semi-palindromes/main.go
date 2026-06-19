package main

// LeetCode #2911: Minimum Changes to Make K Semi-palindromes
// https://leetcode.com/problems/minimum-changes-to-make-k-semi-palindromes/
// Difficulty: Hard
//
// DP + precompute. A semi-palindrome of length L: there exists divisor d of L
// (1 <= d < L) such that grouping characters by residue class modulo d and
// each group forms a palindrome. Precompute cost[i][j] = min changes to make
// s[i:j] a semi-palindrome, then DP[k][n] = min changes to partition into k.
// O(N^3 * sqrt(N)) time, O(N^2) space.

import (
	"fmt"
	"math"
)

func minimumChanges(s string, k int) int {
	n := len(s)

	// Precompute cost[i][j] for substring s[i:j] (exclusive j), 0 <= i < j <= n
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n+1)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 2; j <= n; j++ { // need at least length 2
			length := j - i
			best := math.MaxInt32

			// Try all proper divisors d of length
			for d := 1; d < length; d++ {
				if length%d != 0 {
					continue
				}
				changes := 0
				groups := d
				groupSize := length / d

				// For each residue class (group)
				for r := 0; r < groups; r++ {
					// Characters in this group: s[i+r], s[i+r+d], s[i+r+2d], ...
					// Need to form a palindrome
					// Pair positions p and groupSize-1-p within group
					for p := 0; p < groupSize/2; p++ {
						leftIdx := i + r + p*d
						rightIdx := i + r + (groupSize-1-p)*d
						if s[leftIdx] != s[rightIdx] {
							changes++
						}
					}
				}

				if changes < best {
					best = changes
				}
			}

			cost[i][j] = best
		}
	}

	// DP[t][i] = min changes for first i chars into t semi-palindromes
	dp := make([][]int, k+1)
	for t := range dp {
		dp[t] = make([]int, n+1)
		for i := range dp[t] {
			dp[t][i] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for t := 1; t <= k; t++ {
		for i := 2 * t; i <= n; i++ { // each part needs at least 2 chars
			for j := 2 * (t - 1); j <= i-2; j++ { // previous split must leave >=2 chars
				if dp[t-1][j] == math.MaxInt32 || cost[j][i] == math.MaxInt32 {
					continue
				}
				val := dp[t-1][j] + cost[j][i]
				if val < dp[t][i] {
					dp[t][i] = val
				}
			}
		}
	}

	return dp[k][n]
}

func main() {
	// Example: s="abcac", k=2 => 1
	fmt.Println(minimumChanges("abcac", 2))
	// Example: s="abcdef", k=2 => 2
	fmt.Println(minimumChanges("abcdef", 2))
	// Example: s="aabbaa", k=3 => 0
	fmt.Println(minimumChanges("aabbaa", 3))
	// Single partition
	fmt.Println(minimumChanges("aba", 1))
	// k = n/2
	fmt.Println(minimumChanges("ab", 1))
	fmt.Println(minimumChanges("aabb", 2))
}
