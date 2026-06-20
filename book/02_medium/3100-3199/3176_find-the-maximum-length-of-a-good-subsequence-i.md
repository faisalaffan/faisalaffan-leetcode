# 3176 — Find The Maximum Length Of A Good Subsequence I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumLength(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^2 * k)  |  **Ruang:** O(n * k)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3176: Find the Maximum Length of a Good Subsequence I
// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-i/
// Difficulty: Medium
// Time: O(n^2 * k) | Space: O(n * k)

import "fmt"

func maximumLength(nums []int, k int) int {
	n := len(nums)
  // Matriks 2D
	dp := make([][]int, n)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for p := 0; p < i; p++ {
				if nums[i] == nums[p] {
					if dp[p][j]+1 > dp[i][j] {
						dp[i][j] = dp[p][j] + 1
					}
				} else if j > 0 {
					if dp[p][j-1]+1 > dp[i][j] {
						dp[i][j] = dp[p][j-1] + 1
					}
				}
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 3}, 2)) // Expected: 4
	fmt.Println(maximumLength([]int{1, 2, 3, 4}, 0))     // Expected: 1
}
```
