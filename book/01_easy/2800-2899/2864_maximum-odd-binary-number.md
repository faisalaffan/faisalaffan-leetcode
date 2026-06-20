# 2864 — Maximum Odd Binary Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaximumOddBinaryNumber(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2864: Maximum Odd Binary Number
// https://leetcode.com/problems/maximum-odd-binary-number/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumOddBinaryNumber("010"))
	fmt.Println(MaximumOddBinaryNumber("0101"))
}

func MaximumOddBinaryNumber(s string) string {
	ones := strings.Count(s, "1")
	zeros := len(s) - ones
	return strings.Repeat("1", ones-1) + strings.Repeat("0", zeros) + "1"
}
```
