# 0401 — Binary Watch

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countBits(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #401: Binary Watch
// https://leetcode.com/problems/binary-watch/
// Difficulty: Easy

import (
	"fmt"
)

func countBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

// Time: O(1), Space: O(1)
func BinaryWatch(turnedOn int) []string {
	var result []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if countBits(h)+countBits(m) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return result
}

func main() {
	fmt.Println(BinaryWatch(1))
	fmt.Println(BinaryWatch(9))
}
```
