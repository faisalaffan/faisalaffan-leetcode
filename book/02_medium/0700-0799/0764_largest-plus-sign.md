# 0764 — Largest Plus Sign

## Deskripsi

**Soal:** [0764. Largest Plus Sign](https://leetcode.com/problems/largest-plus-sign/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #764: Largest Plus Sign
// https://leetcode.com/problems/largest-plus-sign/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
	fmt.Println(orderOfLargestPlusSign(1, [][]int{{0, 0}}))
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
  // Membuat map untuk pencarian O(1): key → value
	mineSet := make(map[int]bool)
	for _, m := range mines {
		mineSet[m[0]*n+m[1]] = true
	}

  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
	}

	best := 0

	for r := 0; r < n; r++ {
		count := 0
		for c := 0; c < n; c++ {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = count
		}

		count = 0
		for c := n - 1; c >= 0; c-- {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = min(dp[r][c], count)
		}
	}

	for c := 0; c < n; c++ {
		count := 0
		for r := 0; r < n; r++ {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = min(dp[r][c], count)
		}

		count = 0
		for r := n - 1; r >= 0; r-- {
			if mineSet[r*n+c] {
				count = 0
			} else {
				count++
			}
			dp[r][c] = min(dp[r][c], count)
			if dp[r][c] > best {
				best = dp[r][c]
			}
		}
	}

	return best
}
```
