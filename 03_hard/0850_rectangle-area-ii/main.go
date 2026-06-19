package main

// LeetCode #850: Rectangle Area II
// https://leetcode.com/problems/rectangle-area-ii/
// Difficulty: Hard
// Approach: Sweep line + coordinate compression. Process vertical events (x, y1, y2, type),
// sorted by x. Maintain active y-intervals; at each step, compute covered y-length and
// multiply by delta-x.

import (
	"fmt"
	"sort"
)

const mod = 1_000_000_007

func rectangleArea(rectangles [][]int) int {
	type event struct {
		x    int
		y1   int
		y2   int
		typ  int // +1 = entering, -1 = leaving
	}

	events := make([]event, 0, len(rectangles)*2)
	for _, r := range rectangles {
		events = append(events, event{r[0], r[1], r[2], 1})
		events = append(events, event{r[3], r[1], r[2], -1})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].x < events[j].x
	})

	// Coordinate compress y values
	ys := make(map[int]bool)
	for _, r := range rectangles {
		ys[r[1]] = true
		ys[r[2]] = true
	}
	ySorted := make([]int, 0, len(ys))
	for y := range ys {
		ySorted = append(ySorted, y)
	}
	sort.Ints(ySorted)
	yIndex := make(map[int]int)
	for i, y := range ySorted {
		yIndex[y] = i
	}

	// Segment for y intervals
	m := len(ySorted)
	count := make([]int, m-1) // count[i] = active rectangles covering interval [ySorted[i], ySorted[i+1])

	area := 0
	prevX := events[0].x

	for _, e := range events {
		// Add area from prevX to e.x
		width := e.x - prevX
		if width > 0 {
			coverY := 0
			for i := 0; i < m-1; i++ {
				if count[i] > 0 {
					coverY += ySorted[i+1] - ySorted[i]
				}
			}
			area = (area + width*coverY%mod) % mod
		}

		// Update count for this event's y interval
		l := yIndex[e.y1]
		r := yIndex[e.y2]
		for i := l; i < r; i++ {
			count[i] += e.typ
		}
		prevX = e.x
	}

	return area
}

func main() {
	fmt.Println(rectangleArea([][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}})) // Expected: 6
	fmt.Println(rectangleArea([][]int{{0, 0, 1, 1}, {2, 2, 3, 3}}))               // Expected: 2
}
