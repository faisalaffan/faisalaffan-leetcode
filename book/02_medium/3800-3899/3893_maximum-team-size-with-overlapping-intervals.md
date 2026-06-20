# 3893 — Maximum Team Size With Overlapping Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumTeamSizeWithOverlappingIntervals(startTime []int, endTime []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting

**Waktu:** O(N log N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3893: Maximum Team Size with Overlapping Intervals
// https://leetcode.com/problems/maximum-team-size-with-overlapping-intervals/
// Difficulty: Medium [Paid]
// Time: O(N log N) | Space: O(N)
// Approach: For each employee, count overlapping intervals using binary search
// on sorted start and end times.

import (
	"fmt"
	"sort"
)

func MaximumTeamSizeWithOverlappingIntervals(startTime []int, endTime []int) int {
	n := len(startTime)
  // Alokasi slice
	st := make([]int, n)
  // Alokasi slice
	et := make([]int, n)
	copy(st, startTime)
	copy(et, endTime)
  // Sort O(n log n)
	sort.Ints(st)
  // Sort O(n log n)
	sort.Ints(et)

	ans := 0
	for i := 0; i < n; i++ {
		start := startTime[i]
		end := endTime[i]

		// Count employees whose start <= end
		startsBeforeEnd := sort.SearchInts(st, end+1)
		// Count employees whose end < start
		endsBeforeStart := sort.SearchInts(et, start)

		overlap := startsBeforeEnd - endsBeforeStart
		if overlap > ans {
			ans = overlap
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{1, 2, 3}, []int{4, 5, 6})) // Expected: 3

	// Example 2
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{2, 5, 8}, []int{3, 7, 9})) // Expected: 1

	// Example 3
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{3, 4, 6}, []int{8, 5, 7})) // Expected: 3
}
```
