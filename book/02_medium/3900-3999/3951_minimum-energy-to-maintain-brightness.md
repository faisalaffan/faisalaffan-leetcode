# 3951 — Minimum Energy To Maintain Brightness

## Deskripsi

**Soal:** [3951. Minimum Energy To Maintain Brightness](https://leetcode.com/problems/minimum-energy-to-maintain-brightness/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N log N)  
**Kompleksitas Ruang:** O(N) where N = len(intervals)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumEnergyToMaintainBrightness(n int, brightness int, intervals [][]int) int64`

> **Ide Kunci:** Each bulb illuminates 3 positions (self + adjacents). Need

## Solusi Go

```go
package main

// LeetCode #3951: Minimum Energy to Maintain Brightness
// https://leetcode.com/problems/minimum-energy-to-maintain-brightness/
// Difficulty: Medium
// Time: O(N log N) | Space: O(N) where N = len(intervals)
// Approach: Each bulb illuminates 3 positions (self + adjacents). Need
// ceil(brightness/3) bulbs. Merge overlapping intervals. For each merged
// interval, energy += bulbs * intervalLength. Return total energy.

import (
	"fmt"
	"sort"
)

func MinimumEnergyToMaintainBrightness(n int, brightness int, intervals [][]int) int64 {
	bulbs := int64((brightness + 2) / 3)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var total int64
	i := 0
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end+1 {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
			j++
		}
		total += int64(end - start + 1)
		i = j
	}

	return bulbs * total
}

func main() {
	// Example 1
	fmt.Println(MinimumEnergyToMaintainBrightness(5, 5, [][]int{{6, 12}})) // Expected: 14

	// Example 2
	fmt.Println(MinimumEnergyToMaintainBrightness(2, 1, [][]int{{0, 0}, {2, 2}})) // Expected: 2

	// Example 3
	fmt.Println(MinimumEnergyToMaintainBrightness(4, 2, [][]int{{1, 3}, {2, 4}})) // Expected: 4
}
```
