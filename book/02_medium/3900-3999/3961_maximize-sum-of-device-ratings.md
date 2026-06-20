# 3961 — Maximize Sum Of Device Ratings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaximizeSumOfDeviceRatings(units [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(M * N log N)  |  **Ruang:** O(M) where M = devices, N = units per device

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3961: Maximize Sum of Device Ratings
// https://leetcode.com/problems/maximize-sum-of-device-ratings/
// Difficulty: Medium
// Time: O(M * N log N) | Space: O(M) where M = devices, N = units per device
// Approach: Sort each device's units. Rating after optimal transfers =
// second smallest value per device. Sum across devices, then adjust:
// subtract (minSecond - globalMin) since moving global min to the device
// with smallest second-min lowers that device's rating to globalMin.

import (
	"fmt"
	"math"
	"sort"
)

func MaximizeSumOfDeviceRatings(units [][]int) int64 {
	m := len(units)
	if m == 0 {
		return 0
	}
	n := len(units[0])

	if n == 1 {
		var sum int64
		for _, dev := range units {
			sum += int64(dev[0])
		}
		return sum
	}

	globalMin := math.MaxInt32
	minSecond := math.MaxInt32
	var sumSecond int64

	for _, dev := range units {
  // Alokasi slice
		sorted := make([]int, n)
		copy(sorted, dev)
  // Sort O(n log n)
		sort.Ints(sorted)

		if sorted[0] < globalMin {
			globalMin = sorted[0]
		}
		second := sorted[1]
		sumSecond += int64(second)
		if second < minSecond {
			minSecond = second
		}
	}

	return sumSecond - int64(minSecond-globalMin)
}

func main() {
	// Example 1
	fmt.Println(MaximizeSumOfDeviceRatings([][]int{{1, 3}, {2, 2}})) // Expected: 4

	// Example 2
	fmt.Println(MaximizeSumOfDeviceRatings([][]int{{1, 2, 3}, {4, 5, 6}})) // Expected: 6

	// Example 3
	fmt.Println(MaximizeSumOfDeviceRatings([][]int{{5, 5, 5}, {1, 1, 1}})) // Expected: 6
}
```
