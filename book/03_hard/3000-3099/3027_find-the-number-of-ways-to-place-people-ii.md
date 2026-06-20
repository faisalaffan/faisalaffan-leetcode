# 3027 — Find The Number Of Ways To Place People Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numberOfPairs(points [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3027: Find the Number of Ways to Place People II
// https://leetcode.com/problems/find-the-number-of-ways-to-place-people-ii/
// Difficulty: Hard
//
// Given an array of points on a 2D plane, count how many pairs (i, j)
// (i as Alice, j as Bob) satisfy:
//   - Alice is strictly above and strictly to the left of Bob, OR
//     Alice is directly above AND strictly to the left, OR
//     Alice is strictly above AND directly to the left.
//   - No other point is inside or on the rectangle defined by Alice and Bob
//     (excluding the rectangle border at Alice and Bob themselves).
//
// Approach: Sort + geometry
//   Sort points by x ascending, then y descending.
//   For each Alice point (i), iterate over Bob points (j > i) that
//   are to the right and below Alice. Maintain the minimum allowed
//   x and maximum allowed y from previously counted Bobs to avoid
//   counting points with another point inside the rectangle.

import (
	"fmt"
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	ans := 0
	// Sort by x ascending, then y descending
  // Custom sort
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

  // Linear scan O(n)
	for i := 0; i < len(points)-1; i++ {
		xMax := math.MaxInt32
		yMin := math.MinInt32
		for j := i + 1; j < len(points); j++ {
			// Bob must be to the right (or same x but lower y) of Alice
			if points[j][0] > points[i][0]-1 && points[j][0] < xMax &&
				points[j][1] > yMin && points[j][1] < points[i][1]+1 {
				ans++
				xMax = points[j][0]
				yMin = points[j][1]
			}
		}
	}
	return ans
}

func main() {
	// Test 1: Simple points
	fmt.Println("Test 1:", numberOfPairs([][]int{{1, 1}, {2, 2}, {3, 3}}))

	// Test 2
	fmt.Println("Test 2:", numberOfPairs([][]int{{1, 2}, {2, 1}, {3, 1}}))

	// Test 3: Same x
	fmt.Println("Test 3:", numberOfPairs([][]int{{1, 3}, {1, 2}, {1, 1}}))

	// Test 4: Single pair
	fmt.Println("Test 4:", numberOfPairs([][]int{{0, 0}, {1, 0}}))

	// Test 5: All in a line diagonal
	result := numberOfPairs([][]int{{0, 3}, {1, 2}, {2, 1}, {3, 0}})
	fmt.Println("Test 5:", result)

	// Test 6
	fmt.Println("Test 6:", numberOfPairs([][]int{{3, 1}, {1, 1}, {0, 0}}))
}
```
