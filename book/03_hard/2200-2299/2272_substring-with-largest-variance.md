# 2272 — Substring With Largest Variance

## Deskripsi

**Soal:** [2272. Substring With Largest Variance](https://leetcode.com/problems/substring-with-largest-variance/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func largestVariance(s string) int`

> **Ide Kunci:** For each pair of characters (major, minor),

## Solusi Go

```go
package main

// LeetCode #2272: Substring With Largest Variance
// https://leetcode.com/problems/substring-with-largest-variance/
// Difficulty: Hard
//
// The variance of a string is defined as the largest difference
// between the counts of any two distinct characters in the string.
// Given a string s consisting of lowercase English letters,
// return the largest variance among all substrings of s.

import (
	"fmt"
	"math"
)

// largestVariance returns the maximum variance across all substrings.
//
// Approach: For each pair of characters (major, minor),
// treat major as +1, minor as -1, and find maximum subarray sum,
// ensuring at least one minor appears (variance requires both chars).
func largestVariance(s string) int {
	// count frequency of each char
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	maxVar := 0

	// For each pair (a, b) where a != b
	for a := 0; a < 26; a++ {
		if freq[a] == 0 {
			continue
		}
		for b := 0; b < 26; b++ {
			if a == b || freq[b] == 0 {
				continue
			}

			// Kadane-like: treat a as +1, b as -1
			// We need substrings containing at least one b.
			// Use two DP values:
			//   cur: maximum subarray ending here (allowing no b)
			//   curWithB: maximum subarray ending here with at least one b
			cur := 0
			curWithB := math.MinInt32

  // Loop standar: indeks 0 sampai n-1
			for i := 0; i < len(s); i++ {
				ch := int(s[i] - 'a')

				if ch == a {
					cur++
					if curWithB != math.MinInt32 {
						curWithB++
					}
				} else if ch == b {
					cur--
					curWithB = cur // at least we have b now
					if cur < 0 {
						cur = 0 // reset (start fresh from here)
					}
				} else {
					// other characters: they contribute 0 difference
					// but affect whether we have b
					// They don't change the relative count
				}

				if curWithB > maxVar {
					maxVar = curWithB
				}
			}
		}
	}

	return maxVar
}

func main() {
	// Example 1
	fmt.Println(largestVariance("aababbb")) // Expected: 3

	// Example 2
	fmt.Println(largestVariance("abcde")) // Expected: 0
}
```
