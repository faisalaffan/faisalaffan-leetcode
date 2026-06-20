# 0696 — Count Binary Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countBinarySubstrings(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #696: Count Binary Substrings
// https://leetcode.com/problems/count-binary-substrings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("00110011")) // 6
	fmt.Println(countBinarySubstrings("10101"))    // 4
	fmt.Println(countBinarySubstrings("00110"))    // 3
}

// countBinarySubstrings counts substrings that have equal numbers of 0s and 1s.
// Time: O(n). Space: O(1).
func countBinarySubstrings(s string) int {
	prev, curr, result := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			prev = curr
			curr = 1
		}
		if prev >= curr {
			result++
		}
	}
	return result
}
```
