# 1218 — Longest Arithmetic Subsequence Of Given Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestSubsequence(arr []int, difference int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1218: Longest Arithmetic Subsequence of Given Difference
// https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference/
// Difficulty: Medium

// dp[x] = length of longest arithmetic subsequence ending with value x.
// dp[x] = dp[x-difference] + 1

// Time: O(n)
// Space: O(n)

func longestSubsequence(arr []int, difference int) int {
  // HashMap: O(1) lookup
	dp := make(map[int]int)
	maxLen := 0

	for _, v := range arr {
		prev := v - difference
		if count, exists := dp[prev]; exists {
			dp[v] = count + 1
		} else {
			dp[v] = 1
		}
		if dp[v] > maxLen {
			maxLen = dp[v]
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d (expected: 4)\n", longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Printf("%d (expected: 1)\n", longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Printf("%d (expected: 4)\n", longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}
```
