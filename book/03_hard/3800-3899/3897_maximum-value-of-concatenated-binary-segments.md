# 3897 — Maximum Value Of Concatenated Binary Segments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxValue(nums1 []int, nums0 []int) int
```

> **💡 Hint:** Sort both arrays descending. Greedily pick the largest

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3897: Maximum Value of Concatenated Binary Segments
// https://leetcode.com/problems/maximum-value-of-concatenated-binary-segments/
// Difficulty: Hard
//
// Given two arrays nums1 and nums0, select segments from each.
// Concatenate binary representations of selected elements to
// maximize the resulting value. Segments from nums1 precede
// those from nums0.
//
// Approach: Sort both arrays descending. Greedily pick the largest
// elements since binary concatenation favors larger values in
// higher positions.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxValue([]int{3, 5}, []int{2, 4}))
	// Example 2
	fmt.Println(maxValue([]int{1, 2}, []int{3, 4}))
	// Edge: single elements
	fmt.Println(maxValue([]int{7}, []int{1}))
	// Edge: same values
	fmt.Println(maxValue([]int{5, 5}, []int{5, 5}))
}

func maxValue(nums1 []int, nums0 []int) int {
  // Custom sort dengan comparator
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] > nums1[j] })
  // Custom sort dengan comparator
	sort.Slice(nums0, func(i, j int) bool { return nums0[i] > nums0[j] })

	// Try all possible split points: take first i from nums1, rest from nums0
	best := 0
	for i := 0; i <= len(nums1); i++ {
		for j := 0; j <= len(nums0); j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := 0
			for k := 0; k < i; k++ {
				val = appendBits(val, nums1[k])
			}
			for k := 0; k < j; k++ {
				val = appendBits(val, nums0[k])
			}
			if val > best {
				best = val
			}
		}
	}
	return best
}

func appendBits(current, val int) int {
	if val == 0 {
		return current << 1
	}
	bits := 0
	tmp := val
	for tmp > 0 {
		bits++
		tmp >>= 1
	}
	return (current << bits) | val
}
```
