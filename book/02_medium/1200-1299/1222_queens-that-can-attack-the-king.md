# 1222 — Queens That Can Attack The King

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func queensAttacktheKing(queens [][]int, king []int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) where n = number of queens  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1222: Queens That Can Attack the King
// https://leetcode.com/problems/queens-that-can-attack-the-king/
// Difficulty: Medium

// Queens can attack the king if no other piece blocks.
// From king's position, check 8 directions for nearest queen.

// Time: O(n) where n = number of queens
// Space: O(1)

func queensAttacktheKing(queens [][]int, king []int) [][]int {
  // HashMap: O(1) lookup
	queenSet := make(map[[2]int]bool)
	for _, q := range queens {
		queenSet[[2]int{q[0], q[1]}] = true
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

  // Matriks 2D
	result := make([][]int, 0)

	for _, d := range dirs {
		r, c := king[0]+d[0], king[1]+d[1]
		for r >= 0 && r < 8 && c >= 0 && c < 8 {
			if queenSet[[2]int{r, c}] {
				result = append(result, []int{r, c})
				break
			}
			r += d[0]
			c += d[1]
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [[0 1] [1 0] [3 3]])\n",
		queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}},
			[]int{0, 0}))

	fmt.Printf("%v (expected: [[2 2] [4 2]])\n",
		queensAttacktheKing([][]int{{0, 0}, {2, 2}, {4, 2}, {5, 5}},
			[]int{3, 2}))
}
```
