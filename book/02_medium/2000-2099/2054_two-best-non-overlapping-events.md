# 2054 — Two Best Non Overlapping Events

## Deskripsi

**Soal:** [2054. Two Best Non Overlapping Events](https://leetcode.com/problems/two-best-non-overlapping-events/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maxTwoEvents(events [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2054: Two Best Non-Overlapping Events
// https://leetcode.com/problems/two-best-non-overlapping-events/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxTwoEvents(events [][]int) int {
	// Sort by end time
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	// bestUpTo[i] = max value using events[0..i] (single event, non-overlapping)
  // Membuat slice untuk menyimpan hasil
	bestUpTo := make([]int, n)
	bestUpTo[0] = events[0][2]
	for i := 1; i < n; i++ {
		if events[i][2] > bestUpTo[i-1] {
			bestUpTo[i] = events[i][2]
		} else {
			bestUpTo[i] = bestUpTo[i-1]
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		// Take event i
		result = max(result, events[i][2])
		// Find last event that ends before this event starts
		lo, hi := 0, i-1
		best := -1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if events[mid][1] < events[i][0] {
				best = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if best != -1 {
			result = max(result, events[i][2]+bestUpTo[best])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", maxTwoEvents([][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}))
	// Expected: 8
}
```
