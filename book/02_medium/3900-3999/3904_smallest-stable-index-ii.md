# 3904 — Smallest Stable Index Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SmallestStableIndexIi(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3904: Smallest Stable Index II
// https://leetcode.com/problems/smallest-stable-index-ii/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Precompute suffix minimum, iterate prefix maximum.
// Check if max(nums[0..i]) - min(nums[i..n-1]) <= k.

import "fmt"

func SmallestStableIndexIi(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(nums[i], suffixMin[i+1])
	}

	prefixMax := 0
	for i := 0; i < n; i++ {
		if nums[i] > prefixMax {
			prefixMax = nums[i]
		}
		if prefixMax-suffixMin[i] <= k {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(SmallestStableIndexIi([]int{5, 0, 1, 4}, 3)) // Expected: 3

	// Example 2
	fmt.Println(SmallestStableIndexIi([]int{3, 2, 1}, 1)) // Expected: -1

	// Example 3
	fmt.Println(SmallestStableIndexIi([]int{0}, 0)) // Expected: 0
}
```
