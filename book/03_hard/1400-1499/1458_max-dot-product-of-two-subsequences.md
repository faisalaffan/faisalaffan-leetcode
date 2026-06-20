# 1458 — Max Dot Product Of Two Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxDotProduct(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1458: Max Dot Product of Two Subsequences
// https://leetcode.com/problems/max-dot-product-of-two-subsequences/
// Difficulty: Hard

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
  // Matriks 2D
	dp := make([][]int, n1)
  // Range loop
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
