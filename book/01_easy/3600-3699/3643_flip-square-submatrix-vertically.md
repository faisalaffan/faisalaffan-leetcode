# 3643 — Flip Square Submatrix Vertically

## Deskripsi

**Soal:** [3643. Flip Square Submatrix Vertically](https://leetcode.com/problems/flip-square-submatrix-vertically/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(k^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3643: Flip Square Submatrix Vertically
// https://leetcode.com/problems/flip-square-submatrix-vertically/
// Difficulty: Easy

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(FlipSquareSubmatrixVertically(grid, 1, 1, 3))
}

// Time: O(k^2)
// Space: O(1)
func FlipSquareSubmatrixVertically(grid [][]int, x, y, k int) [][]int {
	for i := x; i < x+k/2; i++ {
		i2 := x + k - 1 - (i - x)
		for j := y; j < y+k; j++ {
			grid[i][j], grid[i2][j] = grid[i2][j], grid[i][j]
		}
	}
	return grid
}
```
