# 1027 — Longest Arithmetic Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestArithSeqLength(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1027: Longest Arithmetic Subsequence
// https://leetcode.com/problems/longest-arithmetic-subsequence/
// Difficulty: Medium
//
// Approach: DP with hash map per index tracking difference -> length
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))    // 4
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10})) // 3
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8})) // 4
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

  // Alokasi slice
	dp := make([]map[int]int, n)
	result := 2

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			length := 2
			if prev, ok := dp[j][diff]; ok {
				length = prev + 1
			}
			dp[i][diff] = length
			if length > result {
				result = length
			}
		}
	}

	return result
}
```
