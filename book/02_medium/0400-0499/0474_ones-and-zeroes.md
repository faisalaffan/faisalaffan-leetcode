# 0474 — Ones And Zeroes

## Deskripsi

**Soal:** [0474. Ones And Zeroes](https://leetcode.com/problems/ones-and-zeroes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n * len(strs))  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #474: Ones and Zeroes
// https://leetcode.com/problems/ones-and-zeroes/
// Difficulty: Medium
// Time: O(m * n * len(strs))
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(OnesAndZeroes([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(OnesAndZeroes([]string{"10", "0", "1"}, 1, 1))
}

func OnesAndZeroes(strs []string, m int, n int) int {
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := countBits(s)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}

	return dp[m][n]
}

func countBits(s string) (int, int) {
	zeros, ones := 0, 0
	for _, c := range s {
		if c == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}
```
