# 1070 — Product Sales Analysis Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func productSalesAnalysisIII(sales [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1070: Product Sales Analysis III
// https://leetcode.com/problems/product-sales-analysis-iii/
// Difficulty: Medium
//
// Approach: Find first year of each product, get its sale data.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Input: sales = [(product_id, year, quantity, price), ...]
	sales := [][]int{{1, 100, 2008, 10, 5000}, {2, 100, 2009, 12, 5000}, {3, 100, 2008, 15, 5000}}
	fmt.Println(productSalesAnalysisIII(sales)) // [[100,2008,10,5000]]
}

func productSalesAnalysisIII(sales [][]int) [][]int {
	// sales[i] = [sale_id, product_id, year, quantity, price]
  // HashMap: O(1) lookup
	firstYear := make(map[int]int) // product_id -> min year
  // HashMap: O(1) lookup
	saleByProd := make(map[int][]int)

	for _, s := range sales {
		pid, year := s[1], s[2]
		if _, ok := firstYear[pid]; !ok || year < firstYear[pid] {
			firstYear[pid] = year
		}
		saleByProd[pid] = s
	}

  // Matriks 2D
	result := make([][]int, 0)
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	for _, s := range sales {
		pid, year := s[1], s[2]
		if !seen[pid] && year == firstYear[pid] {
			// Return [product_id, year, quantity, price]
			result = append(result, []int{pid, year, s[3], s[4]})
			seen[pid] = true
		}
	}

	return result
}
```
