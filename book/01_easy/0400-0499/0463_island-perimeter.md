# 0463 — Island Perimeter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func IslandPerimeter(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #463: Island Perimeter
// https://leetcode.com/problems/island-perimeter/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(1)
func IslandPerimeter(grid [][]int) int {
	perimeter := 0
  // Linear scan O(n)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				if i > 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}

func main() {
	fmt.Println(IslandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
	fmt.Println(IslandPerimeter([][]int{{1}}))
}
```
