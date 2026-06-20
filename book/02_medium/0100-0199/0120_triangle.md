# 0120 — Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minimumTotal(triangle [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #120: Triangle
// https://leetcode.com/problems/triangle/
// Difficulty: Medium

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
  // Alokasi slice
	dp := make([]int, n)
	copy(dp, triangle[n-1])

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if dp[j] < dp[j+1] {
				dp[j] = triangle[i][j] + dp[j]
			} else {
				dp[j] = triangle[i][j] + dp[j+1]
			}
		}
	}

	return dp[0]
}

func main() {
	// Test case 1
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})) // 11

	// Test case 2
	fmt.Println(minimumTotal([][]int{{-10}})) // -10

	// Test case 3
	fmt.Println(minimumTotal([][]int{{1}, {2, 3}})) // 3
}

// Time: O(n^2) | Space: O(n)
```
