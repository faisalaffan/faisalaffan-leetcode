# 0796 — Rotate String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func rotateString(s string, goal string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #796: Rotate String
// https://leetcode.com/problems/rotate-string/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rotateString("abcde", "cdeab")) // true
	fmt.Println(rotateString("abcde", "abced")) // false
}

// rotateString checks if goal can be obtained by rotating s.
// Time: O(n). Space: O(n).
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}
```
