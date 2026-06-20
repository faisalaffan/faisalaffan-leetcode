# 2852 — Sum Of Remoteness Of All Cells

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func SumOfRemotenessOfAllCells(grid [][]int) []int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(n^2)


## 💻 Solusi Go

```go
package main

// LeetCode #2852: Sum of Remoteness of All Cells
// https://leetcode.com/problems/sum-of-remoteness-of-all-cells/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func SumOfRemotenessOfAllCells(grid [][]int) []int64 {
	n := len(grid)
  // Matriks 2D
	visited := make([][]bool, n)
  // Range loop
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

  // Alokasi slice
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
