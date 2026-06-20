# 2744 — Find Maximum Number Of String Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindMaximumNumberOfStringPairs(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2744: Find Maximum Number of String Pairs
// https://leetcode.com/problems/find-maximum-number-of-string-pairs/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindMaximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
	fmt.Println(FindMaximumNumberOfStringPairs([]string{"ab", "ba", "cc"}))
}

func FindMaximumNumberOfStringPairs(words []string) int {
	seen := map[string]bool{}
	count := 0
	for _, w := range words {
		rev := reverse(w)
		if seen[rev] {
			count++
		}
		seen[w] = true
	}
	return count
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```
