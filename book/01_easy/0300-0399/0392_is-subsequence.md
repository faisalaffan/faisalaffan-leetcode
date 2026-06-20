# 0392 — Is Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func IsSubsequence(s, t string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n+m), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #392: Is Subsequence
// https://leetcode.com/problems/is-subsequence/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(1)
func IsSubsequence(s, t string) bool {
	i := 0
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func main() {
	fmt.Println(IsSubsequence("abc", "ahbgdc"))
	fmt.Println(IsSubsequence("axc", "ahbgdc"))
	fmt.Println(IsSubsequence("", "ahbgdc"))
}
```
