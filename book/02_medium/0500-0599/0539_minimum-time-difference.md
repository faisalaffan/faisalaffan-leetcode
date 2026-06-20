# 0539 — Minimum Time Difference

## Deskripsi

**Soal:** [0539. Minimum Time Difference](https://leetcode.com/problems/minimum-time-difference/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #539: Minimum Time Difference
// https://leetcode.com/problems/minimum-time-difference/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(FindMinDifference([]string{"23:59", "00:00"}))
	fmt.Println(FindMinDifference([]string{"00:00", "23:59", "00:00"}))
}

func FindMinDifference(timePoints []string) int {
	n := len(timePoints)
  // Membuat slice untuk menyimpan hasil
	minutes := make([]int, n)

	for i, t := range timePoints {
		h, _ := strconv.Atoi(t[:2])
		m, _ := strconv.Atoi(t[3:])
		minutes[i] = h*60 + m
	}

	sort.Ints(minutes)

	minDiff := 24 * 60 // 1440
	for i := 1; i < n; i++ {
		diff := minutes[i] - minutes[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Check circular difference
	circularDiff := 1440 - minutes[n-1] + minutes[0]
	if circularDiff < minDiff {
		minDiff = circularDiff
	}

	return minDiff
}
```
