# 2565 — Subsequence With The Minimum Score

## Deskripsi

**Soal:** [2565. Subsequence With The Minimum Score](https://leetcode.com/problems/subsequence-with-the-minimum-score/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func minimumScore(s string, t string) int`

## Solusi Go

```go
package main

// LeetCode #2565: Subsequence With the Minimum Score
// https://leetcode.com/problems/subsequence-with-the-minimum-score/
// Difficulty: Hard

import "fmt"

// minimumScore returns the minimum length of a subsequence to remove
// from s such that t becomes a subsequence of s.
//
// Compute prefixPos[i] = position in s where first i chars of t are matched.
// Compute suffixPos[i] = position in s where last i chars of t are matched.
// Then for each split point, maximize the total matched chars.
//
// Complexity: O(n+m) time, O(m) space
func minimumScore(s string, t string) int {
	n, m := len(s), len(t)

	// prefixPos[i] = position in s where first i chars of t are matched (0-indexed)
  // Membuat slice untuk menyimpan hasil
	prefixPos := make([]int, m+1)
  // Iterasi seluruh elemen
	for i := range prefixPos {
		prefixPos[i] = n
	}
	prefixPos[0] = -1
	idx := 0
	for i := 0; i < n && idx < m; i++ {
		if s[i] == t[idx] {
			idx++
			prefixPos[idx] = i
		}
	}

	// suffixPos[i] = position in s where last i chars of t are matched
  // Membuat slice untuk menyimpan hasil
	suffixPos := make([]int, m+1)
  // Iterasi seluruh elemen
	for i := range suffixPos {
		suffixPos[i] = -1
	}
	suffixPos[0] = n
	idx = 0
	for i := n - 1; i >= 0 && idx < m; i-- {
		if s[i] == t[m-1-idx] {
			idx++
			suffixPos[idx] = i
		}
	}

	// t is already a subsequence of s
	if prefixPos[m] < n {
		return 0
	}

	ans := m
	// Try keeping first i chars from left and last j chars from right
	// They must not overlap: prefixPos[i] < suffixPos[j]
	for i := 0; i <= m; i++ {
		if prefixPos[i] == n {
			continue
		}
		// Binary search for max j such that suffixPos[j] > prefixPos[i]
		lo, hi := 0, m-i
		for lo < hi {
			mid := (lo + hi + 1) / 2
			if suffixPos[mid] > prefixPos[i] {
				lo = mid
			} else {
				hi = mid - 1
			}
		}
		kept := i + lo
		if m-kept < ans {
			ans = m - kept
		}
	}
	return ans
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: s=\"abacaba\", t=\"bzaa\" ->", minimumScore("abacaba", "bzaa")) // 1

	// Additional test cases
	fmt.Println("Test 2: s=\"abc\", t=\"abc\" ->", minimumScore("abc", "abc"))      // 0
	fmt.Println("Test 3: s=\"abcde\", t=\"ace\" ->", minimumScore("abcde", "ace"))  // 0
	fmt.Println("Test 4: s=\"abcde\", t=\"xyz\" ->", minimumScore("abcde", "xyz"))  // 3
	fmt.Println("Test 5: s=\"a\", t=\"b\" ->", minimumScore("a", "b"))              // 1
	fmt.Println("Test 6: s=\"\", t=\"a\" ->", minimumScore("", "a"))                // 1
}
```
