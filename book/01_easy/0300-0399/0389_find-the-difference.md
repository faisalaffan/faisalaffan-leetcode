# 0389 — Find The Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindTheDifference(s, t string) byte`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #389: Find the Difference
// https://leetcode.com/problems/find-the-difference/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindTheDifference(s, t string) byte {
	var diff byte
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		diff ^= s[i]
	}
  // Linear scan O(n)
	for i := 0; i < len(t); i++ {
		diff ^= t[i]
	}
	return diff
}

func main() {
	fmt.Printf("%c\n", FindTheDifference("abcd", "abcde"))
	fmt.Printf("%c\n", FindTheDifference("", "y"))
	fmt.Printf("%c\n", FindTheDifference("a", "aa"))
}
```
