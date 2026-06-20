# 2770 — Maximum Number Of Jumps To Reach The Last Index

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MaximumNumberOfJumpsToReachTheLastIndex(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2770: Maximum Number of Jumps to Reach the Last Index
// https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func MaximumNumberOfJumpsToReachTheLastIndex(nums []int, target int) int {
	n := len(nums)
  // Alokasi slice
	dp := make([]int, n)
  // Range loop
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			if diff < 0 {
				diff = -diff
			}
			if diff <= target && dp[j] != -1 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(MaximumNumberOfJumpsToReachTheLastIndex([]int{1, 3, 6, 4, 1, 2}, 2))
	fmt.Println(MaximumNumberOfJumpsToReachTheLastIndex([]int{1, 3, 6, 4, 1, 2}, 3))
}
```
