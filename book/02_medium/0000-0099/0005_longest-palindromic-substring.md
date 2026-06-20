# 0005 — Longest Palindromic Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestPalindrome(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #5: Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/
// Difficulty: Medium

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s)-1; i++ {
		expandAroundCenter(i, i)   // odd length
		expandAroundCenter(i, i+1) // even length
	}

	return s[start : start+maxLen]
}

func main() {
	// Test case 1
	fmt.Println(longestPalindrome("babad")) // "bab" or "aba"

	// Test case 2
	fmt.Println(longestPalindrome("cbbd")) // "bb"

	// Test case 3
	fmt.Println(longestPalindrome("a")) // "a"
}

// Time: O(n^2) | Space: O(1)
```
