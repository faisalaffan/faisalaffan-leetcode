package main

// LeetCode #3144: Minimum Substring Partition of Equal Character Frequency
// https://leetcode.com/problems/minimum-substring-partition-of-equal-character-frequency/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumSubstringsInPartition(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		freq := make([]int, 26)
		var distinct, maxFreq int
		for j := i - 1; j >= 0; j-- {
			idx := s[j] - 'a'
			if freq[idx] == 0 {
				distinct++
			}
			freq[idx]++
			if freq[idx] > maxFreq {
				maxFreq = freq[idx]
			}

			if maxFreq*distinct == i-j {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(minimumSubstringsInPartition("fabccddg")) // Expected: 3
	fmt.Println(minimumSubstringsInPartition("abababaccddb")) // Expected: ?
}
