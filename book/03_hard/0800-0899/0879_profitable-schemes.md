# 0879 — Profitable Schemes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func profitableSchemes(n int, minProfit int, group []int, profit []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #879: Profitable Schemes
// https://leetcode.com/problems/profitable-schemes/
// Difficulty: Hard
//
// DP knapsack. dp[p][m] = number of schemes achieving exactly profit p
// with exactly m members, after processing some subset of crimes.
// Then iterate all crimes (0/1 knapsack style, iterating backwards).
// Cap profit at minProfit since we only care about >= minProfit.

import "fmt"

const mod879 = 1_000_000_007

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	// dp[p][m] = ways to achieve profit p with m members (p capped at minProfit).
  // Matriks 2D
	dp := make([][]int, minProfit+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for idx := 0; idx < len(group); idx++ {
		g := group[idx]
		p := profit[idx]
		// Iterate backwards for 0/1 knapsack.
		for curProfit := minProfit; curProfit >= 0; curProfit-- {
			for curMembers := n - g; curMembers >= 0; curMembers-- {
				if dp[curProfit][curMembers] == 0 {
					continue
				}
				newProfit := curProfit + p
				if newProfit > minProfit {
					newProfit = minProfit
				}
				dp[newProfit][curMembers+g] = (dp[newProfit][curMembers+g] + dp[curProfit][curMembers]) % mod879
			}
		}
	}

	// Sum all schemes with profit >= minProfit (i.e., profit == minProfit after capping).
	result := 0
	for m := 0; m <= n; m++ {
		result = (result + dp[minProfit][m]) % mod879
	}
	return result
}

func main() {
	// Example 1: n=5, minProfit=3, group=[2,2], profit=[2,3] -> 2
	fmt.Println("Test 1:", profitableSchemes(5, 3, []int{2, 2}, []int{2, 3})) // 2

	// Example 2: n=0, minProfit=0, group=[], profit=[] -> 1 (empty scheme)
	fmt.Println("Test 2:", profitableSchemes(0, 0, []int{}, []int{})) // 1

	// Example 3: n=10, minProfit=5, group=[2,3,5], profit=[6,7,8] -> 7
	fmt.Println("Test 3:", profitableSchemes(10, 5, []int{2, 3, 5}, []int{6, 7, 8})) // 7

	// Edge: n=1, minProfit=1, group=[2], profit=[5] -> 0 (only crime needs 2 members, > n)
	fmt.Println("Test 4:", profitableSchemes(1, 1, []int{2}, []int{5})) // 0

	// Single crime that fits
	fmt.Println("Test 5:", profitableSchemes(3, 2, []int{2}, []int{5})) // 1
}
```
