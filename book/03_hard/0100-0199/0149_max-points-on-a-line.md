# 0149 — Max Points On A Line

## Deskripsi

**Soal:** [0149. Max Points On A Line](https://leetcode.com/problems/max-points-on-a-line/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func maxPoints(points [][]int) int`

## Solusi Go

```go
package main

// LeetCode #149: Max Points on a Line
// https://leetcode.com/problems/max-points-on-a-line/
// Difficulty: Hard

import (
	"fmt"
)

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxCount := 0

	for i := 0; i < n; i++ {
  // Membuat map untuk pencarian O(1): key → value
		slopes := make(map[[2]int]int)
		duplicate := 1

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 && dy == 0 {
				duplicate++
				continue
			}

			g := gcd(dx, dy)
			dx /= g
			dy /= g

			// Normalize slope sign: ensure dx is positive, or if dx==0, dy positive
			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}

			key := [2]int{dx, dy}
			slopes[key]++
		}

		localMax := 0
		for _, count := range slopes {
			if count > localMax {
				localMax = count
			}
		}

		if localMax+duplicate > maxCount {
			maxCount = localMax + duplicate
		}
	}

	return maxCount
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	result := maxPoints(points)
	expected := 3

	fmt.Printf("maxPoints(%v) = %d\n", points, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
