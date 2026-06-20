# 0880 — Decoded String At Index

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func DecodedStringAtIndex(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #880: Decoded String at Index
// https://leetcode.com/problems/decoded-string-at-index/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(DecodedStringAtIndex("leet2code3", 10))
	fmt.Println(DecodedStringAtIndex("ha22", 5))
	fmt.Println(DecodedStringAtIndex("a2345678999999999999999", 1))
}

// Time: O(n) | Space: O(1)
func DecodedStringAtIndex(s string, k int) string {
	var size int64 = 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			size++
		} else {
			size *= int64(s[i] - '0')
		}
	}

	target := int64(k)
	for i := len(s) - 1; i >= 0; i-- {
		target %= size
		if target == 0 && s[i] >= 'a' && s[i] <= 'z' {
			return string(s[i])
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			size--
		} else {
			size /= int64(s[i] - '0')
		}
	}

	return ""
}
```
