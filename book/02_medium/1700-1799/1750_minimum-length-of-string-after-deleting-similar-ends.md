# 1750 — Minimum Length Of String After Deleting Similar Ends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumLength(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1750: Minimum Length of String After Deleting Similar Ends
// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minimumLength(s string) int {
	left, right := 0, len(s)-1

  // Two-pointer loop
	for left < right && s[left] == s[right] {
		ch := s[left]
		// Delete from left
  // Binary search loop
		for left <= right && s[left] == ch {
			left++
		}
		// Delete from right
  // Binary search loop
		for left <= right && s[right] == ch {
			right--
		}
	}
	return right - left + 1
}

func main() {
	fmt.Println(minimumLength("ca"))            // Expected: 2
	fmt.Println(minimumLength("cabaabac"))      // Expected: 0
	fmt.Println(minimumLength("aabccabba"))     // Expected: 3
}
```
