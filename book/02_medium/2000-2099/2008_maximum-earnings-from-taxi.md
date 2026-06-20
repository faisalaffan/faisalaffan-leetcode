# 2008 — Maximum Earnings From Taxi

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaximumEarningsFromTaxi(n int, rides [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Sorting

**Waktu:** O(m log m), Space: O(m)  |  **Ruang:** O(m)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2008: Maximum Earnings From Taxi
// https://leetcode.com/problems/maximum-earnings-from-taxi/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumEarningsFromTaxi(5, [][]int{{2, 5, 4}, {1, 5, 1}}))
	fmt.Println(MaximumEarningsFromTaxi(20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}))
}

// Time: O(m log m), Space: O(m)
func MaximumEarningsFromTaxi(n int, rides [][]int) int64 {
	m := len(rides)
  // Custom sort
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})

  // Alokasi slice
	dp := make([]int64, m+1)
  // Alokasi slice
	endTimes := make([]int, m)
	for i := 0; i < m; i++ {
		endTimes[i] = rides[i][1]
	}

	for i := 1; i <= m; i++ {
		start, end, tip := rides[i-1][0], rides[i-1][1], rides[i-1][2]
		earn := int64(end - start + tip)
		dp[i] = dp[i-1]

		j := sort.Search(m, func(k int) bool {
			return rides[k][1] > start
		})
		if dp[j]+earn > dp[i] {
			dp[i] = dp[j] + earn
		}
	}

	return dp[m]
}
```
