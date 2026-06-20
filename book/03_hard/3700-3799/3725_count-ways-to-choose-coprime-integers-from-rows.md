# 3725 — Count Ways To Choose Coprime Integers From Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countWays(mat [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3725: Count Ways to Choose Coprime Integers from Rows
// https://leetcode.com/problems/count-ways-to-choose-coprime-integers-from-rows/
// Difficulty: Hard
//
// Pick exactly one integer from each row such that GCD of all
// chosen integers is 1 (coprime). Return count modulo 1e9+7.
//
// Approach: DP over GCD values. For each row, compute new GCDs.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countWays([][]int{{2, 3}, {3, 5}}))
	// Example 2
	fmt.Println(countWays([][]int{{2, 4}, {6, 8}}))
	// Edge: single row
	fmt.Println(countWays([][]int{{1, 5}}))
	// Edge: already coprime
	fmt.Println(countWays([][]int{{1, 2}, {3, 4}}))
}

func countWays(mat [][]int) int {
	const mod = 1000000007
	m := len(mat)
	if m == 0 {
		return 0
	}

  // HashMap: O(1) lookup
	dp := make(map[int]int)
	for _, v := range mat[0] {
		dp[v] = (dp[v] + 1) % mod
	}

	for r := 1; r < m; r++ {
  // HashMap: O(1) lookup
		ndp := make(map[int]int)
		for g, cnt := range dp {
			for _, v := range mat[r] {
				ng := gcd(g, v)
				ndp[ng] = (ndp[ng] + cnt) % mod
			}
		}
		dp = ndp
	}

	return dp[1]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
