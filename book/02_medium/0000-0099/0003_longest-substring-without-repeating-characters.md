# 0003 — Longest Substring Without Repeating Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func lengthOfLongestSubstring(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(min(n, alphabet_size))

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3: Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Difficulty: Medium

import "fmt"

func lengthOfLongestSubstring(s string) int {
  // Membuat map (HashMap) — pencarian O(1)
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
