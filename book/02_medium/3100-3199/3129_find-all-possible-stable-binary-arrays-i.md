# 3129 — Find All Possible Stable Binary Arrays I

## Deskripsi

**Soal:** [3129. Find All Possible Stable Binary Arrays I](https://leetcode.com/problems/find-all-possible-stable-binary-arrays-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(zero * one * limit)  
**Kompleksitas Ruang:** O(zero * one * 2)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func numberOfStableArrays(zero int, one int, limit int) int`

## Solusi Go

```go
package main

// LeetCode #3129: Find All Possible Stable Binary Arrays I
// https://leetcode.com/problems/find-all-possible-stable-binary-arrays-i/
// Difficulty: Medium
// Time: O(zero * one * limit) | Space: O(zero * one * 2)

import "fmt"

func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007
  // Membuat slice 2D untuk DP/tabel
	dp := make([][][2]int, zero+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}

	for i := 0; i <= zero; i++ {
		for j := 0; j <= one; j++ {
			if i == 0 && j == 0 {
				dp[i][j][0] = 1
				dp[i][j][1] = 1
				continue
			}
			if i > 0 {
				dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
				if i > limit {
					dp[i][j][0] -= dp[i-limit-1][j][1]
				}
				dp[i][j][0] %= mod
				if dp[i][j][0] < 0 {
					dp[i][j][0] += mod
				}
			}
			if j > 0 {
				dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]
				if j > limit {
					dp[i][j][1] -= dp[i][j-limit-1][0]
				}
				dp[i][j][1] %= mod
				if dp[i][j][1] < 0 {
					dp[i][j][1] += mod
				}
			}
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

func main() {
	fmt.Println(numberOfStableArrays(1, 1, 2)) // Expected: 2
	fmt.Println(numberOfStableArrays(1, 2, 1)) // Expected: 1
	fmt.Println(numberOfStableArrays(3, 1, 1)) // Expected: 2
}
```
