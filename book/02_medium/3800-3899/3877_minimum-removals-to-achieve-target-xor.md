# 3877 — Minimum Removals To Achieve Target Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumRemovalsToAchieveTargetXor(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(N * 2^M)  |  **Ruang:** O(2^M) where M = max bit length (14)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3877: Minimum Removals to Achieve Target XOR
// https://leetcode.com/problems/minimum-removals-to-achieve-target-xor/
// Difficulty: Medium
// Time: O(N * 2^M) | Space: O(2^M) where M = max bit length (14)
// Approach: DP tracking max selectable elements to achieve each XOR value.
// Answer = len(nums) - maxElementsForTarget (or -1 if unreachable).

import "fmt"

func MinimumRemovalsToAchieveTargetXor(nums []int, target int) int {
	maxXor := 1
	for _, v := range nums {
		for maxXor <= v {
			maxXor <<= 1
		}
	}
	if maxXor <= target {
		for maxXor <= target {
			maxXor <<= 1
		}
	}

	// dp[x] = max elements selectable to achieve XOR x
  // Alokasi slice
	dp := make([]int, maxXor)
  // Range loop
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for _, v := range nums {
  // Alokasi slice
		ndp := make([]int, maxXor)
		copy(ndp, dp)
		for x := 0; x < maxXor; x++ {
			if dp[x] >= 0 {
				nx := x ^ v
				if dp[x]+1 > ndp[nx] {
					ndp[nx] = dp[x] + 1
				}
			}
		}
		dp = ndp
	}

	if dp[target] < 0 {
		return -1
	}
	return len(nums) - dp[target]
}

func main() {
	// Example 1
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{1, 2, 3}, 2)) // Expected: 1

	// Example 2
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{2, 4}, 1)) // Expected: -1

	// Example 3
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{7}, 7)) // Expected: 0
}
```
