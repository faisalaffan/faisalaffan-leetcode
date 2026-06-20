# 3961 — Maximize Sum Of Device Ratings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeSumOfDeviceRatings(units [][]int) int64
```

> **💡 Hint:** Sort each device's units. Rating after optimal transfers =

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(M * N log N)  
**Kompleksitas Ruang:** O(M) where M = devices, N = units per device

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Alokasi slice integer
		sorted := make([]int, n)
		copy(sorted, dev)
  // Urutkan secara ascending — O(n log n)
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
