# 0931 — Minimum Falling Path Sum

## Deskripsi

**Soal:** [0931. Minimum Falling Path Sum](https://leetcode.com/problems/minimum-falling-path-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1) — modifies input in place

**Algoritma:** —

**Fungsi Solusi:** `func minFallingPathSum(matrix [][]int) int`

## Solusi Go

```go
package main

// LeetCode #931: Minimum Falling Path Sum
// https://leetcode.com/problems/minimum-falling-path-sum/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

// Time: O(n^2) | Space: O(1) — modifies input in place
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			best := matrix[i-1][j]
			if j > 0 {
				best = min(best, matrix[i-1][j-1])
			}
			if j+1 < n {
				best = min(best, matrix[i-1][j+1])
			}
			matrix[i][j] += best
		}
	}

	ans := math.MaxInt32
	for _, v := range matrix[n-1] {
		ans = min(ans, v)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
	fmt.Println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))
}
```
