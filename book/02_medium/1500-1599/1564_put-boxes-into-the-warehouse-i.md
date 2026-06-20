# 1564 — Put Boxes Into The Warehouse I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaxBoxesInWarehouse(boxes []int, warehouse []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(N log N + M), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1564: Put Boxes Into the Warehouse I
// https://leetcode.com/problems/put-boxes-into-the-warehouse-i/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxBoxesInWarehouse([]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}))
	fmt.Println(MaxBoxesInWarehouse([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(MaxBoxesInWarehouse([]int{1, 2, 2, 3, 4}, []int{3, 4, 1, 2}))
}

func MaxBoxesInWarehouse(boxes []int, warehouse []int) int {
	// Time: O(N log N + M), Space: O(1)
  // Sort O(n log n)
	sort.Ints(boxes)

	// Preprocess warehouse: each position's max usable height
	// is min of itself and all previous positions' heights
	for i := 1; i < len(warehouse); i++ {
		if warehouse[i] > warehouse[i-1] {
			warehouse[i] = warehouse[i-1]
		}
	}

	boxIdx := 0
	// Try to fit boxes from the largest to smallest, entering from right
	for i := len(warehouse) - 1; i >= 0 && boxIdx < len(boxes); i-- {
		if boxes[boxIdx] <= warehouse[i] {
			boxIdx++
		}
	}

	return boxIdx
}
```
