# 3727 — Maximum Alternating Sum Of Squares

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumAlternatingSumOfSquares(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3727: Maximum Alternating Sum of Squares
// https://leetcode.com/problems/maximum-alternating-sum-of-squares/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumAlternatingSumOfSquares(nums []int) int64 {
  // Custom sort
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	var ans int64
	for i, v := range nums {
		sq := int64(v) * int64(v)
		if i%2 == 0 {
			ans += sq
		} else {
			ans -= sq
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(maximumAlternatingSumOfSquares([]int{1, 2, 3}))
	fmt.Println(maximumAlternatingSumOfSquares([]int{1, -1, 2, -2, 3, -3}))
	fmt.Println(maximumAlternatingSumOfSquares([]int{0, 0, 0}))
}
```
