# 3090 — Maximum Length Substring With Two Occurrences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaximumLengthSubstringWithTwoOccurrences(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3090: Maximum Length Substring With Two Occurrences
// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maximumLengthSubstring
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("bcbbbcba")) // 4
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("aaaa"))      // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: maximumLengthSubstring
func MaximumLengthSubstringWithTwoOccurrences(s string) int {
	left := 0
  // HashMap: O(1) lookup
	freq := make(map[byte]int)
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for freq[s[right]] > 2 {
			freq[s[left]]--
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```
