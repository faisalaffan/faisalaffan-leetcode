# 2796 — Repeat String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func RepeatString(s string, n int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2796: Repeat String
// https://leetcode.com/problems/repeat-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Repeats string n times.

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RepeatString("abc", 3))
	fmt.Println(RepeatString("x", 5))
}

func RepeatString(s string, n int) string {
	return strings.Repeat(s, n)
}
```
