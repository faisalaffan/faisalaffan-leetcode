# 1610 — Maximum Number Of Visible Points

## Deskripsi

**Soal:** [1610. Maximum Number Of Visible Points](https://leetcode.com/problems/maximum-number-of-visible-points/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser)

**Fungsi Solusi:** `func visiblePoints(points [][]int, angle int, location []int) int`

## Solusi Go

```go
package main

// LeetCode #1610: Maximum Number of Visible Points
// https://leetcode.com/problems/maximum-number-of-visible-points/
// Difficulty: Hard

import (
	"fmt"
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	angles := []float64{}
	same := 0

	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			same++
			continue
		}
		// atan2 returns angle in [-pi, pi], convert to degrees
		rad := math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0]))
		deg := rad * 180.0 / math.Pi
		angles = append(angles, deg)
	}

	sort.Float64s(angles)
	n := len(angles)

	// Duplicate with +360 for circular sliding window
	for i := 0; i < n; i++ {
		angles = append(angles, angles[i]+360.0)
	}

	maxVis := 0
	j := 0
	angleF := float64(angle)

	for i := 0; i < n; i++ {
		for j < len(angles) && angles[j] <= angles[i]+angleF {
			j++
		}
		if j-i > maxVis {
			maxVis = j - i
		}
	}

	return same + maxVis
}

func main() {
	// Test case 1: points=[[2,1],[2,2],[3,4]], angle=90, location=[1,1] -> 3
	points := [][]int{{2, 1}, {2, 2}, {3, 4}}
	result := visiblePoints(points, 90, []int{1, 1})
	fmt.Printf("points=%v angle=90 location=[1,1] -> %d (expected 3)\n", points, result)

	// Test case 2: points=[[1,1],[2,2],[3,3],[1,1]], angle=0, location=[1,1] -> 4
	// Same location points: 2 (both [1,1]).
	// Points [2,2] and [3,3] share the same angle (45 deg), both visible with angle=0.
	points2 := [][]int{{1, 1}, {2, 2}, {3, 3}, {1, 1}}
	result2 := visiblePoints(points2, 0, []int{1, 1})
	fmt.Printf("points=%v angle=0 location=[1,1] -> %d\n", points2, result2)

	// Test case 3: points=[[0,0]], angle=90, location=[1,1] -> 0
	points3 := [][]int{{0, 0}}
	result3 := visiblePoints(points3, 90, []int{1, 1})
	fmt.Printf("points=%v angle=90 location=[1,1] -> %d\n", points3, result3)
}
```
