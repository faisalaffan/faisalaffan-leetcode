# 1839 — Longest Substring Of All Vowels In Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestBeautifulSubstring(word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1839: Longest Substring Of All Vowels in Order
// https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func longestBeautifulSubstring(word string) int {
	vowels := "aeiou"
	maxLen := 0
	i := 0
	n := len(word)

	for i < n {
		// Start of a new substring
		vowelIdx := 0
		start := i

		// Check if starts with 'a'
		if word[i] != 'a' {
			i++
			continue
		}

		for i < n && vowelIdx < 5 {
			if word[i] == vowels[vowelIdx] {
				i++
			} else if vowelIdx+1 < 5 && word[i] == vowels[vowelIdx+1] {
				vowelIdx++
				i++
			} else {
				break
			}
		}

		if vowelIdx == 4 {
			length := i - start
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu")) // Expected: 13
	fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou")) // Expected: 5
	fmt.Println(longestBeautifulSubstring("aaaa")) // Expected: 0 (no 'e')
}
```
