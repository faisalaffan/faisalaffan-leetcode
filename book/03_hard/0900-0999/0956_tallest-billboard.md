# 0956 — Tallest Billboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func tallestBillboard(rods []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #956: Tallest Billboard
// https://leetcode.com/problems/tallest-billboard/
// Difficulty: Hard

import "fmt"

func tallestBillboard(rods []int) int {
	// dp[diff] = max total sum of both sides with this abs difference
	dp := map[int]int{0: 0}

	for _, r := range rods {
  // HashMap: O(1) lookup
		cur := make(map[int]int)
		for diff, total := range dp {
			// 1. skip this rod — use >= to propagate diff=0/total=0 (map zero-value)
			if total >= cur[diff] {
				cur[diff] = total
			}
			// 2. add to taller side: diff increases by r, total increases by r
			left := total + r
			if left > cur[diff+r] {
				cur[diff+r] = left
			}
			// 3. add to shorter side: abs diff changes
			right := total + r
			newDiff := diff - r
			if newDiff < 0 {
				newDiff = -newDiff
			}
			if right > cur[newDiff] {
				cur[newDiff] = right
			}
		}
		dp = cur
	}

	// dp[0] = max total sum when diff=0 (equal sides)
	// Each side height = dp[0] / 2
	return dp[0] / 2
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
	// Expected: 6

	fmt.Println("Example 2:")
	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	// Expected: 10

	fmt.Println("Example 3:")
	fmt.Println(tallestBillboard([]int{1, 2}))
	// Expected: 0
}
```
