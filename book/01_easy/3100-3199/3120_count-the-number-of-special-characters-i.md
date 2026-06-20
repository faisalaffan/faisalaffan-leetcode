# 3120 — Count The Number Of Special Characters I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountTheNumberOfSpecialCharactersI(word string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3120: Count the Number of Special Characters I
// https://leetcode.com/problems/count-the-number-of-special-characters-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfSpecialChars
	fmt.Println(CountTheNumberOfSpecialCharactersI("aaAbcBC")) // 3
	fmt.Println(CountTheNumberOfSpecialCharactersI("abcd"))    // 0
	fmt.Println(CountTheNumberOfSpecialCharactersI("abAB"))   // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: numberOfSpecialChars
func CountTheNumberOfSpecialCharactersI(word string) int {
  // HashMap: O(1) lookup
	lower := make(map[byte]bool)
  // HashMap: O(1) lookup
	upper := make(map[byte]bool)
  // Linear scan O(n)
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'a' && c <= 'z' {
			lower[c] = true
		} else if c >= 'A' && c <= 'Z' {
			upper[c] = true
		}
	}
	count := 0
	for c := byte('a'); c <= 'z'; c++ {
		if lower[c] && upper[c-'a'+'A'] {
			count++
		}
	}
	return count
}
```
