package main

// LeetCode #3186: Maximum Total Damage With Spell Casting
// https://leetcode.com/problems/maximum-total-damage-with-spell-casting/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	freq := make(map[int]int)
	for _, v := range power {
		freq[v]++
	}

	vals := make([]int, 0, len(freq))
	for k := range freq {
		vals = append(vals, k)
	}
	sort.Ints(vals)

	n := len(vals)
	dp := make([]int64, n)

	for i := 0; i < n; i++ {
		val := vals[i]
		count := freq[val]
		dp[i] = int64(val) * int64(count)

		// Find prev valid (val - 2)
		for j := i - 1; j >= 0; j-- {
			if vals[j] < val-2 {
				if dp[j] > dp[i] {
					dp[i] = dp[j]
				}
				break
			}
			if vals[j] <= val-2 {
				dp[i] = maxInt64(dp[i], dp[j]+int64(val)*int64(count))
			}
		}

		if i > 0 && dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}
	}

	return dp[n-1]
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumTotalDamage([]int{1, 1, 3, 4}))          // Expected: 6
	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 3}))           // Expected: 10
}
