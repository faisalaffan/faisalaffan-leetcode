# 3409 — Longest Subsequence With Decreasing Adjacent Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestSubsequence(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * maxDiff) Space: O(maxVal * maxDiff)  |  **Ruang:** O(maxVal * maxDiff)


## 💻 Solusi Go

```go
package main

// LeetCode #3409: Longest Subsequence With Decreasing Adjacent Difference
// https://leetcode.com/problems/longest-subsequence-with-decreasing-adjacent-difference/
// Difficulty: Medium
// Time: O(n * maxDiff) Space: O(maxVal * maxDiff)

import (
	"fmt"
	"slices"
)

func longestSubsequence(nums []int) int {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)

  // Matriks 2D
	f := make([][]int, mx+1)
  // Range loop
	for i := range f {
		f[i] = make([]int, maxD+1)
	}

	ans := 0
	for _, x := range nums {
		fx := 1
		for j := maxD; j >= 0; j-- {
			if x-j >= 0 {
				if f[x-j][j]+1 > fx {
					fx = f[x-j][j] + 1
				}
			}
			if x+j <= mx {
				if f[x+j][j]+1 > fx {
					fx = f[x+j][j] + 1
				}
			}
			f[x][j] = fx
			if fx > ans {
				ans = fx
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestSubsequence([]int{16, 6, 3}))              // 3
	fmt.Println(longestSubsequence([]int{6, 5, 3, 4, 2, 1}))      // 4
	fmt.Println(longestSubsequence([]int{10, 20, 30, 40, 50}))    // 5
}
```
