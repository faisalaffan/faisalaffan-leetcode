# 2730 — Find The Longest Semi Repetitive Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheLongestSemiRepetitiveSubstring(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2730: Find the Longest Semi-Repetitive Substring
// https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func FindTheLongestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}

	maxLen := 0
	left := 0
	lastPair := -1

	for right := 1; right < n; right++ {
		if s[right] == s[right-1] {
			if lastPair != -1 {
				left = lastPair
			}
			lastPair = right
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("52233"))
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("0001"))
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("1111111"))
}
```
