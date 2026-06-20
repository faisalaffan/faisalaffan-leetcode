# 0573 — Squirrel Simulation

## Deskripsi

**Soal:** [0573. Squirrel Simulation](https://leetcode.com/problems/squirrel-simulation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of nuts  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #573: Squirrel Simulation
// https://leetcode.com/problems/squirrel-simulation/
// Difficulty: Medium [Paid]
// Time: O(n) where n = number of nuts
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	height := 5
	width := 7
	tree := []int{2, 2}
	squirrel := []int{4, 4}
	nuts := [][]int{{3, 0}, {2, 5}}
	fmt.Println(MinDistance(height, width, tree, squirrel, nuts))
}

func MinDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	total := 0
	maxSaving := math.MinInt32

	for _, nut := range nuts {
		nutToTree := abs(nut[0]-tree[0]) + abs(nut[1]-tree[1])
		nutToSquirrel := abs(nut[0]-squirrel[0]) + abs(nut[1]-squirrel[1])
		total += 2 * nutToTree
		saving := nutToTree - nutToSquirrel
		if saving > maxSaving {
			maxSaving = saving
		}
	}

	return total - maxSaving
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
