# 1706 — Where Will The Ball Fall

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findBall(grid [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(m * n), Space: O(1) (excluding output)  |  **Ruang:** O(1) (excluding output)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1706: Where Will the Ball Fall
// https://leetcode.com/problems/where-will-the-ball-fall/
// Difficulty: Medium
// Time: O(m * n), Space: O(1) (excluding output)

import "fmt"

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
  // Alokasi slice
	result := make([]int, n)

	for col := 0; col < n; col++ {
		c := col
		for r := 0; r < m; r++ {
			// If the cell is 1 (sloping right), check the cell to the right
			if grid[r][c] == 1 {
				if c+1 >= n || grid[r][c+1] == -1 {
					c = -1
					break
				}
				c++
			} else {
				// Cell is -1 (sloping left), check the cell to the left
				if c-1 < 0 || grid[r][c-1] == 1 {
					c = -1
					break
				}
				c--
			}
		}
		result[col] = c
	}
	return result
}

func main() {
	fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
	// Expected: [1, -1, -1, -1, -1]

	fmt.Println(findBall([][]int{{-1}})) // Expected: [-1]

	fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
	// Expected: [0, 1, 2, 3, 4, -1]
}
```
