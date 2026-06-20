# 0459 — Repeated Substring Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func RepeatedSubstringPattern(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #459: Repeated Substring Pattern
// https://leetcode.com/problems/repeated-substring-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func RepeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}

func main() {
	fmt.Println(RepeatedSubstringPattern("abab"))
	fmt.Println(RepeatedSubstringPattern("aba"))
	fmt.Println(RepeatedSubstringPattern("abcabcabcabc"))
}
```
