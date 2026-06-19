package main

// LeetCode #3104: Find Longest Self-Contained Substring
// https://leetcode.com/problems/find-longest-self-contained-substring/
// Difficulty: Hard [Paid]
//
// A substring s[i..j] is self-contained if for every character c in the substring,
// ALL occurrences of c in the original string are within [i, j].
// Find the longest self-contained substring length.

import (
	"fmt"
)

func longestSelfContainedSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// First and last occurrence of each character
	first := make([]int, 26)
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = n
		last[i] = -1
	}
	for i := 0; i < n; i++ {
		c := int(s[i] - 'a')
		if first[c] > i {
			first[c] = i
		}
		if last[c] < i {
			last[c] = i
		}
	}

	maxLen := 0
	// For each start position, expand the window until it is self-contained
	for i := 0; i < n; i++ {
		end := i
		for j := i; j <= end && j < n; j++ {
			c := int(s[j] - 'a')
			if last[c] > end {
				end = last[c]
			}
		}
		if end-i+1 > maxLen {
			maxLen = end - i + 1
		}
		// Optimization: if the character at i starts at i (first occurrence),
		// we might skip ahead using its last occurrence
		c := int(s[i] - 'a')
		if first[c] == i {
			i = last[c] // loop increment will advance past it
		}
	}

	return maxLen
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", longestSelfContainedSubstring("abacd"))
	// Expected: 4 ("abac" or "baca"? Actually "abac" has a(0,2), b(1), c(3) all within [0,3])

	// Test case 2: single character
	fmt.Println("Test 2:", longestSelfContainedSubstring("a"))
	// Expected: 1

	// Test case 3: all distinct
	fmt.Println("Test 3:", longestSelfContainedSubstring("abcdef"))
	// Expected: 6 (any substring is self-contained)

	// Test case 4: repeating
	fmt.Println("Test 4:", longestSelfContainedSubstring("ababa"))
	// Expected: 5 (all a's and b's must be included)

	// Test case 5: empty
	fmt.Println("Test 5:", longestSelfContainedSubstring(""))
	// Expected: 0
}
