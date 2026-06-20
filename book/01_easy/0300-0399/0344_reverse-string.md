# 0344 — Reverse String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ReverseString(s []byte) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #344: Reverse String
// https://leetcode.com/problems/reverse-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ReverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s1 := []byte("hello")
	ReverseString(s1)
	fmt.Println(string(s1))

	s2 := []byte("Hannah")
	ReverseString(s2)
	fmt.Println(string(s2))
}
```
