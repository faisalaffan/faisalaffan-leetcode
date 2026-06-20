# 1229 — Meeting Scheduler

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n + m log m)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1229: Meeting Scheduler
// https://leetcode.com/problems/meeting-scheduler/
// Difficulty: Medium [Paid]

// Find earliest time slot of given duration that works for both.
// Two pointers approach after sorting slots by start time.

// Time: O(n log n + m log m)
// Space: O(1)

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
  // Custom sort
	sort.Slice(slots1, func(i, j int) bool { return slots1[i][0] < slots1[j][0] })
  // Custom sort
	sort.Slice(slots2, func(i, j int) bool { return slots2[i][0] < slots2[j][0] })

	i, j := 0, 0
	for i < len(slots1) && j < len(slots2) {
		start := max(slots1[i][0], slots2[j][0])
		end := min(slots1[i][1], slots2[j][1])
		if end-start >= duration {
			return []int{start, start + duration}
		}
		if slots1[i][1] < slots2[j][1] {
			i++
		} else {
			j++
		}
	}
	return []int{}
}

func main() {
	fmt.Printf("%v (expected: [60 68])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}},
			[][]int{{0, 15}, {60, 70}}, 8))

	fmt.Printf("%v (expected: [])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}},
			[][]int{{0, 15}, {55, 58}}, 5))
}
```
