# 3230 — Customer Purchasing Behavior Analysis

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func customerPurchasingBehavior(purchases [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3230: Customer Purchasing Behavior Analysis
// https://leetcode.com/problems/customer-purchasing-behavior-analysis/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func customerPurchasingBehavior(purchases [][]int) []int {
  // HashMap: O(1) lookup
	counts := make(map[int]int)
	for _, p := range purchases {
		counts[p[0]]++
	}

	var customers []int
	for id := range counts {
		customers = append(customers, id)
	}
  // Custom sort
	sort.Slice(customers, func(i, j int) bool {
		if counts[customers[i]] != counts[customers[j]] {
			return counts[customers[i]] > counts[customers[j]]
		}
		return customers[i] < customers[j]
	})

  // Alokasi slice
	ans := make([]int, len(customers))
	for i, id := range customers {
		ans[i] = id
	}
	return ans
}

func main() {
	fmt.Println(customerPurchasingBehavior([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 75}})) // Expected: [1 2 3]
	fmt.Println(customerPurchasingBehavior([][]int{{1, 10}, {2, 20}, {2, 30}}))             // Expected: [2 1]
}
```
