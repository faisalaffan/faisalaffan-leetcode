# 0939 — Minimum Area Rectangle

## Deskripsi

**Soal:** [0939. Minimum Area Rectangle](https://leetcode.com/problems/minimum-area-rectangle/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minAreaRect(points [][]int) int`

## Solusi Go

```go
package main

// LeetCode #939: Minimum Area Rectangle
// https://leetcode.com/problems/minimum-area-rectangle/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

// Time: O(n^2) | Space: O(n)
func minAreaRect(points [][]int) int {
  // Membuat map untuk pencarian O(1): key → value
	set := make(map[[2]int]bool)
	for _, p := range points {
		set[[2]int{p[0], p[1]}] = true
	}

	ans := math.MaxInt32
	n := len(points)

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			if x1 == x2 || y1 == y2 {
				continue
			}
			if set[[2]int{x1, y2}] && set[[2]int{x2, y1}] {
				area := abs(x1-x2) * abs(y1-y2)
				if area < ans {
					ans = area
				}
			}
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}))
	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}}))
}
```
