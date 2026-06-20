# 1770 — Maximum Score From Performing Multiplication Operations

## Deskripsi

**Soal:** [1770. Maximum Score From Performing Multiplication Operations](https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func maximumScore(nums []int, multipliers []int) int`

> **Ide Kunci:** DP[l][i] where l = number of operations done and i = left index used.

## Solusi Go

```go
package main

// LeetCode #1770: Maximum Score from Performing Multiplication Operations
// https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/
// Difficulty: Hard
//
// Approach: DP[l][i] where l = number of operations done and i = left index used.
// Equivalent to: dp[l][left] = max(
//     nums[left] * mult[l] + dp[l+1][left+1],   // take from left
//     nums[right] * mult[l] + dp[l+1][left]      // take from right
// )
// where right = n - 1 - (l - left).
// Optimized: 2D DP is fine since m <= 1000.

import (
	"fmt"
)

func maximumScore(nums []int, multipliers []int) int {
	n := len(nums)
	m := len(multipliers)
	// dp[l][left] = max score using l operations with 'left' left-end picks
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for l := m - 1; l >= 0; l-- {
		for left := l; left >= 0; left-- {
			right := n - 1 - (l - left)
			takeLeft := nums[left]*multipliers[l] + dp[l+1][left+1]
			takeRight := nums[right]*multipliers[l] + dp[l+1][left]
			if takeLeft > takeRight {
				dp[l][left] = takeLeft
			} else {
				dp[l][left] = takeRight
			}
		}
	}

	return dp[0][0]
}

func main() {
	// Example test case
	fmt.Println("nums=[1,2,3],mult=[3,2,1] →", maximumScore([]int{1, 2, 3}, []int{3, 2, 1})) // Expected: 14

	// Additional tests
	fmt.Println("nums=[-5,-3,-3,-2,7,1],mult=[-10,-5,3,4,6] →",
		maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}))

	fmt.Println("nums=[5],mult=[10] →", maximumScore([]int{5}, []int{10}))             // Expected: 50
	fmt.Println("nums=[1,2],mult=[1,2] →", maximumScore([]int{1, 2}, []int{1, 2}))     // Expected: 5 (1*1 + 2*2 = 5)
}
```
