# 2663 — Lexicographically Smallest Beautiful String

## Deskripsi

**Soal:** [2663. Lexicographically Smallest Beautiful String](https://leetcode.com/problems/lexicographically-smallest-beautiful-string/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2663: Lexicographically Smallest Beautiful String
// https://leetcode.com/problems/lexicographically-smallest-beautiful-string/
// Difficulty: Hard
//
// A beautiful string has no palindromic substrings of length >= 2.
// Given a string s of first k lowercase letters, find the lexicographically
// smallest beautiful string strictly greater than s. Return "" if impossible.
// Equivalent to: no two adjacent same chars (for length-2 palindromes) and
// no s[i] == s[i+2] (for length-3 palindromes). Higher-length palindromes
// are automatically avoided.

import "fmt"

func main() {
	// Example 1: s="ab", k=3 -> "ac" (palindromes of len 2: "aa" invalid, len 3: "aba" invalid)
	fmt.Println(smallestBeautifulString("ab", 3))
	// Example 2: s="abcz", k=26
	fmt.Println(smallestBeautifulString("abcz", 26))
	// Example 3: s="dc", k=4 -> ""
	fmt.Println(smallestBeautifulString("dc", 4))
}

func smallestBeautifulString(s string, k int) string {
	n := len(s)
	b := []byte(s)
	maxChar := byte('a' + k - 1)

	// Iterate from right to left, try to increment each position
	for i := n - 1; i >= 0; i-- {
		// Try each possible larger character at position i
		for c := b[i] + 1; c <= maxChar; c++ {
			b[i] = c
			// Check local palindrome constraints for position i
			if i > 0 && b[i] == b[i-1] {
				continue
			}
			if i > 1 && b[i] == b[i-2] {
				continue
			}

			// Fill suffix with smallest possible characters
			if fillSuffix(b, i+1, maxChar) {
				return string(b)
			}
		}
	}
	return ""
}

// fillSuffix fills b[pos:] with the smallest possible characters
// that avoid length-2 and length-3 palindromes.
func fillSuffix(b []byte, pos int, maxChar byte) bool {
	for j := pos; j < len(b); j++ {
		found := false
		for c := byte('a'); c <= maxChar; c++ {
			if j > 0 && c == b[j-1] {
				continue
			}
			if j > 1 && c == b[j-2] {
				continue
			}
			b[j] = c
			found = true
			break
		}
		if !found {
			return false
		}
	}
	return true
}
```
