# 3856 — Trim Trailing Vowels

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func TrimTrailingVowels(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3856: Trim Trailing Vowels
// https://leetcode.com/problems/trim-trailing-vowels/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TrimTrailingVowels("idea"))
	fmt.Println(TrimTrailingVowels("day"))
	fmt.Println(TrimTrailingVowels("aeiou"))
}

// Time: O(n)
// Space: O(n)
func TrimTrailingVowels(s string) string {
	i := len(s) - 1
	for i >= 0 && (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u') {
		i--
	}
	return s[:i+1]
}
```
