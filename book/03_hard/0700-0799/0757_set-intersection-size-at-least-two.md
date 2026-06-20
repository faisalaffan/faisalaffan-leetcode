# 0757 — Set Intersection Size At Least Two

## Deskripsi

**Soal:** [0757. Set Intersection Size At Least Two](https://leetcode.com/problems/set-intersection-size-at-least-two/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal)

**Fungsi Solusi:** `func intersectionSizeTwo(intervals [][]int) int`

## Solusi Go

```go
package main

// LeetCode #757: Set Intersection Size At Least Two
// https://leetcode.com/problems/set-intersection-size-at-least-two/
// Difficulty: Hard
//
// Algorithm: Greedy
// 1. Sort intervals by end point ascending, then by start point descending
// 2. Maintain last two chosen points
// 3. For each interval, determine how many of the chosen points are within it
// 4. Add new points greedily (from the end) when needed

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort by end ascending, then by start descending
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})

	// Chosen points (we'll maintain at most 2 per interval at the end)
  // Membuat slice untuk menyimpan hasil
	chosen := make([]int, 0, len(intervals)*2)
	chosen = append(chosen, intervals[0][1]-1, intervals[0][1])

	for _, interval := range intervals[1:] {
		start, end := interval[0], interval[1]

		// Count how many of the last two chosen points are in this interval
		count := 0
		if len(chosen) >= 1 && chosen[len(chosen)-1] >= start {
			count++
		}
		if len(chosen) >= 2 && chosen[len(chosen)-2] >= start {
			count++
		}

		if count < 2 {
			// Need to add points
			// Add from the end: add points starting from 'end' going backwards
			need := 2 - count
			toAdd := end
			for need > 0 {
				// Don't add already chosen points
				if len(chosen) == 0 || toAdd != chosen[len(chosen)-1] {
					chosen = append(chosen, toAdd)
					need--
				}
				toAdd--
			}
		}
	}

	return len(chosen)
}

func main() {
	// Example from problem
	intervals1 := [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}
	result1 := intersectionSizeTwo(intervals1)
	fmt.Printf("Input: %v\nOutput: %d (expected: 3)\n\n", intervals1, result1)

	// Test case 2
	intervals2 := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}
	result2 := intersectionSizeTwo(intervals2)
	fmt.Printf("Input: %v\nOutput: %d\n\n", intervals2, result2)

	// Test case 3: single interval
	intervals3 := [][]int{{1, 5}}
	result3 := intersectionSizeTwo(intervals3)
	fmt.Printf("Input: %v\nOutput: %d (expected: 2)\n\n", intervals3, result3)

	// Test case 4: two overlapping intervals
	intervals4 := [][]int{{1, 2}, {2, 3}}
	result4 := intersectionSizeTwo(intervals4)
	fmt.Printf("Input: %v\nOutput: %d\n\n", intervals4, result4)

	// Test case 5: all overlapping
	intervals5 := [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}}
	result5 := intersectionSizeTwo(intervals5)
	fmt.Printf("Input: %v\nOutput: %d (expected: 2)\n\n", intervals5, result5)

	// Test case 6: chain
	intervals6 := [][]int{{1, 3}, {2, 4}, {4, 6}, {5, 7}}
	result6 := intersectionSizeTwo(intervals6)
	fmt.Printf("Input: %v\nOutput: %d\n\n", intervals6, result6)

	// Test case 7: exactly 2
	intervals7 := [][]int{{2, 10}, {3, 7}, {3, 15}, {4, 11}, {6, 12}, {7, 9}}
	result7 := intersectionSizeTwo(intervals7)
	fmt.Printf("Input: %v\nOutput: %d\n\n", intervals7, result7)

	// Test case 8: leetcode example 2
	intervals8 := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}
	result8 := intersectionSizeTwo(intervals8)
	fmt.Printf("Input: %v\nOutput: %d\n", intervals8, result8)
}
```
