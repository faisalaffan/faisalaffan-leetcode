# 2482 — Difference Between Ones And Zeros In Row And Column

## Deskripsi

**Soal:** [2482. Difference Between Ones And Zeros In Row And Column](https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m + n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2482: Difference Between Ones and Zeros in Row and Column
// https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m + n)
// diff[i][j] = onesRow[i] + onesCol[j] - zerosRow[i] - zerosCol[j]
// = 2*onesRow[i] + 2*onesCol[j] - m - n

import "fmt"

func main() {
	fmt.Println(onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
	// [[0,0,4],[0,0,4],[-2,-2,2]]

	fmt.Println(onesMinusZeros([][]int{{1, 1, 1}, {1, 1, 1}}))
	// [[5,5,5],[5,5,5]]
}

func onesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
  // Membuat slice untuk menyimpan hasil
	rowOnes := make([]int, m)
  // Membuat slice untuk menyimpan hasil
	colOnes := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowOnes[i]++
				colOnes[j]++
			}
		}
	}

  // Membuat slice 2D untuk DP/tabel
	diff := make([][]int, m)
	for i := 0; i < m; i++ {
		diff[i] = make([]int, n)
		for j := 0; j < n; j++ {
			diff[i][j] = 2*rowOnes[i] + 2*colOnes[j] - m - n
		}
	}
	return diff
}
```
