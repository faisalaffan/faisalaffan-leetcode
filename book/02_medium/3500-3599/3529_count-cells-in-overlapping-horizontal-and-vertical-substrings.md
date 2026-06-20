# 3529 — Count Cells In Overlapping Horizontal And Vertical Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountCellsInOverlappingHorizontalAndVerticalSubstrings(rows, cols int, horizontal, vertical []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3529: Count Cells in Overlapping Horizontal and Vertical Substrings
// https://leetcode.com/problems/count-cells-in-overlapping-horizontal-and-vertical-substrings/
// Difficulty: Medium
// Complexity: O(h*w) time, O(h*w) space

import "fmt"

func main() {
	// Test case 1
	h := []int{0, 2}
	v := []int{0, 2}
	fmt.Println("Test 1:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(3, 3, h, v))

	// Test case 2
	h2 := []int{0, 1}
	v2 := []int{0, 1}
	fmt.Println("Test 2:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(3, 3, h2, v2))

	// Test case 3
	h3 := []int{0}
	v3 := []int{0}
	fmt.Println("Test 3:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(1, 1, h3, v3))
}

func CountCellsInOverlappingHorizontalAndVerticalSubstrings(rows, cols int, horizontal, vertical []int) int {
	// Count cells that are in the intersection of selected horizontal and vertical ranges
  // Membuat map (HashMap) — pencarian O(1)
	hSet := make(map[int]bool)
	for _, h := range horizontal {
		hSet[h] = true
	}
  // Membuat map (HashMap) — pencarian O(1)
	vSet := make(map[int]bool)
	for _, v := range vertical {
		vSet[v] = true
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if hSet[r] && vSet[c] {
				count++
			}
		}
	}
	return count
}
```
