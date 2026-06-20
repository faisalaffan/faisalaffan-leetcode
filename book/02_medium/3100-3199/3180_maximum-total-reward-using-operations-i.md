# 3180 — Maximum Total Reward Using Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxTotalReward(rewardValues []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * maxVal)  
**Kompleksitas Ruang:** O(maxVal)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Urutkan secara ascending — O(n log n)
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
