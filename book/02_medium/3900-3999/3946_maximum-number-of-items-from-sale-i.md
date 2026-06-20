# 3946 — Maximum Number Of Items From Sale I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaximumNumberOfItemsFromSaleI(items [][]int, budget int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(N * budget + N^2)  |  **Ruang:** O(budget) where N = len(items)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3946: Maximum Number of Items From Sale I
// https://leetcode.com/problems/maximum-number-of-items-from-sale-i/
// Difficulty: Medium
// Time: O(N * budget + N^2) | Space: O(budget) where N = len(items)
// Approach: 0-1 knapsack. Each item's first copy gives (1 + out_degree)
// copies at cost price_i. Additional copies of cheapest item give 1 each.
// out_degree[i] = count of j where factor_i divides factor_j, j != i.

import (
	"fmt"
	"math"
)

func MaximumNumberOfItemsFromSaleI(items [][]int, budget int) int {
	m := len(items)

	// Compute out_degree: how many other items this item's factor divides
  // Alokasi slice
	outDeg := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j && items[j][0]%items[i][0] == 0 {
				outDeg[i]++
			}
		}
	}

	// Find min price (for additional copies after activation)
	minPrice := math.MaxInt32
	for _, it := range items {
		if it[1] < minPrice {
			minPrice = it[1]
		}
	}

	// 0-1 knapsack: dp[b] = max copies from first copies with budget b
  // Alokasi slice
	dp := make([]int, budget+1)
	for b := 1; b <= budget; b++ {
		dp[b] = math.MinInt32
	}
	dp[0] = 0

	for i := 0; i < m; i++ {
		price := items[i][1]
		val := 1 + outDeg[i] // first copy value
		for b := budget; b >= price; b-- {
			if dp[b-price] > math.MinInt32 {
				cand := dp[b-price] + val
				if cand > dp[b] {
					dp[b] = cand
				}
			}
		}
	}

	// Best result: dp[b] + floor((budget-b)/minPrice) cheap copies
	best := 0
	for b := 0; b <= budget; b++ {
		if dp[b] > math.MinInt32 {
			total := dp[b] + (budget-b)/minPrice
			if total > best {
				best = total
			}
		}
	}

	return best
}

func main() {
	// Example 1
	fmt.Println(MaximumNumberOfItemsFromSaleI([][]int{{6, 2}, {2, 6}, {3, 4}}, 9)) // Expected: 4

	// Example 2
	fmt.Println(MaximumNumberOfItemsFromSaleI([][]int{{2, 4}, {3, 2}, {4, 1}, {6, 4}, {12, 4}}, 8)) // Expected: 10
}
```
