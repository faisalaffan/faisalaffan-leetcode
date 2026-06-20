# 2852 — Sum Of Remoteness Of All Cells

## Deskripsi

**Soal:** [2852. Sum Of Remoteness Of All Cells](https://leetcode.com/problems/sum-of-remoteness-of-all-cells/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

**Fungsi Solusi:** `func SumOfRemotenessOfAllCells(grid [][]int) []int64`

## Solusi Go

```go
package main

// LeetCode #2852: Sum of Remoteness of All Cells
// https://leetcode.com/problems/sum-of-remoteness-of-all-cells/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func SumOfRemotenessOfAllCells(grid [][]int) []int64 {
	n := len(grid)
  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, n)
  // Iterasi seluruh elemen
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var totalSum int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				totalSum += int64(grid[i][j])
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int64, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				result[i*n+j] = totalSum - int64(grid[i][j])
			}
		}
	}

	return result
}

func main() {
	fmt.Println(SumOfRemotenessOfAllCells([][]int{{1, 2}, {3, 4}}))
	fmt.Println(SumOfRemotenessOfAllCells([][]int{{5, 0}, {0, 5}}))
}
```
