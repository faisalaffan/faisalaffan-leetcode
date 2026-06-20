# 0459 — Repeated Substring Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func RepeatedSubstringPattern(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #459: Repeated Substring Pattern
// https://leetcode.com/problems/repeated-substring-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func RepeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}

func main() {
	fmt.Println(RepeatedSubstringPattern("abab"))
	fmt.Println(RepeatedSubstringPattern("aba"))
	fmt.Println(RepeatedSubstringPattern("abcabcabcabc"))
}
```
