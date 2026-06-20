# 2655 — Find Maximal Uncovered Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findMaximalUncoveredRanges(n int, ranges [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2655: Find Maximal Uncovered Ranges
// https://leetcode.com/problems/find-maximal-uncovered-ranges/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findMaximalUncoveredRanges(n int, ranges [][]int) [][]int {
  // Custom sort
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	ans := [][]int{}
	prevEnd := -1

	for _, r := range ranges {
		start, end := r[0], r[1]
		if start > prevEnd+1 {
			ans = append(ans, []int{prevEnd + 1, start - 1})
		}
		if end > prevEnd {
			prevEnd = end
		}
	}

	if prevEnd < n-1 {
		ans = append(ans, []int{prevEnd + 1, n - 1})
	}

	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMaximalUncoveredRanges(10, [][]int{{0, 2}, {5, 7}}))
	// Expected: [[3,4],[8,9]]

	// Test case 2
	fmt.Println("Test 2:", findMaximalUncoveredRanges(5, [][]int{{0, 4}}))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", findMaximalUncoveredRanges(5, [][]int{}))

	// Expected: [[0,4]]
}
```
