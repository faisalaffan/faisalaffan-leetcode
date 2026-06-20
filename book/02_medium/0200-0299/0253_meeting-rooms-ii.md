# 0253 — Meeting Rooms Ii

## Deskripsi

**Soal:** [0253. Meeting Rooms Ii](https://leetcode.com/problems/meeting-rooms-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minMeetingRooms(intervals [][]int) int`

## Solusi Go

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

  // Membuat slice untuk menyimpan hasil
	starts := make([]int, len(intervals))
  // Membuat slice untuk menyimpan hasil
	ends := make([]int, len(intervals))

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	rooms, endIdx := 0, 0

  // Loop standar: indeks 0 sampai n-1
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
