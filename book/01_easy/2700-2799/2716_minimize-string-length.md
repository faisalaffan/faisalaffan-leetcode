# 2716 — Minimize String Length

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimizeStringLength(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2716: Minimize String Length
// https://leetcode.com/problems/minimize-string-length/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinimizeStringLength("aaabc"))
	fmt.Println(MinimizeStringLength("cbbd"))
}

func MinimizeStringLength(s string) int {
	seen := map[rune]bool{}
	for _, c := range s {
		seen[c] = true
	}
	return len(seen)
}
```
