# 3195 — Find The Minimum Area To Cover All Ones I

## Deskripsi

**Soal:** [3195. Find The Minimum Area To Cover All Ones I](https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minimumArea(grid [][]int) int`

## Solusi Go

```go
package main

// LeetCode #3195: Find the Minimum Area to Cover All Ones I
// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minimumArea(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	top, bottom, left, right := m, -1, n, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i < top {
					top = i
				}
				if i > bottom {
					bottom = i
				}
				if j < left {
					left = j
				}
				if j > right {
					right = j
				}
			}
		}
	}

	if bottom == -1 {
		return 0
	}
	return (bottom - top + 1) * (right - left + 1)
}

func main() {
	fmt.Println(minimumArea([][]int{{0, 1, 0}, {1, 0, 1}})) // Expected: 4
	fmt.Println(minimumArea([][]int{{0, 0}, {1, 1}}))        // Expected: 2
}
```
