# 3545 — Minimum Deletions For At Most K Distinct Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumDeletionsForAtMostKDistinctCharacters(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice
	freq := make([]int, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
  // Custom sort
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
