# 3455 — Shortest Matching Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func shortestMatchingSubstring(s string, p string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3455: Shortest Matching Substring
// https://leetcode.com/problems/shortest-matching-substring/
// Difficulty: Hard
//
// Given string s and pattern p containing exactly two '*' wildcards,
// find the length of the shortest substring in s that matches p.
// '*' matches any sequence of characters (including empty).
//
// Approach: Split pattern by '*'. Find positions of each part in s.
// For each match position of part1, find the earliest match of part2
// after it, then part3 after part2.

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(shortestMatchingSubstring("ababa", "a*b*a"))
	// Example 2
	fmt.Println(shortestMatchingSubstring("abcdef", "a*d*f"))
	// Example 3: no match
	fmt.Println(shortestMatchingSubstring("abc", "a*d"))
	// Edge: * matches empty
	fmt.Println(shortestMatchingSubstring("abc", "a**c"))
	// Edge: exact match without *
	fmt.Println(shortestMatchingSubstring("abc", "a**c"))
}

func shortestMatchingSubstring(s string, p string) int {
	// Split pattern by '*' to get parts
	parts := strings.Split(p, "*")
	// Filter empty parts (consecutive * or leading/trailing *)
	var filtered []string
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}

	if len(filtered) == 0 {
		return 0 // pattern is all wildcards
	}

	// For each part, find all start positions
  // Matriks 2D
	partPositions := make([][]int, len(filtered))
	for i, part := range filtered {
		positions := findAllOccurrences(s, part)
		if len(positions) == 0 {
			return -1
		}
		partPositions[i] = positions
	}

	// Find shortest window that covers all parts in order
	ans := math.MaxInt32

	if len(filtered) == 1 {
		for _, pos := range partPositions[0] {
			end := pos + len(filtered[0])
			if end < ans {
				ans = end
			}
		}
		if ans == math.MaxInt32 {
			return -1
		}
		return ans
	}

	if len(filtered) == 2 {
		for _, p1 := range partPositions[0] {
			end1 := p1 + len(filtered[0])
			for _, p2 := range partPositions[1] {
				if p2 >= end1 {
					length := p2 + len(filtered[1]) - p1
					if length < ans {
						ans = length
					}
					break // first match after part1 is shortest
				}
			}
		}
		if ans == math.MaxInt32 {
			return -1
		}
		return ans
	}

	if len(filtered) == 3 {
		for _, p1 := range partPositions[0] {
			end1 := p1 + len(filtered[0])
			for _, p2 := range partPositions[1] {
				if p2 < end1 {
					continue
				}
				end2 := p2 + len(filtered[1])
				for _, p3 := range partPositions[2] {
					if p3 >= end2 {
						length := p3 + len(filtered[2]) - p1
						if length < ans {
							ans = length
						}
						break
					}
				}
			}
		}
		if ans == math.MaxInt32 {
			return -1
		}
		return ans
	}

	return -1
}

func findAllOccurrences(s, sub string) []int {
	var res []int
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			res = append(res, i)
		}
	}
	return res
}
```
