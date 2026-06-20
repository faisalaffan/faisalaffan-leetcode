# 3148 — Maximum Difference Score In A Grid

## Deskripsi

**Soal:** [3148. Maximum Difference Score In A Grid](https://leetcode.com/problems/maximum-difference-score-in-a-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** —

**Fungsi Solusi:** `func maxScore(grid [][]int) int`

## Solusi Go

```go
package main

// LeetCode #3148: Maximum Difference Score in a Grid
// https://leetcode.com/problems/maximum-difference-score-in-a-grid/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import (
	"fmt"
	"math"
)

func maxScore(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat slice 2D untuk DP/tabel
	minVal := make([][]int, m)
  // Iterasi seluruh elemen
	for i := range minVal {
		minVal[i] = make([]int, n)
	}

	ans := math.MinInt32

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prevMin := math.MaxInt32
			if i > 0 {
				prevMin = min(prevMin, minVal[i-1][j])
			}
			if j > 0 {
				prevMin = min(prevMin, minVal[i][j-1])
			}

			if i > 0 || j > 0 {
				ans = max(ans, grid[i][j]-prevMin)
			}

			minVal[i][j] = grid[i][j]
			if i > 0 {
				minVal[i][j] = min(minVal[i][j], minVal[i-1][j])
			}
			if j > 0 {
				minVal[i][j] = min(minVal[i][j], minVal[i][j-1])
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(maxScore([][]int{{9, 5, 7, 3}, {8, 9, 6, 1}, {6, 7, 14, 3}, {2, 5, 3, 1}})) // Expected: 9
	fmt.Println(maxScore([][]int{{4, 3, 2}, {3, 2, 1}}))                                      // Expected: -1
}
```
