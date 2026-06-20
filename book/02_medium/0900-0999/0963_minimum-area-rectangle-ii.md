# 0963 — Minimum Area Rectangle Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minAreaFreeRect(points [][]int) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^3)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #963: Minimum Area Rectangle II
// https://leetcode.com/problems/minimum-area-rectangle-ii/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

// Time: O(n^3) | Space: O(n)
func minAreaFreeRect(points [][]int) float64 {
	n := len(points)
  // HashMap: O(1) lookup
	set := make(map[[2]int]bool)
	for _, p := range points {
		set[[2]int{p[0], p[1]}] = true
	}

	ans := math.MaxFloat64

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x2, y2 := points[j][0], points[j][1]
			dx1, dy1 := x2-x1, y2-y1
			for k := j + 1; k < n; k++ {
				if k == i {
					continue
				}
				x3, y3 := points[k][0], points[k][1]
				dx2, dy2 := x3-x1, y3-y1

				// Check perpendicular
				if dx1*dx2+dy1*dy2 != 0 {
					continue
				}

				// Fourth point: (x2 + dx2, y2 + dy2)
				x4, y4 := x2+dx2, y2+dy2
				if !set[[2]int{x4, y4}] {
					continue
				}

				area := math.Sqrt(float64(dx1*dx1+dy1*dy1)) * math.Sqrt(float64(dx2*dx2+dy2*dy2))
				if area < ans {
					ans = area
				}
			}
		}
	}

	if ans == math.MaxFloat64 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(minAreaFreeRect([][]int{{1, 2}, {2, 1}, {1, 0}, {0, 1}}))
	fmt.Println(minAreaFreeRect([][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}, {1, 1}}))
}
```
