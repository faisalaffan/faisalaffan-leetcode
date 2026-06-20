# 0436 — Find Right Interval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findRightInterval(intervals [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #436: Find Right Interval
// https://leetcode.com/problems/find-right-interval/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	// Create array of (start, index) pairs
  // Alokasi slice
	starts := make([][2]int, n)
	for i, iv := range intervals {
		starts[i] = [2]int{iv[0], i}
	}
  // Custom sort
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})

  // Alokasi slice
	result := make([]int, n)
	for i, iv := range intervals {
		target := iv[1]
		// Binary search
		idx := sort.Search(n, func(j int) bool {
			return starts[j][0] >= target
		})
		if idx < n {
			result[i] = starts[idx][1]
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findRightInterval([][]int{{1, 2}}))
	// Expected: [-1]

	// Test case 2
	fmt.Println("Test 2:", findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
	// Expected: [-1, 0, 1]

	// Test case 3
	fmt.Println("Test 3:", findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
	// Expected: [-1, 2, -1]
}
```
