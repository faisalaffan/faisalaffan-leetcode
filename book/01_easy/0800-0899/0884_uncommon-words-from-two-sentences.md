# 0884 — Uncommon Words From Two Sentences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func uncommonFromSentences(s1 string, s2 string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m). Space: O(n + m).  |  **Ruang:** O(n + m).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #884: Uncommon Words from Two Sentences
// https://leetcode.com/problems/uncommon-words-from-two-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour")) // [sweet sour]
	fmt.Println(uncommonFromSentences("apple apple", "banana"))                     // [banana]
}

// uncommonFromSentences returns all uncommon words across two sentences.
// Time: O(n + m). Space: O(n + m).
func uncommonFromSentences(s1 string, s2 string) []string {
  // HashMap: O(1) lookup
	count := make(map[string]int)
	for _, w := range strings.Fields(s1) {
		count[w]++
	}
	for _, w := range strings.Fields(s2) {
		count[w]++
	}
	result := make([]string, 0)
	for w, c := range count {
		if c == 1 {
			result = append(result, w)
		}
	}
	return result
}
```
