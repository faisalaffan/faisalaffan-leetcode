# 0475 — Heaters

## Deskripsi

**Soal:** [0475. Heaters](https://leetcode.com/problems/heaters/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n + m log n) where n = len(heaters), m = len(houses)  
**Kompleksitas Ruang:** O(log n) for sorting

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

```go
package main

// LeetCode #475: Heaters
// https://leetcode.com/problems/heaters/
// Difficulty: Medium
// Time: O(n log n + m log n) where n = len(heaters), m = len(houses)
// Space: O(log n) for sorting

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Heaters([]int{1, 2, 3}, []int{2}))
	fmt.Println(Heaters([]int{1, 2, 3, 4}, []int{1, 4}))
	fmt.Println(Heaters([]int{1, 5}, []int{2}))
}

func Heaters(houses []int, heaters []int) int {
	sort.Ints(heaters)
	maxRadius := 0

	for _, house := range houses {
		// Binary search to find nearest heater
		idx := sort.SearchInts(heaters, house)
		minDist := int(^uint(0) >> 1) // MaxInt

		if idx < len(heaters) {
			dist := heaters[idx] - house
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
		if idx > 0 {
			dist := house - heaters[idx-1]
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}

		if minDist > maxRadius {
			maxRadius = minDist
		}
	}

	return maxRadius
}
```
