# 0945 — Minimum Increment To Make Array Unique

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minIncrementForUnique(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1) ignoring sort space

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #945: Minimum Increment to Make Array Unique
// https://leetcode.com/problems/minimum-increment-to-make-array-unique/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1) ignoring sort space
func minIncrementForUnique(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	moves := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			moves += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return moves
}

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
```
