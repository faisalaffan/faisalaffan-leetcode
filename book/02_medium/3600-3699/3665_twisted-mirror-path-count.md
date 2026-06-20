# 3665 — Twisted Mirror Path Count

## Deskripsi

**Soal:** [3665. Twisted Mirror Path Count](https://leetcode.com/problems/twisted-mirror-path-count/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func twistedMirrorPathCount(grid [][]int) int`

## Solusi Go

```go
package main

// LeetCode #3665: Twisted Mirror Path Count
// https://leetcode.com/problems/twisted-mirror-path-count/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func twistedMirrorPathCount(grid [][]int) int {
	mod := int(1e9 + 7)
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				continue
			}
			cur := dp[i][j]

			// Try moving right to (i, j+1)
			if j+1 < n {
				if grid[i][j+1] == 1 {
					// Mirror at (i, j+1): reflect down to (i+1, j+1)
					if i+1 < m {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i][j+1] = (dp[i][j+1] + cur) % mod
				}
			}

			// Try moving down to (i+1, j)
			if i+1 < m {
				if grid[i+1][j] == 1 {
					// Mirror at (i+1, j): reflect right to (i+1, j+1)
					if j+1 < n {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i+1][j] = (dp[i+1][j] + cur) % mod
				}
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 0}, {0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1}, {1, 0}}))
}
```
