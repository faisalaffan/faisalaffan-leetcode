# 1876 — Substrings Of Size Three With Distinct Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountGoodSubstrings(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1876: Substrings of Size Three with Distinct Characters
// https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountGoodSubstrings(s string) int {
	count := 0
	for i := 0; i+2 < len(s); i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountGoodSubstrings("xyzzaz"))
	fmt.Println(CountGoodSubstrings("aababcabc"))
}
```
