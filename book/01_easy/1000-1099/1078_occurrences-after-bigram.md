# 1078 — Occurrences After Bigram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func findOcurrences(text, first, second string) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1078: Occurrences After Bigram
// https://leetcode.com/problems/occurrences-after-bigram/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	// ["girl","student"]
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
	// ["we","rock"]
}

// LeetCode submission: findOcurrences
func findOcurrences(text, first, second string) []string {
	words := strings.Fields(text)
	var ans []string
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ans = append(ans, words[i])
		}
	}
	return ans
}
```
