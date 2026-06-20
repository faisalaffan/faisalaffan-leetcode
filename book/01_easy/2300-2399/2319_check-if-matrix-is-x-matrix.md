# 2319 — Check If Matrix Is X Matrix

## Deskripsi

**Soal:** [2319. Check If Matrix Is X Matrix](https://leetcode.com/problems/check-if-matrix-is-x-matrix/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2319: Check if Matrix Is X-Matrix
// https://leetcode.com/problems/check-if-matrix-is-x-matrix/
// Difficulty: Easy
// Time O(n^2) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfMatrixIsXMatrix([][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}})) // true
	fmt.Println(CheckIfMatrixIsXMatrix([][]int{{5, 7, 0}, {0, 3, 1}, {0, 5, 0}}))                        // false
}

func CheckIfMatrixIsXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
```
