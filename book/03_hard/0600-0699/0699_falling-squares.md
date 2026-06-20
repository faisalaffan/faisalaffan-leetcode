# 0699 — Falling Squares

## Deskripsi

**Soal:** [0699. Falling Squares](https://leetcode.com/problems/falling-squares/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #699: Falling Squares
// https://leetcode.com/problems/falling-squares/
// Difficulty: Hard
//
// For each square, check overlap with all previous squares.
// Base height = max height of overlapping squares below.
// Current height = base + size. Track running max.

func main() {
	// [[1,2],[2,3],[6,1]] => [2,5,5]
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
	// [[100,100],[200,100]] => [100,100]
	fmt.Println(fallingSquares([][]int{{100, 100}, {200, 100}}))
	// [[2,1],[2,9],[1,8]] => [1,10,18]
	fmt.Println(fallingSquares([][]int{{2, 1}, {2, 9}, {1, 8}}))
	// Single
	fmt.Println(fallingSquares([][]int{{5, 5}}))
	// [[1,2],[3,4]] => [2,4]
	fmt.Println(fallingSquares([][]int{{1, 2}, {3, 4}}))
}

func fallingSquares(positions [][]int) []int {
	n := len(positions)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	heights := make([]int, n)

	for i, p := range positions {
		left, size := p[0], p[1]
		right := left + size

		base := 0
		for j := 0; j < i; j++ {
			jLeft := positions[j][0]
			jRight := jLeft + positions[j][1]
			if left < jRight && right > jLeft {
				if heights[j] > base {
					base = heights[j]
				}
			}
		}
		heights[i] = base + size

		if i == 0 {
			ans[i] = heights[i]
		} else {
			if heights[i] > ans[i-1] {
				ans[i] = heights[i]
			} else {
				ans[i] = ans[i-1]
			}
		}
	}

	return ans
}
```
