# 0059 — Spiral Matrix Ii

## Deskripsi

**Soal:** [0059. Spiral Matrix Ii](https://leetcode.com/problems/spiral-matrix-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

**Fungsi Solusi:** `func generateMatrix(n int) [][]int`

## Solusi Go

```go
package main

// LeetCode #59: Spiral Matrix II
// https://leetcode.com/problems/spiral-matrix-ii/
// Difficulty: Medium

import "fmt"

func generateMatrix(n int) [][]int {
  // Membuat slice 2D untuk DP/tabel
	matrix := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1

	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			matrix[top][j] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		if top <= bottom {
			for j := right; j >= left; j-- {
				matrix[bottom][j] = num
				num++
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}

func main() {
	// Test case 1
	fmt.Println(generateMatrix(3)) // [[1 2 3] [8 9 4] [7 6 5]]

	// Test case 2
	fmt.Println(generateMatrix(1)) // [[1]]

	// Test case 3
	fmt.Println(generateMatrix(4))
}

// Time: O(n^2) | Space: O(n^2)
```
