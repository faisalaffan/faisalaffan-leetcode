# 1641 — Count Sorted Vowel Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountVowelStrings(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1641: Count Sorted Vowel Strings
// https://leetcode.com/problems/count-sorted-vowel-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountVowelStrings(1))
	fmt.Println(CountVowelStrings(2))
	fmt.Println(CountVowelStrings(33))
}

func CountVowelStrings(n int) int {
	// Time: O(N), Space: O(1)
	// Combinatorics: C(n+4, 4) = (n+4)*(n+3)*(n+2)*(n+1)/24
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}
```
