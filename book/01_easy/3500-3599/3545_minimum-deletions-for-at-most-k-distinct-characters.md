# 3545 — Minimum Deletions For At Most K Distinct Characters

## Deskripsi

**Soal:** [3545. Minimum Deletions For At Most K Distinct Characters](https://leetcode.com/problems/minimum-deletions-for-at-most-k-distinct-characters/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3545: Minimum Deletions for At Most K Distinct Characters
// https://leetcode.com/problems/minimum-deletions-for-at-most-k-distinct-characters/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumDeletionsForAtMostKDistinctCharacters("aabbbcc", 2))
	fmt.Println(MinimumDeletionsForAtMostKDistinctCharacters("abcde", 2))
}

// MinimumDeletionsForAtMostKDistinctCharacters returns min deletions so the string has at most k distinct characters.
// Time: O(n log n). Space: O(1).
func MinimumDeletionsForAtMostKDistinctCharacters(s string, k int) int {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	// Count distinct characters
	distinct := 0
	for _, f := range freq {
		if f > 0 {
			distinct++
		}
	}
	if distinct <= k {
		return 0
	}

	// Delete the least frequent characters (from the end of sorted freq)
	deletions := 0
	for i := len(freq) - 1; i >= 0 && distinct > k; i-- {
		if freq[i] > 0 {
			deletions += freq[i]
			distinct--
		}
	}
	return deletions
}
```
