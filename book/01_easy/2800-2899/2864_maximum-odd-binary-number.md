# 2864 — Maximum Odd Binary Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumOddBinaryNumber(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2864: Maximum Odd Binary Number
// https://leetcode.com/problems/maximum-odd-binary-number/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumOddBinaryNumber("010"))
	fmt.Println(MaximumOddBinaryNumber("0101"))
}

func MaximumOddBinaryNumber(s string) string {
	ones := strings.Count(s, "1")
	zeros := len(s) - ones
	return strings.Repeat("1", ones-1) + strings.Repeat("0", zeros) + "1"
}
```
