# 3414 — Maximum Score Of Non Overlapping Intervals

## Deskripsi

**Soal:** [3414. Maximum Score Of Non Overlapping Intervals](https://leetcode.com/problems/maximum-score-of-non-overlapping-intervals/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner), Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #3414: Maximum Score of Non-overlapping Intervals
// https://leetcode.com/problems/maximum-score-of-non-overlapping-intervals/
// Difficulty: Hard
//
// Weighted interval scheduling. Sort by end time, binary search for
// previous non-overlapping interval, DP.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumScoreOfNonOverlappingIntervals([][]int{{1, 3, 2}, {2, 5, 3}, {4, 6, 1}}))
}

func MaximumScoreOfNonOverlappingIntervals(intervals [][]int) int {
	n := len(intervals)
	type interval struct{ start, end, score int }
  // Membuat slice untuk menyimpan hasil
	ivs := make([]interval, n)
	for i, v := range intervals {
		ivs[i] = interval{v[0], v[1], v[2]}
	}
	sort.Slice(ivs, func(i, j int) bool {
		return ivs[i].end < ivs[j].end
	})

  // Membuat slice untuk menyimpan hasil
	ends := make([]int, n)
	for i, v := range ivs {
		ends[i] = v.end
	}

  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n)
	dp[0] = ivs[0].score
	for i := 1; i < n; i++ {
		// Binary search for last interval ending <= ivs[i].start
		prev := sort.SearchInts(ends, ivs[i].start+1) - 1
		best := ivs[i].score
		if prev >= 0 {
			best += dp[prev]
		}
		if dp[i-1] > best {
			best = dp[i-1]
		}
		dp[i] = best
	}
	return dp[n-1]
}
```
