# 3332 — Maximum Points Tourist Can Earn

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(k * n^2) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3332: Maximum Points Tourist Can Earn
// https://leetcode.com/problems/maximum-points-tourist-can-earn/
// Difficulty: Medium
// Time: O(k * n^2) Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxScore(2, 1, [][]int{{2, 3}}, [][]int{{0, 2}, {1, 0}})) // 3
	fmt.Println(maxScore(3, 2, [][]int{{3, 4, 2}, {2, 1, 3}}, [][]int{{0, 2, 1}, {2, 0, 2}, {1, 3, 0}})) // 8
}

func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {
  // Alokasi slice
	dp := make([]int, n)
	for i := 0; i < k; i++ {
  // Alokasi slice
		ndp := make([]int, n)
		// Copy dp and add stay score
		for curr := 0; curr < n; curr++ {
			ndp[curr] = dp[curr] + stayScore[i][curr]
		}
		// Try travel from any city to any city
		for curr := 0; curr < n; curr++ {
			for dest := 0; dest < n; dest++ {
				val := dp[curr] + travelScore[curr][dest]
				if val > ndp[dest] {
					ndp[dest] = val
				}
			}
		}
		dp = ndp
	}

	maxVal := 0
	for _, v := range dp {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
```
