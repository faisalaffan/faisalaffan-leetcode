# 3417 — Zigzag Grid Traversal With Skip

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ZigzagGridTraversalWithSkip(grid [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n). Space: O(m * n).  |  **Ruang:** O(m * n).


## 💻 Solusi Go

```go
package main

// LeetCode #3417: Zigzag Grid Traversal With Skip
// https://leetcode.com/problems/zigzag-grid-traversal-with-skip/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ZigzagGridTraversalWithSkip([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(ZigzagGridTraversalWithSkip([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}))
}

// ZigzagGridTraversalWithSkip traverses the grid in zigzag order but skips every other element.
// Time: O(m * n). Space: O(m * n).
func ZigzagGridTraversalWithSkip(grid [][]int) []int {
	if len(grid) == 0 {
		return nil
	}
	m, n := len(grid), len(grid[0])
	result := []int{}
	skip := false
	for i := 0; i < m; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				if !skip {
					result = append(result, grid[i][j])
				}
				skip = !skip
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				if !skip {
					result = append(result, grid[i][j])
				}
				skip = !skip
			}
		}
	}
	return result
}
```
