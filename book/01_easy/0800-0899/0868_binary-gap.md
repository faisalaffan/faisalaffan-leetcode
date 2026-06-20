# 0868 — Binary Gap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func binaryGap(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #868: Binary Gap
// https://leetcode.com/problems/binary-gap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(binaryGap(22))  // 2 (10110)
	fmt.Println(binaryGap(8))   // 0 (1000)
	fmt.Println(binaryGap(5))   // 2 (101)
	fmt.Println(binaryGap(6))   // 1 (110)
}

// binaryGap finds the longest distance between two consecutive 1s in binary representation.
// Time: O(log n). Space: O(1).
func binaryGap(n int) int {
	last := -1
	maxDist := 0
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				if i-last > maxDist {
					maxDist = i - last
				}
			}
			last = i
		}
		n >>= 1
	}
	return maxDist
}
```
