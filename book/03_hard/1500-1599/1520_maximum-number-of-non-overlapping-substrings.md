# 1520 — Maximum Number Of Non Overlapping Substrings

## Deskripsi

**Soal:** [1520. Maximum Number Of Non Overlapping Substrings](https://leetcode.com/problems/maximum-number-of-non-overlapping-substrings/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal)

> **Ide Kunci:** Greedy Interval

## Solusi Go

```go
package main

// LeetCode #1520: Maximum Number of Non-Overlapping Substrings
// https://leetcode.com/problems/maximum-number-of-non-overlapping-substrings/
// Difficulty: Hard
//
// Approach: Greedy Interval
// 1. For each character, find its first and last occurrence in s.
// 2. For each character, expand its interval until all chars in the interval
//    have their full range within the interval.
// 3. Sort intervals by end ascending. Greedy pick: if start > last_end, take it.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxNumOfSubstrings("adefaddaccc"))
	// Expected: ["e","f","ccc"] or similar valid order

	// Example 2
	fmt.Println(maxNumOfSubstrings("abbaccd"))
	// Expected: ["d","bb","cc"] or similar
}

type iv struct{ l, r int }

func maxNumOfSubstrings(s string) []string {
	n := len(s)

	// First and last occurrence of each char
  // Membuat slice untuk menyimpan hasil
	first := make([]int, 26)
  // Membuat slice untuk menyimpan hasil
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = n
		last[i] = -1
	}
	for i, ch := range s {
		c := int(ch - 'a')
		if i < first[c] {
			first[c] = i
		}
		if i > last[c] {
			last[c] = i
		}
	}

	// Compute minimal interval for each character that appears
  // Membuat slice untuk menyimpan hasil
	minIntervals := make([]iv, 0)
	for c := 0; c < 26; c++ {
		if first[c] == n {
			continue
		}
		l, r := first[c], last[c]
		changed := true
		for changed {
			changed = false
			for j := l; j <= r; j++ {
				ch := int(s[j] - 'a')
				if first[ch] < l {
					l = first[ch]
					changed = true
				}
				if last[ch] > r {
					r = last[ch]
					changed = true
				}
			}
		}
		minIntervals = append(minIntervals, iv{l, r})
	}

	// Dedup by (l,r)
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]map[int]bool)
  // Membuat slice untuk menyimpan hasil
	unique := make([]iv, 0)
	for _, inv := range minIntervals {
		if seen[inv.l] == nil {
			seen[inv.l] = make(map[int]bool)
		}
		if !seen[inv.l][inv.r] {
			seen[inv.l][inv.r] = true
			unique = append(unique, inv)
		}
	}

	// Sort by end ascending
	sort.Slice(unique, func(i, j int) bool {
		return unique[i].r < unique[j].r
	})

	// Greedy pick non-overlapping substrings
  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)
	end := -1
	for _, inv := range unique {
		if inv.l > end {
			result = append(result, s[inv.l:inv.r+1])
			end = inv.r
		}
	}

	return result
}
```
