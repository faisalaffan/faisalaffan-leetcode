# 3323 — Minimize Connected Groups By Inserting Interval

## Deskripsi

**Soal:** [3323. Minimize Connected Groups By Inserting Interval](https://leetcode.com/problems/minimize-connected-groups-by-inserting-interval/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3323: Minimize Connected Groups by Inserting Interval
// https://leetcode.com/problems/minimize-connected-groups-by-inserting-interval/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minConnectedGroups([][]int{{1, 3}, {5, 6}, {8, 10}}, 3)) // 2
	fmt.Println(minConnectedGroups([][]int{{1, 2}, {3, 4}, {5, 6}}, 1)) // 2
	fmt.Println(minConnectedGroups([][]int{{1, 10}}, 5))               // 1
}

func minConnectedGroups(intervals [][]int, k int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Merge overlapping intervals
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}

	n := len(merged)
	if n <= 1 {
		return 1
	}

	// For each group, try to bridge as many groups as possible
	minGroups := n
	j := 0
	for i := 0; i < n; i++ {
		for j < n && merged[j][0]-merged[i][1]-1 <= k {
			j++
		}
		bridged := j - i - 1
		remaining := n - bridged
		if remaining < minGroups {
			minGroups = remaining
		}
	}

	return minGroups
}
```
