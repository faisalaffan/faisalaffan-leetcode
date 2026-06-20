# 0790 — Domino And Tromino Tiling

## Deskripsi

**Soal:** [0790. Domino And Tromino Tiling](https://leetcode.com/problems/domino-and-tromino-tiling/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #790: Domino and Tromino Tiling
// https://leetcode.com/problems/domino-and-tromino-tiling/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numTilings(3))
	fmt.Println(numTilings(1))
	fmt.Println(numTilings(5))
}

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	mod := 1000000007
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
  // Membuat slice untuk menyimpan hasil
	dp2 := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp2[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + 2*dp2[i-1]) % mod
		dp2[i] = (dp[i-2] + dp2[i-1]) % mod
	}

	return dp[n]
}
```
