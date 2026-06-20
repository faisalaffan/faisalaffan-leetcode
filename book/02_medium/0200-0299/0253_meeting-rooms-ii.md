# 0253 — Meeting Rooms Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minMeetingRooms(intervals [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #253: Meeting Rooms II
// https://leetcode.com/problems/meeting-rooms-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

  // Alokasi slice
	starts := make([]int, len(intervals))
  // Alokasi slice
	ends := make([]int, len(intervals))

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

  // Sort O(n log n)
	sort.Ints(starts)
  // Sort O(n log n)
	sort.Ints(ends)

	rooms, endIdx := 0, 0

  // Linear scan O(n)
	for i := 0; i < len(starts); i++ {
		if starts[i] < ends[endIdx] {
			rooms++
		} else {
			endIdx++
		}
	}

	return rooms
}

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(minMeetingRooms([][]int{{7, 10}, {2, 4}}))
	fmt.Println(minMeetingRooms([][]int{{0, 5}, {5, 10}, {10, 15}}))
}
```
