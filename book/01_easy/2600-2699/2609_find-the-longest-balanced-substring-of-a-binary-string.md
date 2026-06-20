# 2609 — Find The Longest Balanced Substring Of A Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindTheLongestBalancedSubstringOfABinaryString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2609: Find the Longest Balanced Substring of a Binary String
// https://leetcode.com/problems/find-the-longest-balanced-substring-of-a-binary-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("01000111")) // 6
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("00111"))    // 4
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("111"))      // 0
}

func FindTheLongestBalancedSubstringOfABinaryString(s string) int {
	maxLen := 0
	i := 0
	n := len(s)

	for i < n {
		zeros, ones := 0, 0
		for i < n && s[i] == '0' {
			zeros++
			i++
		}
		for i < n && s[i] == '1' {
			ones++
			i++
		}
		pairLen := min(zeros, ones) * 2
		if pairLen > maxLen {
			maxLen = pairLen
		}
	}
	return maxLen
}
```
