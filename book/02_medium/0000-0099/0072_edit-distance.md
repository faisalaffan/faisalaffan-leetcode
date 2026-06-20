# 0072 — Edit Distance

## Deskripsi

**Soal:** [0072. Edit Distance](https://leetcode.com/problems/edit-distance/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func minDistance(word1 string, word2 string) int`

## Solusi Go

```go
package main

// LeetCode #72: Edit Distance
// https://leetcode.com/problems/edit-distance/
// Difficulty: Medium

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test case 1
	fmt.Println(minDistance("horse", "ros")) // 3

	// Test case 2
	fmt.Println(minDistance("intention", "execution")) // 5

	// Test case 3
	fmt.Println(minDistance("", "a")) // 1
}

// Time: O(m*n) | Space: O(m*n)
```
