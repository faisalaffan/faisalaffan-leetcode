# 1706 — Where Will The Ball Fall

## Deskripsi

**Soal:** [1706. Where Will The Ball Fall](https://leetcode.com/problems/where-will-the-ball-fall/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n), Space: O(1) (excluding output)  
**Kompleksitas Ruang:** O(1) (excluding output)

**Algoritma:** —

**Fungsi Solusi:** `func findBall(grid [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #1706: Where Will the Ball Fall
// https://leetcode.com/problems/where-will-the-ball-fall/
// Difficulty: Medium
// Time: O(m * n), Space: O(1) (excluding output)

import "fmt"

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)

	for col := 0; col < n; col++ {
		c := col
		for r := 0; r < m; r++ {
			// If the cell is 1 (sloping right), check the cell to the right
			if grid[r][c] == 1 {
				if c+1 >= n || grid[r][c+1] == -1 {
					c = -1
					break
				}
				c++
			} else {
				// Cell is -1 (sloping left), check the cell to the left
				if c-1 < 0 || grid[r][c-1] == 1 {
					c = -1
					break
				}
				c--
			}
		}
		result[col] = c
	}
	return result
}

func main() {
	fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
	// Expected: [1, -1, -1, -1, -1]

	fmt.Println(findBall([][]int{{-1}})) // Expected: [-1]

	fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
	// Expected: [0, 1, 2, 3, 4, -1]
}
```
