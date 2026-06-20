# 2895 — Minimum Processing Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumProcessingTime(processorTime []int, taskTime []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2895: Minimum Processing Time
// https://leetcode.com/problems/minimum-processing-time/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MinimumProcessingTime(processorTime []int, taskTime []int) int {
  // Sort O(n log n)
	sort.Ints(processorTime)
  // Sort O(n log n)
	sort.Ints(taskTime)

	// Each processor handles 4 tasks
	n := len(processorTime)
	maxTime := 0

	for i := 0; i < n; i++ {
		// Assign 4 largest remaining tasks to the slowest processor
		// Processor runs concurrently, so time = processorTime[i] + max of its 4 tasks
		taskIdx := len(taskTime) - 1 - i*4
		time := processorTime[i] + taskTime[taskIdx]
		if time > maxTime {
			maxTime = time
		}
	}

	return maxTime
}

func main() {
	fmt.Println(MinimumProcessingTime([]int{8, 10}, []int{2, 2, 3, 1, 8, 7, 4, 5}))
	fmt.Println(MinimumProcessingTime([]int{10, 20}, []int{2, 3, 1, 2, 5, 8, 4, 3}))
}
```
