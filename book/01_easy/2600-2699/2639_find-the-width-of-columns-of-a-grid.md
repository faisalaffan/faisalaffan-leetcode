# 2639 — Find The Width Of Columns Of A Grid

## Deskripsi

**Soal:** [2639. Find The Width Of Columns Of A Grid](https://leetcode.com/problems/find-the-width-of-columns-of-a-grid/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2639: Find the Width of Columns of a Grid
// https://leetcode.com/problems/find-the-width-of-columns-of-a-grid/
// Difficulty: Easy
// Time: O(m * n) | Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{1}, {22}, {333}}))
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}}))
}

func FindTheWidthOfColumnsOfAGrid(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
	for j := 0; j < n; j++ {
		maxLen := 0
		for i := 0; i < m; i++ {
			width := len(strconv.Itoa(grid[i][j]))
			if width > maxLen {
				maxLen = width
			}
		}
		ans[j] = maxLen
	}
	return ans
}
```
