# 1784 — Check If Binary String Has At Most One Segment Of Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckOnesSegment(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1784: Check if Binary String Has at Most One Segment of Ones
// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(1)
func CheckOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func main() {
	fmt.Println(CheckOnesSegment("1001"))
	fmt.Println(CheckOnesSegment("110"))
	fmt.Println(CheckOnesSegment("1"))
}
```
