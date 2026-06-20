# 0693 — Binary Number With Alternating Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func hasAlternatingBits(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #693: Binary Number with Alternating Bits
// https://leetcode.com/problems/binary-number-with-alternating-bits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(5))  // true (101)
	fmt.Println(hasAlternatingBits(7))  // false (111)
	fmt.Println(hasAlternatingBits(11)) // false (1011)
	fmt.Println(hasAlternatingBits(10)) // true (1010)
}

// hasAlternatingBits checks if the binary representation of n has alternating bits.
// Time: O(log n). Space: O(1).
func hasAlternatingBits(n int) bool {
	// XOR with n>>1 gives all 1s if alternating
	x := n ^ (n >> 1)
	// Check if x is all 1s (i.e., x & (x+1) == 0)
	return x&(x+1) == 0
}
```
