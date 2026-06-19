package main

// LeetCode #2478: Number of Beautiful Partitions
// https://leetcode.com/problems/number-of-beautiful-partitions/
// Difficulty: Hard
//
// DP with prefix sum. Partition s into k substrings, each:
//   - length >= minLength
//   - first char is prime digit (2,3,5,7)
//   - last char is non-prime digit
//
// DP[i][j] = ways to partition s[0:i] into j valid substrings.
// Optimized with prefix sum to O(n*k).

import "fmt"

const MOD = 1000000007

func main() {
	// Example 1: "23542185131", 3, 2 => 3
	fmt.Println(beautifulPartitions("23542185131", 3, 2))
	// Example 2: "23542185131", 3, 5 => 0
	fmt.Println(beautifulPartitions("23542185131", 3, 5))
	// Edge: single partition whole string
	fmt.Println(beautifulPartitions("331", 1, 3))
	// Edge: no valid partition
	fmt.Println(beautifulPartitions("22", 1, 2))
	// Edge: k=0
	fmt.Println(beautifulPartitions("235", 0, 2))
}

func isPrimeDigit(c byte) bool {
	return c == '2' || c == '3' || c == '5' || c == '7'
}

func beautifulPartitions(s string, k int, minLength int) int {
	n := len(s)

	// Quick check: first char must be prime, last must be non-prime
	if !isPrimeDigit(s[0]) || isPrimeDigit(s[n-1]) {
		return 0
	}
	if n < minLength*k {
		return 0
	}
	if k == 0 {
		return 1
	}

	// validSplit[i] = true if we can split after position i
	// meaning s[i-1] is non-prime and s[i] is prime
	validSplit := make([]bool, n+1)
	validSplit[0] = true
	for p := 1; p < n; p++ {
		if !isPrimeDigit(s[p-1]) && isPrimeDigit(s[p]) {
			validSplit[p] = true
		}
	}
	validSplit[n] = !isPrimeDigit(s[n-1])

	// dp[i][j] = ways using first i chars, j partitions
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, k+1)
	}
	dp[0][0] = 1

	for j := 1; j <= k; j++ {
		sum := int64(0)
		for i := 1; i <= n; i++ {
			// Add dp[p][j-1] where p = i - minLength is a valid split
			if i >= minLength {
				p := i - minLength
				if validSplit[p] {
					sum = (sum + dp[p][j-1]) % MOD
				}
			}
			// Current position i is a valid end if s[i-1] is non-prime
			if !isPrimeDigit(s[i-1]) {
				dp[i][j] = sum
			}
		}
	}

	return int(dp[n][k])
}
