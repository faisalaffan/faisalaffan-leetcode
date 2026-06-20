# 2351 — First Letter To Appear Twice

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FirstLetterToAppearTwice(s string) byte`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2351: First Letter to Appear Twice
// https://leetcode.com/problems/first-letter-to-appear-twice/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(string(FirstLetterToAppearTwice("abccbaacz"))) // "c"
	fmt.Println(string(FirstLetterToAppearTwice("abcdd")))      // "d"
}

func FirstLetterToAppearTwice(s string) byte {
	seen := [26]bool{}
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if seen[idx] {
			return s[i]
		}
		seen[idx] = true
	}
	return 0
}
```
