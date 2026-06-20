package main

// LeetCode #3579: Minimum Steps to Convert String with Operations
// https://leetcode.com/problems/minimum-steps-to-convert-string-with-operations/
// Difficulty: Hard
//
// Given two strings word1 and word2 of equal length, find the minimum number of
// operations to transform word1 into word2. Allowed operations per substring:
// replace (change one char), swap (any two chars), reverse (entire substring).
// Each character can be involved in each operation type at most once.
//
// Approach: DP over intervals. For substring i..j, compute min operations
// to convert word1[i..j] to word2[i..j].

import "fmt"
import "math"

func main() {
	// Example 1
	fmt.Println(minOperations("abc", "cba"))
	// Example 2
	fmt.Println(minOperations("abcd", "bcda"))
	// Example 3
	fmt.Println(minOperations("ab", "ab"))
	// Edge: single char
	fmt.Println(minOperations("a", "b"))
	// Edge: already equal
	fmt.Println(minOperations("abc", "abc"))
}

func minOperations(word1 string, word2 string) int {
	n := len(word1)
	if n == 0 {
		return 0
	}

	// dp[i][j] = min ops to convert word1[i..j] to word2[i..j]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Single character: need replace if diff
	for i := 0; i < n; i++ {
		if word1[i] == word2[i] {
			dp[i][i] = 0
		} else {
			dp[i][i] = 1
		}
	}

	// Interval DP
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			best := math.MaxInt32

			// Option 1: split into two parts
			for k := i; k < j; k++ {
				cost := dp[i][k] + dp[k+1][j]
				if cost < best {
					best = cost
				}
			}

			// Option 2: swap ends
			if word1[i] == word2[j] && word1[j] == word2[i] {
				// Swap first and last chars
				if i+1 <= j-1 {
					if dp[i+1][j-1] < best {
						best = dp[i+1][j-1]
					}
				} else {
					if 0 < best {
						best = 0
					}
				}
			}

			// Option 3: reverse the whole substring
			// Check if word2 is the reverse of word1 for this range
			rev := true
			for offset := 0; offset < length; offset++ {
				if word1[i+offset] != word2[j-offset] {
					rev = false
					break
				}
			}
			if rev {
				// One reverse operation
				if 1 < best {
					best = 1
				}
			}

			dp[i][j] = best
		}
	}

	return dp[0][n-1]
}
