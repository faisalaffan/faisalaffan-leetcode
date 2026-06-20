# 0720 — Longest Word In Dictionary

## Deskripsi

**Soal:** [0720. Longest Word In Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * L + n log n)  
**Kompleksitas Ruang:** O(n * L)

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
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
