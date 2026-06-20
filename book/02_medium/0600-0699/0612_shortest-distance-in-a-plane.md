# 0612 — Shortest Distance In A Plane

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ShortestDistance(points [][]float64) float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #612: Shortest Distance in a Plane
// https://leetcode.com/problems/shortest-distance-in-a-plane/
// Difficulty: Medium [Paid]
// Time: O(n^2)
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	points := [][]float64{
		{-1, -1},
		{0, 0},
		{1, 1},
		{2, 2},
	}
	fmt.Printf("%.4f\n", ShortestDistance(points))
}

func ShortestDistance(points [][]float64) float64 {
	if len(points) < 2 {
		return 0
	}

	minDist := math.MaxFloat64
  // Linear scan O(n)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}
```
