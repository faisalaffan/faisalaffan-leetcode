package main

// LeetCode #2478: Number of Beautiful Partitions
// https://leetcode.com/problems/number-of-beautiful-partitions/
// Difficulty: Hard
//
// DP[i][j] = number of ways to partition s[0:i] into j substrings where each
// substring has length >= minLength, starts with a prime digit, and ends with
// a non-prime digit. s[i-1] must be the end of the j-th substring.
// Transition: DP[i][j] = sum(DP[p][j-1]) where s[p] is prime (start of current)
// and i-p >= minLength.

import "fmt"

const MOD = 1000000007

func main() {
	// Example 1: "23542185131", 3, 2 => 3
	fmt.Println(beautifulPartitions("23542185131", 3, 2))
	// Example 2: "23542185131", 3, 5 => 0
	fmt.Println(beautifulPartitions("23542185131", 3, 5))
	// Edge: single partition whole string
	fmt.Println(beautifulPartitions("331", 3, 1))
	// Edge: no valid partition
	fmt.Println(beautifulPartitions("22", 2, 1))
}

func isPrimeDigit(c byte) bool {
	return c == '2' || c == '3' || c == '5' || c == '7'
}

func beautifulPartitions(s string, minLength int, k int) int {
	n := len(s)

	// Quick pruning: first char must be prime, last must be non-prime
	if !isPrimeDigit(s[0]) || isPrimeDigit(s[n-1]) {
		return 0
	}
	if n < minLength*k {
		return 0
	}

	// validEnd[i] = true if we can end a partition at position i (0-indexed, exclusive)
	// A partition s[l..r] ends at r+1. s[r] must be non-prime and s[l] must be prime.
	// For a split at position p (end of one substring), s[p-1] must be non-prime
	// and s[p] must be prime.
	validSplit := make([]bool, n+1) // split position (exclusive end)
	for p := 1; p < n; p++ {
		if !isPrimeDigit(s[p-1]) && isPrimeDigit(s[p]) {
			validSplit[p] = true
		}
	}
	validSplit[0] = true                                                    // start of string
	validSplit[n] = validSplit[n] || !isPrimeDigit(s[n-1])                  // end of string

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	// Base: one way to have 0 partitions of empty prefix
	dp[0][0] = 1

	// dp[i][j] = sum over p where i-p >= minLength, validSplit[p] is true
	// and p < i, of dp[p][j-1], and s[p] is prime (start of new substring)
	// Optimization using prefix sums
	for j := 1; j <= k; j++ {
		sum := 0
		for i := 1; i <= n; i++ {
			// Check if s[i-1] ends a partition (must be non-prime)
			if i >= minLength {
				p := i - minLength
				if validSplit[p] {
					sum = (sum + dp[p][j-1]) % MOD
				}
			}
			// s[i-1] must be non-prime for this to be a valid end
			if !isPrimeDigit(s[i-1]) {
				dp[i][j] = sum
			}
		}
	}

	return dp[n][k]
}
