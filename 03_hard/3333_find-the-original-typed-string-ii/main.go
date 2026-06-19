package main

// LeetCode #3333: Find the Original Typed String II
// https://leetcode.com/problems/find-the-original-typed-string-ii/
// Difficulty: Hard
//
// Given a typed string (result of pressing keys where a key may be pressed
// multiple times), count the number of possible original strings that could
// result in the given typed string when the typist may have pressed each key
// one or more times. The original string must have length exactly k.
//
// Approach: DP with combinatorics. Group consecutive identical characters.
// For each group length len, we can choose to keep 1..len characters.
// Use sliding window DP over groups to count ways to achieve total length k.

import "fmt"

func main() {
	// Example 1
	fmt.Println(possibleStringCount("aabb", 3))
	// Example 2
	fmt.Println(possibleStringCount("aabb", 2))
	// Example 3
	fmt.Println(possibleStringCount("aaaa", 2))
	// Edge: k equals original length
	fmt.Println(possibleStringCount("abc", 3))
	// Edge: k = 1
	fmt.Println(possibleStringCount("aa", 1))
}

const MOD = 1000000007

func possibleStringCount(word string, k int) int {
	// Group consecutive characters
	var groups []int
	count := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			count++
		} else {
			groups = append(groups, count)
			count = 1
		}
	}
	groups = append(groups, count)

	// If k is larger than total possible characters, return 0
	totalPossible := len(word)
	if k > totalPossible {
		return 0
	}

	// Each group must contribute at least 1 character
	n := len(groups)
	if k < n {
		return 0
	}

	// If k >= totalPossible (all chars), there's exactly 1 way
	if k >= totalPossible {
		return 1
	}

	// DP: dp[j] = number of ways to reach length j using processed groups
	extra := k - n // additional characters beyond the minimum 1 per group
	m := totalPossible - n

	// dp[j] = ways to use j extra characters so far
	dp := make([]int, extra+1)
	dp[0] = 1

	for _, g := range groups {
		maxAdd := g - 1 // max extra from this group
		ndp := make([]int, extra+1)
		// Prefix sum for sliding window
		prefix := make([]int, extra+2)
		for j := 0; j <= extra; j++ {
			prefix[j+1] = (prefix[j] + dp[j]) % MOD
		}
		for j := 0; j <= extra; j++ {
			// We use t extra characters from this group (0 <= t <= min(maxAdd, j))
			// ndp[j] = sum_{t=0}^{min(maxAdd, j)} dp[j-t]
			l := j - min(maxAdd, j)
			ndp[j] = (prefix[j+1] - prefix[l] + MOD) % MOD
		}
		dp = ndp
	}

	ans := 0
	for j := 0; j <= extra; j++ {
		ans = (ans + dp[j]) % MOD
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
