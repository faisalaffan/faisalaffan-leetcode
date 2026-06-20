# 3180 — Maximum Total Reward Using Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxTotalReward(rewardValues []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Sorting

**Waktu:** O(n * maxVal)  |  **Ruang:** O(maxVal)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3180: Maximum Total Reward Using Operations I
// https://leetcode.com/problems/maximum-total-reward-using-operations-i/
// Difficulty: Medium
// Time: O(n * maxVal) | Space: O(maxVal)

import (
	"fmt"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
  // Sort O(n log n)
	sort.Ints(rewardValues)
	maxVal := rewardValues[len(rewardValues)-1]
	size := 2 * maxVal
	dp := make([]bool, size)
	dp[0] = true

	for _, v := range rewardValues {
		for x := size - 1 - v; x >= 0; x-- {
			if dp[x] && v > x {
				dp[x+v] = true
			}
		}
	}

	for x := size - 1; x >= 0; x-- {
		if dp[x] {
			return x
		}
	}
	return 0
}

func main() {
	fmt.Println(maxTotalReward([]int{1, 1, 3, 3})) // Expected: 4
	fmt.Println(maxTotalReward([]int{1, 6, 4, 3, 2})) // Expected: 11
}
```
