# 0252 — Meeting Rooms

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CanAttendMeetings(intervals [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #252: Meeting Rooms
// https://leetcode.com/problems/meeting-rooms/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func CanAttendMeetings(intervals [][]int) bool {
  // Custom sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CanAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(CanAttendMeetings([][]int{{7, 10}, {2, 4}}))
}
```
