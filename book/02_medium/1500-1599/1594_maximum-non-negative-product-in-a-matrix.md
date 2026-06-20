# 1594 — Maximum Non Negative Product In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaxProductPath(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(R*C), Space: O(C)  |  **Ruang:** O(C)


## 💻 Solusi Go

```go
package main

// LeetCode #1594: Maximum Non Negative Product in a Matrix
// https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/
// Difficulty: Medium

import "fmt"

func main() {
	grid1 := [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}
	fmt.Println(MaxProductPath(grid1))

	grid2 := [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}
	fmt.Println(MaxProductPath(grid2))

	grid3 := [][]int{{1, 3}, {0, -4}}
	fmt.Println(MaxProductPath(grid3))
}

func MaxProductPath(grid [][]int) int {
	// Time: O(R*C), Space: O(C)
	const mod = 1_000_000_007

	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return -1
	}

	// minDP[r][c] = minimum product to reach (r,c)
	// maxDP[r][c] = maximum product to reach (r,c)
  // Matriks 2D
	minDP := make([][]int64, rows)
  // Matriks 2D
	maxDP := make([][]int64, rows)
	for i := 0; i < rows; i++ {
		minDP[i] = make([]int64, cols)
		maxDP[i] = make([]int64, cols)
	}

	maxDP[0][0] = int64(grid[0][0])
	minDP[0][0] = int64(grid[0][0])

	// First row
	for c := 1; c < cols; c++ {
		val := int64(grid[0][c])
		maxDP[0][c] = maxDP[0][c-1] * val
		minDP[0][c] = maxDP[0][c-1] * val
	}

	// First column
	for r := 1; r < rows; r++ {
		val := int64(grid[r][0])
		maxDP[r][0] = maxDP[r-1][0] * val
		minDP[r][0] = maxDP[r-1][0] * val
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			val := int64(grid[r][c])

			options := []int64{
				maxDP[r-1][c] * val,
				maxDP[r][c-1] * val,
				minDP[r-1][c] * val,
				minDP[r][c-1] * val,
			}

			maxVal := options[0]
			minVal := options[0]
			for _, opt := range options {
				if opt > maxVal {
					maxVal = opt
				}
				if opt < minVal {
					minVal = opt
				}
			}

			maxDP[r][c] = maxVal
			minDP[r][c] = minVal
		}
	}

	if maxDP[rows-1][cols-1] < 0 {
		return -1
	}

	return int(maxDP[rows-1][cols-1] % mod)
}
```
