# 2826 — Sorting Three Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SortingThreeGroups(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2826: Sorting Three Groups
// https://leetcode.com/problems/sorting-three-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func SortingThreeGroups(nums []int) int {
	n := len(nums)
	// dp[i][j] = min operations to make first i+1 elements sorted with last element == j+1
  // Alokasi slice
	dp := make([][3]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 0; j < 3; j++ {
			change := 0
			if nums[i-1] != j+1 {
				change = 1
			}
			dp[i][j] = dp[i-1][j] + change
			if j > 0 && dp[i][j-1] < dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[n][2]
}

func main() {
	fmt.Println(SortingThreeGroups([]int{2, 1, 3, 2, 1}))
	fmt.Println(SortingThreeGroups([]int{1, 2, 3, 1, 2, 3}))
}
```
