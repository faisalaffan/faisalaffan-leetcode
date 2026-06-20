# 3794 — Reverse String Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ReverseStringPrefix(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3794: Reverse String Prefix
// https://leetcode.com/problems/reverse-string-prefix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseStringPrefix("abcd", 2))
	fmt.Println(ReverseStringPrefix("xyz", 3))
	fmt.Println(ReverseStringPrefix("hey", 1))
}

// Time: O(n)
// Space: O(n)
func ReverseStringPrefix(s string, k int) string {
	b := []byte(s)
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
```
