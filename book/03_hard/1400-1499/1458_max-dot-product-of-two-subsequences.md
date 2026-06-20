# 1458 — Max Dot Product Of Two Subsequences

## Deskripsi

**Soal:** [1458. Max Dot Product Of Two Subsequences](https://leetcode.com/problems/max-dot-product-of-two-subsequences/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func maxDotProduct(nums1 []int, nums2 []int) int`

## Solusi Go

```go
package main

// LeetCode #1458: Max Dot Product of Two Subsequences
// https://leetcode.com/problems/max-dot-product-of-two-subsequences/
// Difficulty: Hard

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n2)
	}

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			product := nums1[i] * nums2[j]
			dp[i][j] = product
			if i > 0 && dp[i-1][j] > dp[i][j] {
				dp[i][j] = dp[i-1][j]
			}
			if j > 0 && dp[i][j-1] > dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
			if i > 0 && j > 0 {
				candidate := dp[i-1][j-1]
				if candidate > 0 {
					candidate += product
				} else {
					candidate = product
				}
				if candidate > dp[i][j] {
					dp[i][j] = candidate
				}
			}
		}
	}
	return dp[n1-1][n2-1]
}

func main() {
	// Example: [2,1,-2,5], [3,0,-6] -> 18
	fmt.Println(maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6}))
}
```
