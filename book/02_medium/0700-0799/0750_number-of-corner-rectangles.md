# 0750 — Number Of Corner Rectangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countCornerRectangles(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(R * C^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #750: Number of Corner Rectangles
// https://leetcode.com/problems/number-of-corner-rectangles/
// Difficulty: Medium [Paid]
// Time: O(R * C^2)
// Space: O(1)

import "fmt"

func main() {
	grid := [][]int{
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1},
	}
	fmt.Println(countCornerRectangles(grid))
}

func countCornerRectangles(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	for c1 := 0; c1 < cols; c1++ {
		for c2 := c1 + 1; c2 < cols; c2++ {
			pairs := 0
			for r := 0; r < rows; r++ {
				if grid[r][c1] == 1 && grid[r][c2] == 1 {
					pairs++
				}
			}
			count += pairs * (pairs - 1) / 2
		}
	}

	return count
}
```
