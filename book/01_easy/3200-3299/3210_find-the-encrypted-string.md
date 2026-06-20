# 3210 — Find The Encrypted String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindTheEncryptedString(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #3210: Find the Encrypted String
// https://leetcode.com/problems/find-the-encrypted-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheEncryptedString("dart", 3))
	fmt.Println(FindTheEncryptedString("aaa", 1))
	fmt.Println(FindTheEncryptedString("abcd", 5))
}

// FindTheEncryptedString returns the encrypted string by rotating each character by k positions forward.
// Time: O(n). Space: O(n).
func FindTheEncryptedString(s string, k int) string {
	n := len(s)
  // Edge case: input kosong
	if n == 0 {
		return s
	}
	k %= n
	return s[k:] + s[:k]
}
```
