# 3407 — Substring Matching Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func SubstringMatchingPattern(s string, p string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n * m). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3407: Substring Matching Pattern
// https://leetcode.com/problems/substring-matching-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SubstringMatchingPattern("leetcode", "ee*e"))
	fmt.Println(SubstringMatchingPattern("car", "c*r"))
	fmt.Println(SubstringMatchingPattern("test", "t*t"))
}

// SubstringMatchingPattern checks if s matches pattern p where '*' matches any sequence of characters.
// Time: O(n * m). Space: O(n).
func SubstringMatchingPattern(s string, p string) bool {
	starIdx := -1
	for i, ch := range p {
		if ch == '*' {
			starIdx = i
			break
		}
	}

	left := p[:starIdx]
	right := p[starIdx+1:]

	// Left part must be prefix of some substring, right part must be suffix
	return strings.Contains(s, left) && strings.Contains(s, right) &&
		strings.Index(s, left) <= len(s)-len(right) &&
		strings.Index(s, left)+len(left) <= strings.LastIndex(s, right)
}
```
