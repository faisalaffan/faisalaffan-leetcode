# 3760 — Maximum Substrings With Distinct Start

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumSubstringsWithDistinctStart(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3760: Maximum Substrings With Distinct Start
// https://leetcode.com/problems/maximum-substrings-with-distinct-start/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSubstringsWithDistinctStart(s string) int {
	seen := [26]bool{}
	ans := 0
	for _, ch := range s {
		idx := ch - 'a'
		if !seen[idx] {
			seen[idx] = true
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumSubstringsWithDistinctStart("abacaba"))
	fmt.Println(maximumSubstringsWithDistinctStart("aaaa"))
	fmt.Println(maximumSubstringsWithDistinctStart("abc"))
}
```
