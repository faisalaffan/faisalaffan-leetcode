# 1033 — Moving Stones Until Consecutive

## Deskripsi

**Soal:** [1033. Moving Stones Until Consecutive](https://leetcode.com/problems/moving-stones-until-consecutive/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

> **Ide Kunci:** Sort the positions. Min moves = 0, 1, or 2 based on gaps.

## Solusi Go

```go
package main

// LeetCode #1033: Moving Stones Until Consecutive
// https://leetcode.com/problems/moving-stones-until-consecutive/
// Difficulty: Medium
//
// Approach: Sort the positions. Min moves = 0, 1, or 2 based on gaps.
//           Max moves = distance between extremes - 2.
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStones(1, 2, 5)) // [1,2]
	fmt.Println(numMovesStones(4, 3, 2)) // [0,0]
	fmt.Println(numMovesStones(3, 5, 1)) // [1,2]
}

func numMovesStones(a int, b int, c int) []int {
	stones := []int{a, b, c}
	sort.Ints(stones)
	x, y, z := stones[0], stones[1], stones[2]

	minMoves := 2
	if z-x == 2 {
		minMoves = 0
	} else if z-y <= 2 || y-x <= 2 {
		minMoves = 1
	}

	maxMoves := (z - x - 2)

	return []int{minMoves, maxMoves}
}
```
