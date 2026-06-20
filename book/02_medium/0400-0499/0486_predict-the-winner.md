# 0486 — Predict The Winner

## Deskripsi

**Soal:** [0486. Predict The Winner](https://leetcode.com/problems/predict-the-winner/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #486: Predict the Winner
// https://leetcode.com/problems/predict-the-winner/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			// Max of (pick left) or (pick right), minus opponent's optimal play
			left := nums[i] - dp[i+1][j]
			right := nums[j] - dp[i][j-1]
			if left > right {
				dp[i][j] = left
			} else {
				dp[i][j] = right
			}
		}
	}

	return dp[0][n-1] >= 0
}
```
