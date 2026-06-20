# 0527 — Word Abbreviation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func wordsAbbreviation(words []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #527: Word Abbreviation
// https://leetcode.com/problems/word-abbreviation/
// Difficulty: Hard

import (
	"fmt"
	"strconv"
)

func main() {
	input := []string{"like", "god", "internal", "me", "internet", "interval", "intension", "face", "intrusion"}
	output := wordsAbbreviation(input)
	fmt.Println(output)
	// Expected: ["l2e","god","internal","me","i6t","interval","inte4n","f4e","intr4n"]
}

func wordsAbbreviation(words []string) []string {
	n := len(words)
	ans := make([]string, n)
  // Alokasi slice
	prefix := make([]int, n)

	for i := 0; i < n; i++ {
		prefix[i] = 1
		ans[i] = abbreviate(words[i], 1)
	}

	for i := 0; i < n; i++ {
		for {
			conflict := false
			for j := i + 1; j < n; j++ {
				if ans[i] == ans[j] {
					prefix[j]++
					ans[j] = abbreviate(words[j], prefix[j])
					conflict = true
				}
			}
			if conflict {
				prefix[i]++
				ans[i] = abbreviate(words[i], prefix[i])
			} else {
				break
			}
		}
	}
	return ans
}

func abbreviate(s string, k int) string {
	if k >= len(s)-2 {
		return s
	}
	abbr := s[:k] + strconv.Itoa(len(s)-k-1) + s[len(s)-1:]
	if len(abbr) >= len(s) {
		return s
	}
	return abbr
}
```
