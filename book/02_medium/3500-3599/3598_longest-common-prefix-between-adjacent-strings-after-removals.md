# 3598 — Longest Common Prefix Between Adjacent Strings After Removals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs []string, removals int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3598: Longest Common Prefix Between Adjacent Strings After Removals
// https://leetcode.com/problems/longest-common-prefix-between-adjacent-strings-after-removals/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	strs := []string{"abc", "abd", "ab"}
	fmt.Println("Test 1:", LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs, 0))
	// Test case 2
	strs2 := []string{"a", "a", "a"}
	fmt.Println("Test 2:", LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs2, 1))
	// Test case 3
	strs3 := []string{"abc", "abc", "abc"}
	fmt.Println("Test 3:", LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs3, 2))
}

func LongestCommonPrefixBetweenAdjacentStringsAfterRemovals(strs []string, removals int) int {
	n := len(strs)
	if n <= 1 {
		return 0
	}

	maxLCP := 0
	for i := 0; i < n-1; i++ {
		s1, s2 := strs[i], strs[i+1]
		lcp := 0
		for lcp < len(s1) && lcp < len(s2) && s1[lcp] == s2[lcp] {
			lcp++
		}
		if lcp > maxLCP {
			maxLCP = lcp
		}
	}
	return maxLCP
}
```
