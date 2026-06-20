# 0395 — Longest Substring With At Least K Repeating Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestSubstring(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n^2) worst case, O(n) average  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #395: Longest Substring with At Least K Repeating Characters
// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
// Difficulty: Medium
// Time: O(n^2) worst case, O(n) average | Space: O(n)

import "fmt"

func longestSubstring(s string, k int) int {
	return longestSubstringHelper(s, 0, len(s), k)
}

func longestSubstringHelper(s string, start, end, k int) int {
	if end-start < k {
		return 0
	}

	// Count frequencies
	freq := [26]int{}
	for i := start; i < end; i++ {
		freq[s[i]-'a']++
	}

	// Find split point where char freq < k
	for i := start; i < end; i++ {
		if freq[s[i]-'a'] < k {
			left := longestSubstringHelper(s, start, i, k)
			right := longestSubstringHelper(s, i+1, end, k)
			if left > right {
				return left
			}
			return right
		}
	}

	return end - start
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", longestSubstring("aaabb", 3))
	// Expected: 3 ("aaa")

	// Test case 2
	fmt.Println("Test 2:", longestSubstring("ababbc", 2))
	// Expected: 5 ("ababb")

	// Test case 3
	fmt.Println("Test 3:", longestSubstring("aaabbb", 3))
	// Expected: 6
}
```
