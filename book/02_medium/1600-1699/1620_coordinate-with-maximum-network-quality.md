# 1620 — Coordinate With Maximum Network Quality

## Deskripsi

**Soal:** [1620. Coordinate With Maximum Network Quality](https://leetcode.com/problems/coordinate-with-maximum-network-quality/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * R^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1620: Coordinate With Maximum Network Quality
// https://leetcode.com/problems/coordinate-with-maximum-network-quality/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(BestCoordinate([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2))
	fmt.Println(BestCoordinate([][]int{{23, 11, 21}}, 9))
	fmt.Println(BestCoordinate([][]int{{2, 1, 9}, {0, 1, 9}}, 2))
}

func BestCoordinate(towers [][]int, radius int) []int {
	// Time: O(N * R^2), Space: O(1)
	// Signal quality for a point (x,y) = sum of floor(tower_signal / (1 + d))
	// where d = Euclidean distance from tower to (x,y)
	// Range of coordinates is bounded by tower positions +/- radius

	minX, minY := 51, 51
	maxX, maxY := 0, 0

	for _, t := range towers {
		if t[0] < minX {
			minX = t[0]
		}
		if t[0] > maxX {
			maxX = t[0]
		}
		if t[1] < minY {
			minY = t[1]
		}
		if t[1] > maxY {
			maxY = t[1]
		}
	}

	bestX, bestY := 0, 0
	bestQuality := 0

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			quality := 0
			for _, t := range towers {
				d := sqrtDist(x, y, t[0], t[1])
				if d > radius*radius {
					continue
				}
				// signal = floor(t[2] / (1 + sqrt(d)))
				// We avoid sqrt by computing: quality += t[2] / (1 + sqrt(d))
				// For integer comparisons we use the formula directly
				signal := float64(t[2]) / (1.0 + sqrt(float64(d)))
				quality += int(signal)
			}
			if quality > bestQuality || (quality == bestQuality && (x < bestX || (x == bestX && y < bestY))) {
				bestQuality = quality
				bestX, bestY = x, y
			}
		}
	}

	return []int{bestX, bestY}
}

func sqrtDist(x1, y1, x2, y2 int) int {
	dx := x1 - x2
	dy := y1 - y2
	return dx*dx + dy*dy
}

func sqrt(n float64) float64 {
	if n <= 0 {
		return 0
	}
	x := n
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + n/x) / 2
	}
	return x
}
```
