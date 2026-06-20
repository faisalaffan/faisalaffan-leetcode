# 3365 — Rearrange K Substrings To Form Target String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isPossibleToRearrange(s string, t string, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3365: Rearrange K Substrings to Form Target String
// https://leetcode.com/problems/rearrange-k-substrings-to-form-target-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(isPossibleToRearrange("abcd", "cdab", 2)) // true
	fmt.Println(isPossibleToRearrange("aabb", "bbaa", 2)) // true
	fmt.Println(isPossibleToRearrange("abcd", "acbd", 2)) // false
}

func isPossibleToRearrange(s string, t string, k int) bool {
	n := len(s)
	m := n / k

  // HashMap: O(1) lookup
	cnt := make(map[string]int)
	for i := 0; i < n; i += m {
		cnt[s[i:i+m]]++
		cnt[t[i:i+m]]--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
```
