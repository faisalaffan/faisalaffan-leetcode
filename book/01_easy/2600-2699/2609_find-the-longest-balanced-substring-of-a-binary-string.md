# 2609 — Find The Longest Balanced Substring Of A Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheLongestBalancedSubstringOfABinaryString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2609: Find the Longest Balanced Substring of a Binary String
// https://leetcode.com/problems/find-the-longest-balanced-substring-of-a-binary-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("01000111")) // 6
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("00111"))    // 4
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("111"))      // 0
}

func FindTheLongestBalancedSubstringOfABinaryString(s string) int {
	maxLen := 0
	i := 0
	n := len(s)

	for i < n {
		zeros, ones := 0, 0
		for i < n && s[i] == '0' {
			zeros++
			i++
		}
		for i < n && s[i] == '1' {
			ones++
			i++
		}
		pairLen := min(zeros, ones) * 2
		if pairLen > maxLen {
			maxLen = pairLen
		}
	}
	return maxLen
}
```
