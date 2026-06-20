# 0073 — Set Matrix Zeroes

## Deskripsi

**Soal:** [0073. Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func setZeroes(matrix [][]int) `

## Solusi Go

```go
package main

// LeetCode #73: Set Matrix Zeroes
// https://leetcode.com/problems/set-matrix-zeroes/
// Difficulty: Medium

import "fmt"

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero := false
	firstColZero := false

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	// Test case 1
	m1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m1)
	fmt.Println(m1) // [[1 0 1] [0 0 0] [1 0 1]]

	// Test case 2
	m2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m2)
	fmt.Println(m2) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}

// Time: O(m*n) | Space: O(1)
```
