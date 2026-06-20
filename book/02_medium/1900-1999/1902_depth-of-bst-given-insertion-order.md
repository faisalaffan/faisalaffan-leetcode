# 1902 — Depth Of Bst Given Insertion Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxDepthBST(order []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1902: Depth of BST Given Insertion Order
// https://leetcode.com/problems/depth-of-bst-given-insertion-order/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaxDepthBST([]int{2, 1, 4, 3}))
	fmt.Println(MaxDepthBST([]int{2, 1, 3, 4}))
	fmt.Println(MaxDepthBST([]int{1, 2, 3, 4}))
}

// Time: O(n log n), Space: O(n)
func MaxDepthBST(order []int) int {
	// Use map to store depth of each value
	// Use ordered map simulation via lower/higher logic
  // Membuat map (HashMap) — pencarian O(1)
	depth := make(map[int]int)
	depth[order[0]] = 1
	ans := 1

	// We need to find the nearest smaller and larger values already inserted
	// Since Go doesn't have an ordered map, we'll maintain a sorted slice
	sorted := []int{order[0]}

	for i := 1; i < len(order); i++ {
		v := order[i]
		// Find lower (floor) and higher (ceiling)
		lower, higher := -1, -1
		lo, hi := 0, len(sorted)-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if sorted[mid] < v {
				lower = sorted[mid]
				lo = mid + 1
			} else {
				higher = sorted[mid]
				hi = mid - 1
			}
		}

		lowerDepth, higherDepth := 0, 0
		if lower != -1 {
			lowerDepth = depth[lower]
		}
		if higher != -1 {
			higherDepth = depth[higher]
		}

		curDepth := 1 + max(lowerDepth, higherDepth)
		depth[v] = curDepth
		if curDepth > ans {
			ans = curDepth
		}

		// Insert into sorted slice (maintain sorted order)
		pos := lo
		sorted = append(sorted[:pos], append([]int{v}, sorted[pos:]...)...)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
