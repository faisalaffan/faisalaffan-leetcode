# 1621 — Number Of Sets Of K Non Overlapping Line Segments

## Deskripsi

**Soal:** [1621. Number Of Sets Of K Non Overlapping Line Segments](https://leetcode.com/problems/number-of-sets-of-k-non-overlapping-line-segments/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N*K), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #1621: Number of Sets of K Non-Overlapping Line Segments
// https://leetcode.com/problems/number-of-sets-of-k-non-overlapping-line-segments/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfSets(4, 2))
	fmt.Println(NumberOfSets(3, 1))
	fmt.Println(NumberOfSets(30, 7))
}

func NumberOfSets(n int, k int) int {
	// Time: O(N*K), Space: O(N)
	const mod = 1_000_000_007

	// dp[i][j][0] = ways using i points with j segments, not ending at i
	// dp[i][j][1] = ways using i points with j segments, ending at i

  // Membuat slice 2D untuk DP/tabel
	dp := make([][][2]int, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}

	// Base: with 0 segments, there's 1 way
	for i := 1; i <= n; i++ {
		dp[i][0][0] = 1
	}

	for j := 1; j <= k; j++ {
		for i := 2; i <= n; i++ {
			// Not ending at i: just carry forward
			dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod

			// Ending at i: either extend previous ending at i-1, or start new ending at i
			dp[i][j][1] = (dp[i-1][j-1][0] + dp[i-1][j][1]) % mod
		}
	}

	return (dp[n][k][0] + dp[n][k][1]) % mod
}
```
