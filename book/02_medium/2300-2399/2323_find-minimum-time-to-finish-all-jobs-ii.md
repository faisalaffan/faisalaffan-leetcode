# 2323 — Find Minimum Time To Finish All Jobs Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumTime(jobs []int, workers []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2323: Find Minimum Time to Finish All Jobs II
// https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumTime(jobs []int, workers []int) int {
  // Sort O(n log n)
	sort.Ints(jobs)
  // Sort O(n log n)
	sort.Ints(workers)
	maxDays := 0

  // Linear scan O(n)
	for i := 0; i < len(jobs); i++ {
		days := (jobs[i] + workers[i] - 1) / workers[i] // ceil division
		if days > maxDays {
			maxDays = days
		}
	}
	return maxDays
}

func main() {
	// Test case 1
	fmt.Println(minimumTime([]int{5, 2, 4}, []int{1, 7, 5}))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumTime([]int{3, 18, 30}, []int{3, 15, 5}))
	// Expected: 6
}
```
