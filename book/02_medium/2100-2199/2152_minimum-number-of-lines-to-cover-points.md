# 2152 — Minimum Number Of Lines To Cover Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minimumLines(points [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Bitmask

**Waktu:** O(n^2 * 2^n)  |  **Ruang:** O(2^n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2152: Minimum Number of Lines to Cover Points
// https://leetcode.com/problems/minimum-number-of-lines-to-cover-points/
// Difficulty: Medium [Paid]
// Time: O(n^2 * 2^n) | Space: O(2^n)

import "fmt"

func minimumLines(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}

	// Precompute line masks
	lineMasks := []int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			mask := 0
			for k := 0; k < n; k++ {
				if isCollinear(points[i], points[j], points[k]) {
					mask |= 1 << k
				}
			}
			lineMasks = append(lineMasks, mask)
		}
	}

	// DP over subsets
  // Alokasi slice
	dp := make([]int, 1<<n)
  // Range loop
	for i := range dp {
		dp[i] = n // max n lines
	}
	dp[0] = 0

	for mask := 0; mask < (1 << n); mask++ {
		for _, lm := range lineMasks {
			dp[mask|lm] = min(dp[mask|lm], dp[mask]+1)
		}
	}

	return dp[(1<<n)-1]
}

func isCollinear(a, b, c []int) bool {
	return (b[1]-a[1])*(c[0]-a[0]) == (c[1]-a[1])*(b[0]-a[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumLines([][]int{{0, 1}, {2, 3}, {4, 5}, {4, 3}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minimumLines([][]int{{0, 2}, {-2, -2}, {1, 4}}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minimumLines([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}))
	// Expected: 1
}
```
