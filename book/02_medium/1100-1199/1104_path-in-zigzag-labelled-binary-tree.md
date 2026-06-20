# 1104 — Path In Zigzag Labelled Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func pathInZigZagTree(label int) []int
```

> **💡 Hint:** Find level, compute reverse label, traverse to root

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1104: Path In Zigzag Labelled Binary Tree
// https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/
// Difficulty: Medium
//
// Approach: Find level, compute reverse label, traverse to root
// Time: O(log n)
// Space: O(log n)

import "fmt"

func main() {
	fmt.Println(pathInZigZagTree(14)) // [1,3,4,14]
	fmt.Println(pathInZigZagTree(26)) // [1,2,6,10,26]
}

func pathInZigZagTree(label int) []int {
  // Alokasi slice integer
	result := make([]int, 0)

	for label > 0 {
		result = append(result, label)
		level := 0
		for (1 << level) <= label {
			level++
		}
		level--

		// In zigzag levels, the "position" is reversed
		// Min and max of this level
		minVal := 1 << level
		maxVal := (1 << (level + 1)) - 1
		// The parent of label (in zigzag) needs to find the "mirror" position
		parent := minVal + maxVal - label
		label = parent / 2
	}

	// Reverse to get root-to-leaf
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
```
