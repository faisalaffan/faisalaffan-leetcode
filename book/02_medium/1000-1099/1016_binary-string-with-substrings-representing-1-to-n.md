# 1016 — Binary String With Substrings Representing 1 To N

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func queryString(s string, n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * len(s) * log n)  |  **Ruang:** O(log n)


## 💻 Solusi Go

```go
package main

// LeetCode #1016: Binary String With Substrings Representing 1 To N
// https://leetcode.com/problems/binary-string-with-substrings-representing-1-to-n/
// Difficulty: Medium
//
// Approach: Check if binary representation of each number from n down to 1
//           is a substring of s. Start from n and go down for efficiency.
// Time: O(n * len(s) * log n)
// Space: O(log n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(queryString("0110", 3))  // true
	fmt.Println(queryString("0110", 4))  // false
	fmt.Println(queryString("1", 1))     // true
}

func queryString(s string, n int) bool {
	for i := n; i >= 1; i-- {
		binary := fmt.Sprintf("%b", i)
		if !strings.Contains(s, binary) {
			return false
		}
	}
	return true
}
```
