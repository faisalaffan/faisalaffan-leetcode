# 0003 — Longest Substring Without Repeating Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func lengthOfLongestSubstring(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(min(n, alphabet_size))

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3: Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Difficulty: Medium

import "fmt"

func lengthOfLongestSubstring(s string) int {
  // HashMap: O(1) lookup
	lastSeen := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, ok := lastSeen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	// Test case 1
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3

	// Test case 2
	fmt.Println(lengthOfLongestSubstring("bbbbb")) // 1

	// Test case 3
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3
}

// Time: O(n) | Space: O(min(n, alphabet_size))
```
