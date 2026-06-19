package main

// LeetCode #761: Special Binary String
// https://leetcode.com/problems/special-binary-string/
// Difficulty: Hard
//
// A special binary string is one that:
//   - Has equal number of 0s and 1s
//   - Every prefix has at least as many 1s as 0s
//
// Operation: take two consecutive special substrings and swap them.
// Goal: return the lexicographically largest string achievable.
//
// Approach: Recursive
// 1. For a given special binary string, split it into top-level
//    special substrings (balanced substrings that are not nested
//    inside another balanced substring).
// 2. Recursively process each substring.
// 3. Sort the processed substrings in reverse order (descending).
// 4. Concatenate and return.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(makeLargestSpecial("11011000")) // "11100100"
	fmt.Println(makeLargestSpecial("10"))        // "10"
	fmt.Println(makeLargestSpecial("1010"))      // "1010" (already maximal)
}

func makeLargestSpecial(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// Split into top-level special substrings
	var subs []string
	count := 0
	start := 0
	for i, ch := range s {
		if ch == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			// s[start:i+1] is a special binary string
			// Recursively process the inner part (between 1 and 0)
			inner := makeLargestSpecial(s[start+1 : i])
			subs = append(subs, "1"+inner+"0")
			start = i + 1
		}
	}

	// Sort in descending order (lexicographically largest first)
	sort.Slice(subs, func(i, j int) bool {
		return subs[i] > subs[j]
	})

	// Concatenate
	result := ""
	for _, sub := range subs {
		result += sub
	}
	return result
}
