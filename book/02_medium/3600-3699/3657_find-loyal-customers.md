# 3657 — Find Loyal Customers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findLoyalCustomers(purchases [][]int, minPurchases int, minAmount float64) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3657: Find Loyal Customers
// https://leetcode.com/problems/find-loyal-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findLoyalCustomers(purchases [][]int, minPurchases int, minAmount float64) []int {
  // HashMap: O(1) lookup
	customerTotals := make(map[int]int)
  // HashMap: O(1) lookup
	customerCounts := make(map[int]int)

	for _, p := range purchases {
		custID, amount := p[0], p[1]
		customerTotals[custID] += amount
		customerCounts[custID]++
	}

	var loyal []int
	for id := range customerTotals {
		if customerCounts[id] >= minPurchases && float64(customerTotals[id]) >= minAmount {
			loyal = append(loyal, id)
		}
	}

  // Sort O(n log n)
	sort.Ints(loyal)
	return loyal
}

func main() {
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 300}, {2, 150}, {1, 50}}, 2, 200))
	fmt.Println(findLoyalCustomers([][]int{{1, 50}, {1, 50}}, 2, 100))
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 200}}, 2, 100))
}
```
