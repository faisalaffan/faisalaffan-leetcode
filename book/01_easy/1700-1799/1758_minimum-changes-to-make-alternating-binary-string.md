# 1758 — Minimum Changes To Make Alternating Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinOperations(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1758: Minimum Changes to Make Alternating Binary String
// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MinOperations(s string) int {
	changes := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		expected := byte('0' + i%2)
		if s[i] != expected {
			changes++
		}
	}
	if changes < len(s)-changes {
		return changes
	}
	return len(s) - changes
}

func main() {
	fmt.Println(MinOperations("0100"))
	fmt.Println(MinOperations("10"))
	fmt.Println(MinOperations("1111"))
}
```
