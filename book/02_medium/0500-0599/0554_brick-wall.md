# 0554 — Brick Wall

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func LeastBricks(wall [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m) where n = rows, m = avg bricks per row  |  **Ruang:** O(n * m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #554: Brick Wall
// https://leetcode.com/problems/brick-wall/
// Difficulty: Medium
// Time: O(n * m) where n = rows, m = avg bricks per row
// Space: O(n * m)

import "fmt"

func main() {
	wall := [][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	}
	fmt.Println(LeastBricks(wall))
}

func LeastBricks(wall [][]int) int {
  // HashMap: O(1) lookup
	gapCount := make(map[int]int)

	for _, row := range wall {
		pos := 0
  // Linear scan O(n)
		for i := 0; i < len(row)-1; i++ {
			pos += row[i]
			gapCount[pos]++
		}
	}

	maxGaps := 0
	for _, count := range gapCount {
		if count > maxGaps {
			maxGaps = count
		}
	}

	return len(wall) - maxGaps
}
```
