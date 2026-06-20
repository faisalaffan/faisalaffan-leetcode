# 2228 — Users With Two Purchases Within Seven Days

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findUsers(purchases [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2228: Users With Two Purchases Within Seven Days
// https://leetcode.com/problems/users-with-two-purchases-within-seven-days/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findUsers(purchases [][]int) []int {
	// Group purchases by user
  // HashMap: O(1) lookup
	userPurchases := make(map[int][]int)
	for _, p := range purchases {
		userID, date := p[0], p[1]
		userPurchases[userID] = append(userPurchases[userID], date)
	}

	result := []int{}
	for userID, dates := range userPurchases {
		if len(dates) < 2 {
			continue
		}
  // Sort O(n log n)
		sort.Ints(dates)
		for i := 1; i < len(dates); i++ {
			if dates[i]-dates[i-1] <= 7 {
				result = append(result, userID)
				break
			}
		}
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	// Test case 1
	fmt.Println(findUsers([][]int{{1, 1}, {2, 2}, {1, 7}, {1, 15}, {2, 3}}))
	// Expected: [1]

	// Test case 2
	fmt.Println(findUsers([][]int{{1, 1}, {2, 1}, {3, 1}}))
	// Expected: []
}
```
