# 1444 — Number Of Ways Of Cutting A Pizza

## Deskripsi

**Soal:** [1444. Number Of Ways Of Cutting A Pizza](https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func ways(pizza []string, k int) int`

## Solusi Go

```go
package main

// LeetCode #1444: Number of Ways of Cutting a Pizza
// https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/
// Difficulty: Hard

import "fmt"

const mod1444 = 1_000_000_007

func ways(pizza []string, k int) int {
	rows := len(pizza)
	cols := len(pizza[0])

	// Prefix sum to check if any sub-rectangle has apple
  // Membuat slice 2D untuk DP/tabel
	pref := make([][]int, rows+1)
  // Iterasi seluruh elemen
	for i := range pref {
		pref[i] = make([]int, cols+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			apple := 0
			if pizza[i][j] == 'A' {
				apple = 1
			}
			pref[i+1][j+1] = pref[i][j+1] + pref[i+1][j] - pref[i][j] + apple
		}
	}

	// Helper to check if rectangle has apple
	hasApple := func(r1, c1, r2, c2 int) bool {
		total := pref[r2+1][c2+1] - pref[r1][c2+1] - pref[r2+1][c1] + pref[r1][c1]
		return total > 0
	}

	// dp[r][c][p] = ways to cut pizza from (r,c) to bottom-right with p pieces
  // Membuat slice 2D untuk DP/tabel
	dp := make([][][]int, rows)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([][]int, cols)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			if hasApple(r, c, rows-1, cols-1) {
				dp[r][c][1] = 1
			}
		}
	}

	for p := 2; p <= k; p++ {
		for r := rows - 1; r >= 0; r-- {
			for c := cols - 1; c >= 0; c-- {
				// Horizontal cut
				for nr := r + 1; nr < rows; nr++ {
					if hasApple(r, c, nr-1, cols-1) && dp[nr][c][p-1] > 0 {
						dp[r][c][p] = (dp[r][c][p] + dp[nr][c][p-1]) % mod1444
					}
				}
				// Vertical cut
				for nc := c + 1; nc < cols; nc++ {
					if hasApple(r, c, rows-1, nc-1) && dp[r][nc][p-1] > 0 {
						dp[r][c][p] = (dp[r][c][p] + dp[r][nc][p-1]) % mod1444
					}
				}
			}
		}
	}
	return dp[0][0][k]
}

func main() {
	// Example: ["A..","AAA","..."], k=3 -> 3
	fmt.Println(ways([]string{"A..", "AAA", "..."}, 3))
}
```
