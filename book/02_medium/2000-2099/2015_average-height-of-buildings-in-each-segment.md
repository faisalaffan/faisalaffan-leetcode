# 2015 — Average Height Of Buildings In Each Segment

## Deskripsi

**Soal:** [2015. Average Height Of Buildings In Each Segment](https://leetcode.com/problems/average-height-of-buildings-in-each-segment/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2015: Average Height of Buildings in Each Segment
// https://leetcode.com/problems/average-height-of-buildings-in-each-segment/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 4, 2}, {3, 9, 4}}))
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 3, 2}, {2, 5, 3}, {2, 8, 3}}))
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 2, 1}, {5, 6, 1}}))
}

// Time: O(n log n), Space: O(n)
func AverageHeightOfBuildingsInEachSegment(buildings [][]int) [][]int {
	type event struct {
		pos    int
		height int
		change int
	}

  // Membuat slice untuk menyimpan hasil
	events := make([]event, 0)
	for _, b := range buildings {
		start, end, height := b[0], b[1], b[2]
		events = append(events, event{start, height, 1})
		events = append(events, event{end, height, -1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].pos != events[j].pos {
			return events[i].pos < events[j].pos
		}
		return events[i].height < events[j].height
	})

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0)
	totalHeight := 0
	buildingCount := 0
	prevPos := -1

	for _, e := range events {
		if prevPos != -1 && prevPos < e.pos && buildingCount > 0 {
			avg := totalHeight / buildingCount
			n := len(result)
			if n > 0 && result[n-1][2] == avg && result[n-1][1] == prevPos {
				result[n-1][1] = e.pos
			} else {
				result = append(result, []int{prevPos, e.pos, avg})
			}
		}

		totalHeight += e.height * e.change
		buildingCount += e.change
		prevPos = e.pos
	}

	return result
}
```
