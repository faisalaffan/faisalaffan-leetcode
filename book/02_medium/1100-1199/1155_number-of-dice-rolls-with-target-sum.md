# 1155 — Number Of Dice Rolls With Target Sum

## Deskripsi

**Soal:** [1155. Number Of Dice Rolls With Target Sum](https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * target * k)  
**Kompleksitas Ruang:** O(target) with 1D DP optimization

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func numRollsToTarget(n int, k int, target int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1155: Number of Dice Rolls With Target Sum
// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
// Difficulty: Medium

// dp[d][t] = number of ways to get sum t using d dice with k faces.
// dp[d][t] = sum(dp[d-1][t-f] for f in 1..k if t-f >= 0)

// Time: O(n * target * k)
// Space: O(target) with 1D DP optimization

func numRollsToTarget(n int, k int, target int) int {
	const mod = 1_000_000_007

  // Membuat slice untuk menyimpan hasil
	dp := make([]int, target+1)
	dp[0] = 1

	for dice := 0; dice < n; dice++ {
  // Membuat slice untuk menyimpan hasil
		next := make([]int, target+1)
		for sum := 0; sum <= target; sum++ {
			if dp[sum] == 0 {
				continue
			}
			for face := 1; face <= k && sum+face <= target; face++ {
				next[sum+face] = (next[sum+face] + dp[sum]) % mod
			}
		}
		dp = next
	}
	return dp[target]
}

func main() {
	fmt.Printf("%d (expected: 1)\n", numRollsToTarget(1, 6, 3))
	fmt.Printf("%d (expected: 6)\n", numRollsToTarget(2, 6, 7))
	fmt.Printf("%d (expected: 222616187)\n", numRollsToTarget(30, 30, 500))
}
```
