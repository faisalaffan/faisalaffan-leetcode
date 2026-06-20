# 0819 — Most Common Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func mostCommonWord(paragraph string, banned []string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m). Space: O(n + m).  |  **Ruang:** O(n + m).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #819: Most Common Word
// https://leetcode.com/problems/most-common-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})) // "ball"
	fmt.Println(mostCommonWord("a.", []string{}))                                                          // "a"
}

// mostCommonWord finds the most frequent word in paragraph that is not banned.
// Time: O(n + m). Space: O(n + m).
func mostCommonWord(paragraph string, banned []string) string {
  // HashMap: O(1) lookup
	bannedSet := make(map[string]bool)
	for _, w := range banned {
		bannedSet[w] = true
	}

	// Normalize: lowercase and split by non-letter characters
	normalized := strings.ToLower(paragraph)
	cleaned := strings.NewReplacer("!", " ", "?", " ", "'", " ", ",", " ", ";", " ", ".", " ").Replace(normalized)

	words := strings.Fields(cleaned)
  // HashMap: O(1) lookup
	counts := make(map[string]int)
	maxCount := 0
	result := ""
	for _, w := range words {
		if bannedSet[w] {
			continue
		}
		counts[w]++
		if counts[w] > maxCount {
			maxCount = counts[w]
			result = w
		}
	}
	return result
}
```
