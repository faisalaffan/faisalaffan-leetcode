# 2738 — Count Occurrences In Text

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountOccurrencesInText(text string, word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2738: Count Occurrences in Text
// https://leetcode.com/problems/count-occurrences-in-text/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func CountOccurrencesInText(text string, word string) int {
	count := 0
	words := strings.Fields(text)
	for _, w := range words {
		if w == word {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountOccurrencesInText("hello world hello", "hello"))
	fmt.Println(CountOccurrencesInText("this is a test test this", "test"))
	fmt.Println(CountOccurrencesInText("unique", "none"))
}
```
