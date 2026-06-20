# 2185 — Counting Words With A Given Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountingWordsWithAGivenPrefix(words []string, pref string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n * m), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2185: Counting Words With a Given Prefix
// https://leetcode.com/problems/counting-words-with-a-given-prefix/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"pay", "attention", "practice", "attend"}, "at")) // 2
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"leetcode", "win", "loops", "success"}, "code"))   // 0
}

// Time: O(n * m), Space: O(1)
func CountingWordsWithAGivenPrefix(words []string, pref string) int {
	count := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			count++
		}
	}
	return count
}
```
