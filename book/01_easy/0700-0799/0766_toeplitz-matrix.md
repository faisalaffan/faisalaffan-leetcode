# 0766 — Toeplitz Matrix

## Deskripsi

**Soal:** [0766. Toeplitz Matrix](https://leetcode.com/problems/toeplitz-matrix/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(m*n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #766: Toeplitz Matrix
// https://leetcode.com/problems/toeplitz-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}})) // true
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))                            // false
}

// isToeplitzMatrix checks if every diagonal from top-left to bottom-right has the same element.
// Time: O(m*n). Space: O(1).
func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
```
