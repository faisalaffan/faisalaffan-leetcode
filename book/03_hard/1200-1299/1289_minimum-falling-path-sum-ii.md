# 1289 — Minimum Falling Path Sum Ii

## Deskripsi

**Soal:** [1289. Minimum Falling Path Sum Ii](https://leetcode.com/problems/minimum-falling-path-sum-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1289: Minimum Falling Path Sum II
// https://leetcode.com/problems/minimum-falling-path-sum-ii/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("1289. Minimum Falling Path Sum II")
	fmt.Println("[[1,2,3],[4,5,6],[7,8,9]]:", minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), "(expected 13)")
	fmt.Println("[[7]]:", minFallingPathSum([][]int{{7}}), "(expected 7)")
	fmt.Println("[[-37,51,-36,34,-22],[40,4,50,34,6],[62,32,-6,3,10],[4,12,-14,-34,13],[57,76,76,14,-40]]:", minFallingPathSum([][]int{{-37, 51, -36, 34, -22}, {40, 4, 50, 34, 6}, {62, 32, -6, 3, 10}, {4, 12, -14, -34, 13}, {57, 76, 76, 14, -40}}), "(expected -113)")
}

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}
	if n == 1 {
		return grid[0][0]
	}

	prevMin1, prevMin2 := 0, 0
	prevMinCol := -1

	for _, row := range grid {
		curMin1, curMin2 := math.MaxInt32, math.MaxInt32
		curMinCol := -1

		for j, val := range row {
			if j != prevMinCol {
				val += prevMin1
			} else {
				val += prevMin2
			}

			if val < curMin1 {
				curMin2 = curMin1
				curMin1 = val
				curMinCol = j
			} else if val < curMin2 {
				curMin2 = val
			}
		}

		prevMin1, prevMin2 = curMin1, curMin2
		prevMinCol = curMinCol
	}

	return prevMin1
}
```
