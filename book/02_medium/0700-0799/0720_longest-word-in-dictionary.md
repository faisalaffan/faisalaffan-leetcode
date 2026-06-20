# 0720 — Longest Word In Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestWord(words []string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n * L + n log n)  |  **Ruang:** O(n * L)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #720: Longest Word in Dictionary
// https://leetcode.com/problems/longest-word-in-dictionary/
// Difficulty: Medium
// Time: O(n * L + n log n)
// Space: O(n * L)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}

func longestWord(words []string) string {
  // HashMap: O(1) lookup
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}

	sort.Strings(words)

	result := ""
	for _, w := range words {
		if len(w) <= len(result) {
			continue
		}
		valid := true
		for i := 1; i < len(w); i++ {
			if !wordSet[w[:i]] {
				valid = false
				break
			}
		}
		if valid {
			result = w
		}
	}

	return result
}
```
