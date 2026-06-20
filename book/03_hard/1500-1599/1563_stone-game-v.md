# 1563 — Stone Game V

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func stoneGameV(stoneValue []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1563: Stone Game V
// https://leetcode.com/problems/stone-game-v/
// Difficulty: Hard
//
// DP interval approach:
// 1. Compute prefix sums for O(1) range sum queries.
// 2. DP[i][j] = maximum score Alice can get from stones[i..j].
// 3. For each split k in [i, j-1]:
//    - leftSum = sum[i..k], rightSum = sum[k+1..j]
//    - If leftSum < rightSum: score = DP[i][k] + leftSum
//    - If leftSum > rightSum: score = DP[k+1][j] + rightSum
//    - If equal: score = max(DP[i][k], DP[k+1][j]) + leftSum
// 4. Take max over all splits.

import (
	"fmt"
)

func main() {
	// Example: [6,2,3,4,5,5] -> 18
	fmt.Println(stoneGameV([]int{6, 2, 3, 4, 5, 5}))

	// Additional tests
	fmt.Println(stoneGameV([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(stoneGameV([]int{1}))
	fmt.Println(stoneGameV([]int{1, 2}))
	fmt.Println(stoneGameV([]int{3, 3, 3}))
}

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n == 1 {
		return 0
	}

	// Prefix sums
  // Alokasi slice
	prefix := make([]int, n+1)
	for i, v := range stoneValue {
		prefix[i+1] = prefix[i] + v
	}

  // Matriks 2D
	dp := make([][]int, n)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			best := 0
			for k := i; k < j; k++ {
				leftSum := prefix[k+1] - prefix[i]
				rightSum := prefix[j+1] - prefix[k+1]

				var score int
				if leftSum < rightSum {
					score = dp[i][k] + leftSum
				} else if leftSum > rightSum {
					score = dp[k+1][j] + rightSum
				} else {
					score = max(dp[i][k], dp[k+1][j]) + leftSum
				}

				if score > best {
					best = score
				}
			}
			dp[i][j] = best
		}
	}

	return dp[0][n-1]
}
```
