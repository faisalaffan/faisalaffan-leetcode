# 1680 — Concatenation Of Consecutive Binary Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func concatenatedBinary(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1680: Concatenation of Consecutive Binary Numbers
// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

const mod = 1_000_000_007

func concatenatedBinary(n int) int {
	result := 0
	lenBits := 0

	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			lenBits++
		}
		result = (result<<lenBits | i) % mod
	}
	return result
}

func main() {
	fmt.Println(concatenatedBinary(1))   // Expected: 1
	fmt.Println(concatenatedBinary(3))   // Expected: 27
	fmt.Println(concatenatedBinary(12))  // Expected: 505379714
}
```
