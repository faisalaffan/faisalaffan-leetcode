# 1509 — Minimum Difference Between Largest And Smallest Value In Three Moves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinDifference(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(N log N), Space: O(1) if ignoring sort space  |  **Ruang:** O(1) if ignoring sort space

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1509: Minimum Difference Between Largest and Smallest Value in Three Moves
// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinDifference([]int{5, 3, 2, 4}))
	fmt.Println(MinDifference([]int{1, 5, 0, 10, 14}))
	fmt.Println(MinDifference([]int{3, 100, 20}))
}

func MinDifference(nums []int) int {
	// Time: O(N log N), Space: O(1) if ignoring sort space
	if len(nums) <= 4 {
		return 0
	}

  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)

	// After 3 moves, we can change up to 3 values.
	// The minimum difference will be between some combination
	// of removing 0-3 from left and 3-0 from right.
	minDiff := nums[n-1] - nums[0]
	for i := 0; i <= 3; i++ {
		diff := nums[n-1-(3-i)] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
```
