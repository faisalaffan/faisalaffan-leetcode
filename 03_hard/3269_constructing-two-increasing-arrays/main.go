package main

// LeetCode #3269: Constructing Two Increasing Arrays
// https://leetcode.com/problems/constructing-two-increasing-arrays/
// Difficulty: Hard [Paid]
//
// Given two strings s1 and s2 representing integers, and lengths len1, len2,
// count the number of ways to construct two strictly increasing arrays
// of lengths len1 and len2 from the digits of s1 and s2 respectively,
// preserving order within each string.
//
// This is a DP problem: for each position in s1 and s2, we decide whether
// to take the current digit (if it continues the increasing trend) or skip it.

import (
	"fmt"
)

const mod = 1_000_000_007

func main() {
	// Example 1
	fmt.Println(constructingTwoIncreasingArrays("123", "456", 2, 2))
	// Example 2
	fmt.Println(constructingTwoIncreasingArrays("12", "34", 1, 1))
	// Example 3
	fmt.Println(constructingTwoIncreasingArrays("1234", "5678", 3, 3))
	// Example 4: single digit each
	fmt.Println(constructingTwoIncreasingArrays("1", "2", 1, 1))
	// Example 5
	fmt.Println(constructingTwoIncreasingArrays("111", "222", 2, 2))
}

func constructingTwoIncreasingArrays(s1 string, s2 string, len1 int, len2 int) int {
	n1, n2 := len(s1), len(s2)

	// dp[i][j][a][b] = number of ways using first i chars of s1 and first j chars of s2,
	// with last chosen value of a (from s1) and b (from s2).
	// This is too large. Instead use:
	// dp[i][j][k][last]: i chars from s1 considered, j from s2 considered,
	// k elements chosen for array1, last element value (0-9 for digits, 10 = none).

	// Since digits are only 0-9 and we must pick strictly increasing subsequence,
	// we can use DP over (pos1, pos2, taken1, taken2, last1, last2).
	// This is O(n1*n2*len1*len2*10*10) which might be large.
	//
	// Optimized approach: use DP where state is (i, j, k) = ways using first i s1 chars,
	// first j s2 chars, k elements picked for array1 (and don't track the actual last values).
	// Instead track that the last two picked values form an increasing pair.

	// dp[i][j][a][b] = ways using prefix i of s1, prefix j of s2,
	// with last value a from s1 (or 10 for none) and last value b from s2 (or 10 for none).
	dp := make([][][][]int, n1+1)
	for i := range dp {
		dp[i] = make([][][]int, n2+1)
		for j := range dp[i] {
			dp[i][j] = make([][]int, 11)
			for a := range dp[i][j] {
				dp[i][j][a] = make([]int, 11)
			}
		}
	}

	// Initialize: empty state has 1 way, with no last values (10 = none).
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			dp[i][j][10][10] = 1
		}
	}

	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			for a := 0; a <= 10; a++ {
				for b := 0; b <= 10; b++ {
					cur := dp[i][j][a][b]
					if cur == 0 {
						continue
					}
					// Decide to build the next pair by picking one digit from s1 and one from s2.
					// Try picking s1[i] as the next element of array 1 (if it increases).
					if i+1 <= n1 {
						// Skip s1[i]
						dp[i+1][j][a][b] = (dp[i+1][j][a][b] + cur) % mod
					}
					if j+1 <= n2 {
						// Skip s2[j]
						dp[i][j+1][a][b] = (dp[i][j+1][a][b] + cur) % mod
					}
					// Pick s1[i] for array 1 and s2[j] for array 2 simultaneously.
					if i+1 <= n1 && j+1 <= n2 {
						d1 := int(s1[i] - '0')
						d2 := int(s2[j] - '0')
						if (a == 10 || d1 > a) && (b == 10 || d2 > b) {
							dp[i+1][j+1][d1][d2] = (dp[i+1][j+1][d1][d2] + cur) % mod
						}
						// Also skip just this pair but advance both pointers.
						dp[i+1][j+1][a][b] = (dp[i+1][j+1][a][b] + cur) % mod
					}
				}
			}
		}
	}

	ans := 0
	for a := 0; a <= 10; a++ {
		for b := 0; b <= 10; b++ {
			ans = (ans + dp[n1][n2][a][b]) % mod
		}
	}
	return ans
}
