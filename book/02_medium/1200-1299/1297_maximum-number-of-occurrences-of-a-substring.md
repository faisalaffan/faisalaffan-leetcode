# 1297 — Maximum Number Of Occurrences Of A Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func maxFreq(s string, maxLetters int, minSize int, maxSize int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * minSize)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1297: Maximum Number of Occurrences of a Substring
// https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/
// Difficulty: Medium

// Find max occurrences of any substring meeting constraints:
// - size between 1 and maxLetters distinct characters
// - substring size = minSize (longer substrings are less frequent)

// Time: O(n * minSize)
// Space: O(n)

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
  // HashMap: O(1) lookup
	count := make(map[string]int)
	maxOccur := 0

	for i := 0; i+minSize <= len(s); i++ {
		sub := s[i : i+minSize]
		// Count distinct chars
  // HashMap: O(1) lookup
		chars := make(map[byte]bool)
		for j := 0; j < len(sub); j++ {
			chars[sub[j]] = true
		}
		if len(chars) <= maxLetters {
			count[sub]++
			if count[sub] > maxOccur {
				maxOccur = count[sub]
			}
		}
	}

	return maxOccur
}

func main() {
	fmt.Printf("%d (expected: 2)\n", maxFreq("aababcaab", 2, 3, 4))
	fmt.Printf("%d (expected: 2)\n", maxFreq("aaaa", 1, 3, 3))
	fmt.Printf("%d (expected: 3)\n", maxFreq("aabcabcab", 2, 3, 3))
}
```
