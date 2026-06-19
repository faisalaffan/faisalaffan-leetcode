package main

import (
	"fmt"
	"math"
)

// LeetCode #1182: Shortest Distance to Target Color
// https://leetcode.com/problems/shortest-distance-to-target-color/
// Difficulty: Medium [Paid]

// Precompute nearest distance to each color from both directions.

// Time: O(n + m) where m = len(queries)
// Space: O(n)

func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(colors)

	// left[i][c] = nearest distance to color c from left side up to i
	left := make([][3]int, n)
	for i := 0; i < n; i++ {
		for c := 0; c < 3; c++ {
			left[i][c] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			left[i] = left[i-1]
		}
		left[i][colors[i]-1] = 0
		// Update distances: all non-zero entries get +1
		for c := 0; c < 3; c++ {
			if left[i][c] != 0 && left[i][c] != math.MaxInt32 {
				// This gets complicated. Let's take a simpler approach.
			}
		}
	}

	// Simpler approach: for each color, store sorted positions
	positions := make([][]int, 4) // 1-indexed colors
	for i, c := range colors {
		positions[c] = append(positions[c], i)
	}

	binarySearch := func(pos []int, target int) int {
		lo, hi := 0, len(pos)-1
		if target <= pos[lo] {
			return pos[lo]
		}
		if target >= pos[hi] {
			return pos[hi]
		}
		for lo <= hi {
			mid := (lo + hi) / 2
			if pos[mid] == target {
				return target
			}
			if pos[mid] < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		// lo is insertion point
		dist := abs(target - pos[lo])
		if lo > 0 && abs(target-pos[lo-1]) < dist {
			dist = abs(target - pos[lo-1])
			return pos[lo-1]
		}
		return pos[lo]
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		idx, color := q[0], q[1]
		if len(positions[color]) == 0 {
			result[i] = -1
		} else {
			nearest := binarySearch(positions[color], idx)
			result[i] = abs(idx - nearest)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("%v (expected: [3 0 3])\n",
		shortestDistanceColor([]int{1, 1, 2, 1, 3, 2, 2, 3, 3},
			[][]int{{1, 3}, {2, 2}, {6, 1}}))

	fmt.Printf("%v (expected: [-1])\n",
		shortestDistanceColor([]int{1, 2},
			[][]int{{0, 3}}))
}
