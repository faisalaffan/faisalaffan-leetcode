# 3454 — Separate Squares Ii

## Deskripsi

**Soal:** [3454. Separate Squares Ii](https://leetcode.com/problems/separate-squares-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

> **Ide Kunci:** Binary search on y-coordinate. Compute total area below

## Solusi Go

```go
package main

// LeetCode #3454: Separate Squares II
// https://leetcode.com/problems/separate-squares-ii/
// Difficulty: Hard
//
// Given squares [x, y, side], find the minimum y-coordinate where
// a horizontal line splits the total area of all squares into
// equal halves.
//
// Approach: Binary search on y-coordinate. Compute total area below
// a given y. For each square, area below line is the portion of
// the square that lies below y.

import "fmt"

func main() {
	// Example 1
	fmt.Println(separateSquares([][]int{{0, 0, 2}, {1, 1, 2}}))
	// Example 2
	fmt.Println(separateSquares([][]int{{0, 0, 1}, {0, 2, 1}}))
	// Edge: single square
	fmt.Println(separateSquares([][]int{{0, 0, 3}}))
}

func separateSquares(squares [][]int) float64 {
	minY, maxY := 1<<30, 0
	for _, sq := range squares {
		sy, l := sq[1], sq[2]
		if sy < minY {
			minY = sy
		}
		if sy+l > maxY {
			maxY = sy + l
		}
	}

	areaBelow := func(y float64) float64 {
		area := 0.0
		for _, sq := range squares {
			sy := float64(sq[1])
			l := float64(sq[2])
			if y <= sy {
				continue
			}
			top := sy + l
			if y >= top {
				area += l * l
			} else {
				area += l * (y - sy)
			}
		}
		return area
	}

	totalArea := areaBelow(float64(maxY))
	halfArea := totalArea / 2.0

	lo, hi := float64(minY), float64(maxY)
	for i := 0; i < 100; i++ {
		mid := (lo + hi) / 2.0
		if areaBelow(mid) >= halfArea {
			hi = mid
		} else {
			lo = mid
		}
	}
	return (lo + hi) / 2.0
}
```
