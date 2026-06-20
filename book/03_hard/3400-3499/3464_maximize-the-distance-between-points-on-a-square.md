# 3464 — Maximize The Distance Between Points On A Square

## Deskripsi

**Soal:** [3464. Maximize The Distance Between Points On A Square](https://leetcode.com/problems/maximize-the-distance-between-points-on-a-square/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

> **Ide Kunci:** Binary search on minimum distance. For a given min

## Solusi Go

```go
package main

// LeetCode #3464: Maximize the Distance Between Points on a Square
// https://leetcode.com/problems/maximize-the-distance-between-points-on-a-square/
// Difficulty: Hard
//
// Place k points on the perimeter of a square of given side length
// at integer coordinates. Maximize the minimum Manhattan distance
// between any two chosen points.
//
// Approach: Binary search on minimum distance. For a given min
// distance candidate, check if k points can be placed on the
// perimeter with that minimum separation.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxDistance(4, [][]int{{0, 0}, {4, 0}}, 3))
	// Example 2
	fmt.Println(maxDistance(10, [][]int{{0, 0}, {10, 0}}, 2))
	// Edge: k = 1
	fmt.Println(maxDistance(5, [][]int{}, 1))
}

func maxDistance(side int, points [][]int, k int) int {
	// Map each point to its position along the perimeter (clockwise)
	// Start from (0,0), go right, then up, then left, then down
  // Membuat slice untuk menyimpan hasil
	pos := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		if y == 0 {
			pos[i] = x // bottom edge, left to right
		} else if x == side {
			pos[i] = side + y // right edge, bottom to top
		} else if y == side {
			pos[i] = side*3 - x // top edge, right to left
		} else {
			pos[i] = side*4 - y // left edge, top to bottom
		}
	}

	perimeter := side * 4
	can := func(minDist int) bool {
		// Try starting at each point
		for start := 0; start < len(points); start++ {
			cnt := 1
			last := pos[start]
			for i := 1; i < len(points); i++ {
				idx := (start + i) % len(points)
				dist := pos[idx] - last
				if dist < 0 {
					dist += perimeter
				}
				if dist >= minDist {
					cnt++
					last = pos[idx]
				}
				if cnt >= k {
					return true
				}
			}
			// Also check wrap-around
			if cnt >= k {
				return true
			}
		}
		return false
	}

	lo, hi := 0, perimeter
	for lo < hi {
		mid := lo + (hi-lo+1)/2
		if can(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}
```
