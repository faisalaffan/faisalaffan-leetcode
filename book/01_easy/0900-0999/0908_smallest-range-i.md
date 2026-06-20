# 0908 — Smallest Range I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func smallestRangeI(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #908: Smallest Range I
// https://leetcode.com/problems/smallest-range-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))          // 0
	fmt.Println(smallestRangeI([]int{0, 10}, 2))      // 6
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))    // 0
}

// smallestRangeI returns the smallest possible range after modifying each element by at most k.
// Time: O(n). Space: O(1).
func smallestRangeI(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	diff := (maxVal - k) - (minVal + k)
	if diff < 0 {
		return 0
	}
	return diff
}
```
