package main

// LeetCode #3253: Construct String with Minimum Cost (Easy)
// https://leetcode.com/problems/construct-string-with-minimum-cost-easy/
// Difficulty: Medium [Paid]
// Time: O(n * m * L) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumCost(target string, words []string, costs []int) int {
	n := len(target)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		if dp[i] == math.MaxInt32 {
			continue
		}
		for j, w := range words {
			if i+len(w) <= n && target[i:i+len(w)] == w {
				if dp[i]+costs[j] < dp[i+len(w)] {
					dp[i+len(w)] = dp[i] + costs[j]
				}
			}
		}
	}

	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func main() {
	fmt.Println(minimumCost("abc", []string{"a", "bc", "abc"}, []int{1, 2, 3})) // Expected: 3
	fmt.Println(minimumCost("xyz", []string{"ab", "cd"}, []int{1, 2}))          // Expected: -1
}
