package main

// LeetCode #727: Minimum Window Subsequence
// https://leetcode.com/problems/minimum-window-subsequence/
// Difficulty: Hard [Paid]
//
// Algorithm: DP with next occurrence array
// 1. Precompute next[i][c] = next occurrence of char c at or after position i in S
// 2. For each starting position in S, try to match T greedily
// 3. Find the shortest valid window
//
// Alternative: For each char in T, build an array of positions.
// Then starting from each position in S, extend to match all of T.

import (
	"fmt"
	"math"
)

func minWindow(S string, T string) string {
	m, n := len(S), len(T)
	if m == 0 || n == 0 {
		return ""
	}

	// dp[i][j] = starting position of the minimum window in S[i:] that contains T[j:]
	// We'll use a simpler approach: for each start, try to match T greedily.

	// Precompute next occurrence of each character from each position
	// nextPos[i][char] = smallest index >= i such that S[index] == char
	const totalChars = 26
	nextPos := make([][totalChars]int, m+1)
	for c := 0; c < totalChars; c++ {
		nextPos[m][c] = math.MaxInt32
	}
	for i := m - 1; i >= 0; i-- {
		for c := 0; c < totalChars; c++ {
			nextPos[i][c] = nextPos[i+1][c]
		}
		nextPos[i][S[i]-'a'] = i
	}

	bestStart, bestLen := -1, math.MaxInt32

	// Try each starting position
	for start := 0; start < m; start++ {
		// Skip if S[start] doesn't match T[0]
		if S[start] != T[0] {
			continue
		}
		// Greedy match
		pos := start
		j := 0
		for j < n && pos < m {
			c := T[j] - 'a'
			if nextPos[pos][c] == math.MaxInt32 {
				break
			}
			pos = nextPos[pos][c] + 1
			j++
		}
		if j == n {
			// Found a valid window S[start:pos]
			length := pos - start
			if length < bestLen {
				bestLen = length
				bestStart = start
			}
		}
	}

	if bestStart == -1 {
		return ""
	}
	return S[bestStart : bestStart+bestLen]
}

func main() {
	// Example from problem
	S1 := "abcdebdde"
	T1 := "bde"
	result1 := minWindow(S1, T1)
	fmt.Printf("Input: S=%q, T=%q\nOutput: %q (expected: bcde)\n\n", S1, T1, result1)

	// Test case 2
	S2 := "abcde"
	T2 := "ace"
	result2 := minWindow(S2, T2)
	fmt.Printf("Input: S=%q, T=%q\nOutput: %q (expected: abcde)\n\n", S2, T2, result2)

	// Test case 3: no match
	S3 := "aaaa"
	T3 := "bb"
	result3 := minWindow(S3, T3)
	fmt.Printf("Input: S=%q, T=%q\nOutput: %q (expected: \"\")\n\n", S3, T3, result3)

	// Test case 4: S equals T
	S4 := "xyz"
	T4 := "xyz"
	result4 := minWindow(S4, T4)
	fmt.Printf("Input: S=%q, T=%q\nOutput: %q (expected: xyz)\n\n", S4, T4, result4)

	// Test case 5: single char
	S5 := "ababa"
	T5 := "a"
	result5 := minWindow(S5, T5)
	fmt.Printf("Input: S=%q, T=%q\nOutput: %q (expected: a)\n\n", S5, T5, result5)

	// Test case 6: multiple windows, pick shortest
	S6 := "fffnntt"
	T6 := "nt"
	result6 := minWindow(S6, T6)
	fmt.Printf("Input: S=%q, T=%q\nOutput: %q (expected: nt)\n\n", S6, T6, result6)

	// Test case 7: need to skip earlier match for shorter window
	S7 := "abacaba"
	T7 := "aa"
	result7 := minWindow(S7, T7)
	fmt.Printf("Input: S=%q, T=%q\nOutput: %q\n", S7, T7, result7)
}
