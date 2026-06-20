# 2406 — Divide Intervals Into Minimum Number Of Groups

## Deskripsi

**Soal:** [2406. Divide Intervals Into Minimum Number Of Groups](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2406: Divide Intervals Into Minimum Number of Groups
// https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Sweep line: count concurrent intervals, answer is max concurrency.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}})) // 3
	fmt.Println(minGroups([][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}))        // 1
}

func minGroups(intervals [][]int) int {
  // Membuat slice untuk menyimpan hasil
	events := make([][2]int, 0, len(intervals)*2)
	for _, iv := range intervals {
		events = append(events, [2]int{iv[0], 1})   // start
		events = append(events, [2]int{iv[1] + 1, -1}) // end (inclusive)
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	cur, ans := 0, 0
	for _, e := range events {
		cur += e[1]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```
