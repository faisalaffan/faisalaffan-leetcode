# 3588 — Find Maximum Area Of A Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindMaximumAreaOfATriangle(points [][]int) float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3588: Find Maximum Area of a Triangle
// https://leetcode.com/problems/find-maximum-area-of-a-triangle/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import (
	"fmt"
	"math"
)

func main() {
	// Test case 1
	points := [][]int{{0, 0}, {1, 0}, {0, 1}}
	fmt.Println("Test 1:", FindMaximumAreaOfATriangle(points))
	// Test case 2
	points2 := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	fmt.Println("Test 2:", FindMaximumAreaOfATriangle(points2))
	// Test case 3
	points3 := [][]int{{0, 0}}
	fmt.Println("Test 3:", FindMaximumAreaOfATriangle(points3))
}

func FindMaximumAreaOfATriangle(points [][]int) float64 {
	n := len(points)
	maxArea := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := math.Abs(float64(
					points[i][0]*(points[j][1]-points[k][1]) +
						points[j][0]*(points[k][1]-points[i][1]) +
						points[k][0]*(points[i][1]-points[j][1]),
				)) / 2.0
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
```
