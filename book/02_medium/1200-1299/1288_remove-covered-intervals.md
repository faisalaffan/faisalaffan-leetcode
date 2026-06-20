# 1288 — Remove Covered Intervals

## Deskripsi

**Soal:** [1288. Remove Covered Intervals](https://leetcode.com/problems/remove-covered-intervals/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func removeCoveredIntervals(intervals [][]int) int`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1288: Remove Covered Intervals
// https://leetcode.com/problems/remove-covered-intervals/
// Difficulty: Medium

// Sort by start ascending, end descending. Track current max end.
// If end <= maxEnd, interval is covered.

// Time: O(n log n)
// Space: O(1)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	count := 0
	maxEnd := 0

	for _, inv := range intervals {
		if inv[1] > maxEnd {
			count++
			maxEnd = inv[1]
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}))

	fmt.Printf("%d (expected: 1)\n",
		removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))

	fmt.Printf("%d (expected: 2)\n",
		removeCoveredIntervals([][]int{{0, 10}, {5, 12}}))
}
```
