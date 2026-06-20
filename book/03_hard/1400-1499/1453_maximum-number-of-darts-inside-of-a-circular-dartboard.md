# 1453 — Maximum Number Of Darts Inside Of A Circular Dartboard

## Deskripsi

**Soal:** [1453. Maximum Number Of Darts Inside Of A Circular Dartboard](https://leetcode.com/problems/maximum-number-of-darts-inside-of-a-circular-dartboard/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser)

> **Ide Kunci:** Angular sweep. For each point as center candidate, compute

## Solusi Go

```go
package main

// LeetCode #1453: Maximum Number of Darts Inside of a Circular Dartboard
// https://leetcode.com/problems/maximum-number-of-darts-inside-of-a-circular-dartboard/
// Difficulty: Hard
//
// Given points on a 2D plane and a radius r, find the maximum number
// of points that can be covered by a circle of radius r.
//
// Approach: Angular sweep. For each point as center candidate, compute
// angles of all other points that lie within 2r distance, then use
// sliding window to find max points in any semicircle.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(numPoints([][]int{{-2, 0}, {2, 0}, {0, 2}, {0, -2}}, 2))
	// Example 2
	fmt.Println(numPoints([][]int{{-3, 0}, {3, 0}, {2, 6}, {5, 4}, {0, 9}, {7, 8}}, 5))
	// Edge: single point
	fmt.Println(numPoints([][]int{{0, 0}}, 1))
}

func numPoints(darts [][]int, r int) int {
	n := len(darts)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	rr := float64(r) * float64(r)
	ans := 1

	for i := 0; i < n; i++ {
  // Membuat slice untuk menyimpan hasil
		angles := make([]float64, 0, n*2)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dx := float64(darts[j][0] - darts[i][0])
			dy := float64(darts[j][1] - darts[i][1])
			d2 := dx*dx + dy*dy
			if d2 > 4*rr {
				continue
			}
			if d2 == 0 {
				continue
			}
			dist := math.Sqrt(d2)
			half := dist / 2.0
			// Angle from i to j
			base := math.Atan2(dy, dx)
			// Maximum deviation from base
			delta := math.Acos(half / float64(r))
			// Ensure delta is valid
			if !math.IsNaN(delta) {
				angles = append(angles, base-delta, base+delta+2*math.Pi)
			}
		}
		if len(angles) == 0 {
			continue
		}
		sort.Float64s(angles)

		m := len(angles)
		cnt := 0
		left := 0
		for right := 0; right < m; right++ {
			if angles[right]-angles[left] > 2*math.Pi+1e-9 {
				if angles[right]-angles[left] > 2*math.Pi+1e-9 {
					left++
				}
			}
			cnt = right - left + 1
			if cnt+1 > ans {
				ans = cnt + 1
			}
		}
	}
	return ans
}
```
