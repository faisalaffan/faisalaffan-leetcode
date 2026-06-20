# 2548 — Maximum Price To Fill A Bag

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxPrice(items [][]int, capacity int) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2548: Maximum Price to Fill a Bag
// https://leetcode.com/problems/maximum-price-to-fill-a-bag/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxPrice(items [][]int, capacity int) float64 {
	// items[i] = [price, weight]
	// Sort by price/weight ratio descending
  // Custom sort
	sort.Slice(items, func(i, j int) bool {
		return float64(items[i][0])/float64(items[i][1]) > float64(items[j][0])/float64(items[j][1])
	})

	var totalPrice float64
	remaining := capacity

	for _, item := range items {
		if remaining <= 0 {
			break
		}
		price, weight := item[0], item[1]
		take := weight
		if take > remaining {
			take = remaining
		}
		totalPrice += float64(price) * float64(take) / float64(weight)
		remaining -= take
	}

	if remaining > 0 {
		return -1
	}
	return totalPrice
}

func main() {
	// Test case 1
	fmt.Printf("Test 1: %.5f\n", maxPrice([][]int{{50, 10}, {100, 20}, {120, 30}}, 50))
	// Expected: 240.00000

	// Test case 2: can't fill
	fmt.Printf("Test 2: %.5f\n", maxPrice([][]int{{50, 10}}, 20))
	// Expected: -1

	// Test case 3: exact fit
	fmt.Printf("Test 3: %.5f\n", maxPrice([][]int{{60, 10}, {100, 20}}, 30))
	// Expected: 160.00000
}
```
