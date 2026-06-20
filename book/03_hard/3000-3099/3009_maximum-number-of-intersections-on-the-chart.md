# 3009 — Maximum Number Of Intersections On The Chart

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxIntersectionCount(y []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3009: Maximum Number of Intersections on the Chart
// https://leetcode.com/problems/maximum-number-of-intersections-on-the-chart/
// Difficulty: Hard [Paid]
//
// Given an array y representing y-coordinates at integer x positions,
// consider the polyline connecting (i, y[i]) for i=0..n-1.
// A vertical line at x = c (c may be any real number) intersects the
// polyline at some points. Find the maximum number of such intersections
// possible.
//
// Approach: Sweep line
//   Each line segment between (i-1, y[i-1]) and (i, y[i]) occupies a range
//   of y-values. We track overlap counts using events (sweep line).
//   Use 2*y coordinate space to handle half-integer intersection points.

import (
	"fmt"
	"sort"
)

func maxIntersectionCount(y []int) int {
	n := len(y)
	type event struct {
		x     int
		delta int
	}
	events := make([]event, 0, 2*n)

	for i := 1; i < n; i++ {
		s := 2 * y[i-1]
		e := 2 * y[i]

		// For intermediate vertices, adjust endpoint to avoid double-counting
		// at the vertex itself (unless it's the last vertex)
		if i != n-1 {
			if y[i-1] < y[i] {
				e--
			} else {
				e++
			}
		}

		if s > e {
			s, e = e, s
		}
		events = append(events, event{s, 1}, event{e + 1, -1})
	}

  // Custom sort
	sort.Slice(events, func(i, j int) bool {
		if events[i].x != events[j].x {
			return events[i].x < events[j].x
		}
		return events[i].delta < events[j].delta
	})

	ans := 0
	cur := 0
	for _, ev := range events {
		cur += ev.delta
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

func main() {
	// Test 1
	fmt.Println("Test 1:", maxIntersectionCount([]int{1, 2, 1, 2, 1, 3, 2}))

	// Test 2
	fmt.Println("Test 2:", maxIntersectionCount([]int{2, 1, 3, 4, 5}))

	// Flat line
	fmt.Println("Test 3:", maxIntersectionCount([]int{5, 5, 5, 5}))

	// Strictly increasing
	fmt.Println("Test 4:", maxIntersectionCount([]int{1, 2, 3, 4, 5}))

	// Up-down-up
	fmt.Println("Test 5:", maxIntersectionCount([]int{1, 5, 2, 6, 3}))

	// Two elements
	fmt.Println("Test 6:", maxIntersectionCount([]int{1, 3}))
}
```
