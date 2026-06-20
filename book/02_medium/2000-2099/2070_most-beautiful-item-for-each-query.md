# 2070 — Most Beautiful Item For Each Query

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumBeauty(items [][]int, queries []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting, Monotonic Stack

**Waktu:** O((n+q) log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2070: Most Beautiful Item for Each Query
// https://leetcode.com/problems/most-beautiful-item-for-each-query/
// Difficulty: Medium
// Time: O((n+q) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	// Sort items by price
  // Custom sort
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	// For each price, keep max beauty so far (monotonic)
	type item struct{ price, beauty int }
	filtered := []item{}
	maxBeauty := 0
	for _, it := range items {
		if it[1] > maxBeauty {
			maxBeauty = it[1]
		}
		// Only add if beauty increases (since sorted by price)
		if len(filtered) == 0 || it[1] > filtered[len(filtered)-1].beauty {
			filtered = append(filtered, item{it[0], maxBeauty})
		}
	}

	// Handle queries
  // Alokasi slice
	result := make([]int, len(queries))
	for i, q := range queries {
		// Binary search for last item with price <= q
		lo, hi := 0, len(filtered)-1
		best := 0
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if filtered[mid].price <= q {
				best = filtered[mid].beauty
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		result[i] = best
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumBeauty([][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))
	// Expected: [2, 4, 5, 5, 6, 6]

	// Test case 2
	fmt.Println("Test 2:", maximumBeauty([][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}, []int{1}))
	// Expected: [4]

	// Test case 3
	fmt.Println("Test 3:", maximumBeauty([][]int{{10, 100}}, []int{5, 10, 15}))
	// Expected: [0, 100, 100]
}
```
