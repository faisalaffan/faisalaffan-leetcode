# 0686 — Repeated String Match

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func repeatedStringMatch(a string, b string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m) worst case  |  **Ruang:** O(n + m)


## 💻 Solusi Go

```go
package main

// LeetCode #686: Repeated String Match
// https://leetcode.com/problems/repeated-string-match/
// Difficulty: Medium
// Time: O(n * m) worst case
// Space: O(n + m)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("a", "aa"))
	fmt.Println(repeatedStringMatch("abc", "wxyz"))
}

func repeatedStringMatch(a string, b string) int {
	maxRepeats := len(b)/len(a) + 3
	var sb strings.Builder

	for i := 1; i <= maxRepeats; i++ {
		sb.WriteString(a)
		if strings.Contains(sb.String(), b) {
			return i
		}
	}

	return -1
}
```
