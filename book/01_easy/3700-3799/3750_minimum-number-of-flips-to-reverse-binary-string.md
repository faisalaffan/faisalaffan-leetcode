# 3750 — Minimum Number Of Flips To Reverse Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfFlipsToReverseBinaryString(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3750: Minimum Number of Flips to Reverse Binary String
// https://leetcode.com/problems/minimum-number-of-flips-to-reverse-binary-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfFlipsToReverseBinaryString(7))
	fmt.Println(MinimumNumberOfFlipsToReverseBinaryString(10))
}

// Time: O(log n)
// Space: O(1)
func MinimumNumberOfFlipsToReverseBinaryString(n int) int {
	// XOR n with its reverse, count bits
	rev := 0
	for x := n; x > 0; x >>= 1 {
		rev = (rev << 1) | (x & 1)
	}
	xor := n ^ rev
	ans := 0
	for xor > 0 {
		ans += xor & 1
		xor >>= 1
	}
	return ans
}
```
