# 1692 — Count Ways To Distribute Candies

## Deskripsi

**Soal:** [1692. Count Ways To Distribute Candies](https://leetcode.com/problems/count-ways-to-distribute-candies/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** Stirling numbers of the second kind.

## Solusi Go

```go
package main

// LeetCode #1692: Count Ways to Distribute Candies
// https://leetcode.com/problems/count-ways-to-distribute-candies/
// Difficulty: Hard [Paid]
//
// Given n distinct candies and k bags, count the number of ways to
// distribute all candies into exactly k non-empty bags.
// Two ways are different if at least one candy goes to a different bag.
//
// Approach: Stirling numbers of the second kind.
// dp[i][j] = ways to distribute i candies into j bags.
// Transition: dp[i][j] = j * dp[i-1][j] + dp[i-1][j-1].

import "fmt"

func main() {
	// Example 1
	fmt.Println(waysToDistribute(3, 2))
	// Example 2
	fmt.Println(waysToDistribute(4, 3))
	// Edge: k = n
	fmt.Println(waysToDistribute(3, 3))
	// Edge: k = 1
	fmt.Println(waysToDistribute(5, 1))
}

const W2DMOD = 1000000007

func waysToDistribute(n int, k int) int {
	if k > n || k == 0 {
		return 0
	}
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			dp[i][j] = (j*dp[i-1][j] + dp[i-1][j-1]) % W2DMOD
		}
	}
	return dp[n][k]
}
```
