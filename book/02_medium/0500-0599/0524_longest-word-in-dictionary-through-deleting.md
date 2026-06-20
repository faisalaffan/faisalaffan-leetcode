# 0524 — Longest Word In Dictionary Through Deleting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindLongestWord(s string, dictionary []string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n * L) where n = len(dictionary), L = max(len(s), len(word))  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #524: Longest Word in Dictionary through Deleting
// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/
// Difficulty: Medium
// Time: O(n * L) where n = len(dictionary), L = max(len(s), len(word))
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(FindLongestWord("abpcplea", []string{"a", "b", "c"}))
}

func FindLongestWord(s string, dictionary []string) string {
	// Sort by length desc, then lexicographically
  // Custom sort
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		}
		return dictionary[i] < dictionary[j]
	})

	for _, word := range dictionary {
		if isSubsequence(word, s) {
			return word
		}
	}

	return ""
}

func isSubsequence(word, s string) bool {
	i := 0
	for j := 0; i < len(word) && j < len(s); j++ {
		if word[i] == s[j] {
			i++
		}
	}
	return i == len(word)
}
```
