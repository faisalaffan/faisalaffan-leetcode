# 1140 — Stone Game Ii

## Deskripsi

**Soal:** [1140. Stone Game Ii](https://leetcode.com/problems/stone-game-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^3) but typically O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func stoneGameII(piles []int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1140: Stone Game II
// https://leetcode.com/problems/stone-game-ii/
// Difficulty: Medium

// DP with memoization. dp[i][m] = max stones current player can get
// from piles[i:] with current M = m.

// Time: O(n^3) but typically O(n^2)
// Space: O(n^2)

func stoneGameII(piles []int) int {
	n := len(piles)
  // Membuat slice untuk menyimpan hasil
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for m := 1; m <= n; m++ {
			best := 0
			for x := 1; x <= 2*m && i+x <= n; x++ {
				opponent := dp[i+x][max(m, x)]
				player := suffix[i] - opponent
				if player > best {
					best = player
				}
			}
			dp[i][m] = best
		}
	}
	return dp[0][1]
}

func main() {
	fmt.Printf("%d (expected: 10)\n", stoneGameII([]int{2, 7, 9, 4, 4}))
	fmt.Printf("%d (expected: 104)\n", stoneGameII([]int{1, 2, 3, 4, 5, 100}))
	fmt.Printf("%d (expected: 1)\n", stoneGameII([]int{1}))
}
```
