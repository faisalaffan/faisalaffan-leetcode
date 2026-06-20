# 3453 — Separate Squares I

## Deskripsi

**Soal:** [3453. Separate Squares I](https://leetcode.com/problems/separate-squares-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func separateSquares(squares [][]int) float64`

## Solusi Go

```go
package main

// LeetCode #3453: Separate Squares I
// https://leetcode.com/problems/separate-squares-i/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"math"
	"sort"
)

func separateSquares(squares [][]int) float64 {
	type event struct {
		y   float64
		dy  float64
		x1  float64
		x2  float64
	}
	var events []event
	var totalArea float64

	for _, sq := range squares {
		x, y, l := float64(sq[0]), float64(sq[1]), float64(sq[2])
		events = append(events, event{y, l, x, x + l})
		totalArea += l * l
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	halfArea := totalArea / 2.0

	lo, hi := float64(squares[0][1]), float64(squares[0][1])
	for _, sq := range squares {
		y := float64(sq[1])
		ye := y + float64(sq[2])
		if y < lo {
			lo = y
		}
		if ye > hi {
			hi = ye
		}
	}

	areaBelow := func(y float64) float64 {
		var area float64
		for _, sq := range squares {
			sy, l := float64(sq[1]), float64(sq[2])
			sye := sy + l
			if y <= sy {
				continue
			}
			overlap := math.Min(y, sye) - sy
			if overlap > 0 {
				area += overlap * l
			}
		}
		return area
	}

	for i := 0; i < 60; i++ {
		mid := (lo + hi) / 2
		if areaBelow(mid) < halfArea {
			lo = mid
		} else {
			hi = mid
		}
	}
	return (lo + hi) / 2
}

func main() {
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 2}, {1, 1, 1}})) // 1.00000
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 1}, {2, 2, 1}})) // 1.00000
}
```
