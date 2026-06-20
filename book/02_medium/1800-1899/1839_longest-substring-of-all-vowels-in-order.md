# 1839 — Longest Substring Of All Vowels In Order

## Deskripsi

**Soal:** [1839. Longest Substring Of All Vowels In Order](https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func longestBeautifulSubstring(word string) int`

## Solusi Go

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
