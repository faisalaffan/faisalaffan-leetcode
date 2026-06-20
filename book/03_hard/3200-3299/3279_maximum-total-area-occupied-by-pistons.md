# 3279 — Maximum Total Area Occupied By Pistons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxTotalArea(startTime, endTime []int, yRanges [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3279: Maximum Total Area Occupied by Pistons
// https://leetcode.com/problems/maximum-total-area-occupied-by-pistons/
// Difficulty: Hard [Paid]
//
// Each piston moves vertically between a minimum and maximum y-position
// over time. Given start times, end times, and y-ranges for each piston,
// compute the maximum total area (over time) occupied collectively by
// the pistons.
//
// At any time t, the union of vertical intervals occupied by active pistons
// forms a set of segments. The total area = integral over time of the
// total length of covered y-range.
//
// This can be solved by scanning events: for each time where the state
// of piston activity changes, compute the union of active y-intervals
// and multiply by the elapsed time since the last event.
//
// Equivalent to: given N rectangles [time_start, time_end] x [y_min, y_max],
// compute the area of their union.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: two non-overlapping pistons
	// Piston A: t=[0,5], y=[0,2]; Piston B: t=[3,8], y=[1,3]
	fmt.Println(maxTotalArea([]int{0, 3}, []int{5, 8}, [][]int{{0, 2}, {1, 3}}))
	// Example 2: single piston
	fmt.Println(maxTotalArea([]int{0}, []int{10}, [][]int{{0, 5}}))
	// Example 3: overlapping time, disjoint y
	fmt.Println(maxTotalArea([]int{0, 0}, []int{10, 10}, [][]int{{0, 2}, {3, 5}}))
	// Example 4: no pistons
	fmt.Println(maxTotalArea([]int{}, []int{}, [][]int{}))
	// Example 5: three pistons
	fmt.Println(maxTotalArea([]int{0, 2, 4}, []int{6, 8, 10}, [][]int{{0, 3}, {1, 4}, {2, 5}}))
}

func maxTotalArea(startTime, endTime []int, yRanges [][]int) int64 {
	n := len(startTime)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Collect all unique time events.
	type event struct {
		t    int
		idx  int
		add bool // true = start, false = end
	}
	events := make([]event, 0, 2*n)
	for i := 0; i < n; i++ {
		events = append(events, event{startTime[i], i, true})
		events = append(events, event{endTime[i], i, false})
	}
  // Custom sort
	sort.Slice(events, func(i, j int) bool {
		if events[i].t != events[j].t {
			return events[i].t < events[j].t
		}
		return events[i].add && !events[j].add // start before end at same time
	})

	// Active intervals: sweeping over y-axis.
	// Maintain count of active intervals covering each y-position.
	// Since y values are integers, we use a map for the difference array.
  // HashMap: O(1) lookup
	active := make(map[int]int) // diff[y] = net change in active intervals at position y

	addInterval := func(y1, y2 int) {
		active[y1]++
		active[y2+1]--
	}
	removeInterval := func(y1, y2 int) {
		active[y1]--
		active[y2+1]++
	}

	computeUnionLength := func() int64 {
		if len(active) == 0 {
			return 0
		}
		// Sort the y-boundary positions.
  // Alokasi slice
		ys := make([]int, 0, len(active))
		for y := range active {
			ys = append(ys, y)
		}
  // Sort O(n log n)
		sort.Ints(ys)

		var length int64
		var count int
  // Linear scan O(n)
		for i := 0; i < len(ys)-1; i++ {
			count += active[ys[i]]
			if count > 0 {
				length += int64(ys[i+1] - ys[i])
			}
		}
		return length
	}

	var totalArea int64
	prevTime := events[0].t

	for _, e := range events {
		elapsed := int64(e.t - prevTime)
		if elapsed > 0 {
			totalArea += elapsed * computeUnionLength()
		}

		if e.add {
			addInterval(yRanges[e.idx][0], yRanges[e.idx][1])
		} else {
			removeInterval(yRanges[e.idx][0], yRanges[e.idx][1])
		}
		prevTime = e.t
	}

	return totalArea
}
```
