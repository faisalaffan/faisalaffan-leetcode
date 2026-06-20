# 3380 — Maximum Area Rectangle With Point Constraints I

## Deskripsi

**Soal:** [3380. Maximum Area Rectangle With Point Constraints I](https://leetcode.com/problems/maximum-area-rectangle-with-point-constraints-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^4) Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3380: Maximum Area Rectangle With Point Constraints I
// https://leetcode.com/problems/maximum-area-rectangle-with-point-constraints-i/
// Difficulty: Medium
// Time: O(n^4) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxRectangleArea([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}}))          // 4
	fmt.Println(maxRectangleArea([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}})) // -1
}

func maxRectangleArea(points [][]int) int {
	n := len(points)
  // Membuat slice untuk menyimpan hasil
	sp := make([][2]int, n)
	for i := 0; i < n; i++ {
		sp[i] = [2]int{points[i][0], points[i][1]}
	}

	ans := -1

	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				for d := c + 1; d < n; d++ {
					if area, ok := isRect3380([]int{a, b, c, d}, sp); ok && area > ans {
						ans = area
					}
				}
			}
		}
	}
	return ans
}

func isRect3380(idx []int, pts [][2]int) (int, bool) {
  // Membuat map untuk pencarian O(1): key → value
	xSet := make(map[int]bool)
  // Membuat map untuk pencarian O(1): key → value
	ySet := make(map[int]bool)
	for _, i := range idx {
		xSet[pts[i][0]] = true
		ySet[pts[i][1]] = true
	}
	if len(xSet) != 2 || len(ySet) != 2 {
		return 0, false
	}

	var xs, ys []int
	for x := range xSet {
		xs = append(xs, x)
	}
	for y := range ySet {
		ys = append(ys, y)
	}
	sort.Ints(xs)
	sort.Ints(ys)
	x1, x2 := xs[0], xs[1]
	y1, y2 := ys[0], ys[1]

  // Membuat map untuk pencarian O(1): key → value
	cornerSet := make(map[[2]int]bool)
	for _, i := range idx {
		cornerSet[pts[i]] = true
	}
	expected := [][2]int{{x1, y1}, {x1, y2}, {x2, y1}, {x2, y2}}
	for _, p := range expected {
		if !cornerSet[p] {
			return 0, false
		}
	}

	for _, p := range pts {
		if cornerSet[p] {
			continue
		}
		if p[0] >= x1 && p[0] <= x2 && p[1] >= y1 && p[1] <= y2 {
			return 0, false
		}
	}

	return (x2 - x1) * (y2 - y1), true
}
```
