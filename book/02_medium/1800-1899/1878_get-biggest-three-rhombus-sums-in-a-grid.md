# 1878 — Get Biggest Three Rhombus Sums In A Grid

## Deskripsi

**Soal:** [1878. Get Biggest Three Rhombus Sums In A Grid](https://leetcode.com/problems/get-biggest-three-rhombus-sums-in-a-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n*min(m,n)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1878: Get Biggest Three Rhombus Sums in a Grid
// https://leetcode.com/problems/get-biggest-three-rhombus-sums-in-a-grid/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(GetBiggestThree([][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}}))
	fmt.Println(GetBiggestThree([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(GetBiggestThree([][]int{{7, 7, 7}}))
}

// Time: O(m*n*min(m,n)), Space: O(1)
func GetBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
  // Membuat slice untuk menyimpan hasil
	top3 := make([]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Single cell rhombus (size 0)
			addToTop3(&top3, grid[i][j])
			// Expand rhombus size
			maxSize := min(min(i, m-1-i), min(j, n-1-j))
			for s := 1; s <= maxSize; s++ {
				sum := 0
				// Traverse 4 edges of rhombus
				// Top to right: (i-k, j+k) for k=0..s
				// Right to bottom: (i+s-k, j+s-k) for k=0..s
				for k := 0; k < s; k++ {
					sum += grid[i-k][j+k]     // top-left to top-right
					sum += grid[i+k][j+s-k]   // top-right to bottom
				}
				for k := 0; k < s; k++ {
					sum += grid[i+s-k][j-k]   // bottom to bottom-left
					sum += grid[i-k][j-(s-k)] // bottom-left to top
				}
				// Corners counted twice, subtract once
				// Actually the 4 loops above need careful boundary handling
				// Let's fix: iterate each edge separately
				sum = 0
				// Top-right edge: (i-t, j+t) for t=0..s
				for t := 0; t <= s; t++ {
					sum += grid[i-t][j+t]
				}
				// Right-bottom edge: (i+s-t, j+s-t) for t=1..s
				for t := 1; t <= s; t++ {
					sum += grid[i+s-t][j+s-t]
				}
				// Bottom-left edge: (i+t, j-t) for t=1..s-1
				for t := 1; t < s; t++ {
					sum += grid[i-t][j-t]
				}
				// Left-top edge: (i-t, j-t) for t=1..s-1
				for t := 1; t < s; t++ {
					sum += grid[i+t][j-t]
				}
				addToTop3(&top3, sum)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(top3)))
	if len(top3) > 3 {
		top3 = top3[:3]
	}
	return top3
}

func addToTop3(top3 *[]int, val int) {
	for _, v := range *top3 {
		if v == val {
			return
		}
	}
	*top3 = append(*top3, val)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
