# 2249 — Count Lattice Points Inside A Circle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countLatticePoints(circles [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * r^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2249: Count Lattice Points Inside a Circle
// https://leetcode.com/problems/count-lattice-points-inside-a-circle/
// Difficulty: Medium
// Time: O(n * r^2) | Space: O(n)

import "fmt"

func countLatticePoints(circles [][]int) int {
  // HashMap: O(1) lookup
	points := make(map[[2]int]bool)

	for _, c := range circles {
		x, y, r := c[0], c[1], c[2]
		rr := r * r
		for dx := -r; dx <= r; dx++ {
			for dy := -r; dy <= r; dy++ {
				if dx*dx+dy*dy <= rr {
					points[[2]int{x + dx, y + dy}] = true
				}
			}
		}
	}
	return len(points)
}

func main() {
	// Test case 1
	fmt.Println(countLatticePoints([][]int{{2, 2, 1}}))
	// Expected: 5

	// Test case 2
	fmt.Println(countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
	// Expected: 16
}
```
