# 1967 — Number Of Strings That Appear As Substrings In Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NumberOfStringsThatAppearAsSubstringsInWord(patterns []string, word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1967: Number of Strings That Appear as Substrings in Word
// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "abc", "bc", "d"}, "abc")) // 3
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "b", "c"}, "aaaaabbbbb"))  // 2
}

// Time: O(n * m), Space: O(1)
func NumberOfStringsThatAppearAsSubstringsInWord(patterns []string, word string) int {
	count := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			count++
		}
	}
	return count
}
```
