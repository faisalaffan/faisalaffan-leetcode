package main

// LeetCode #497: Random Point in Non-overlapping Rectangles
// https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/
// Difficulty: Medium
// Time: O(n) for init, O(log n) per pick
// Space: O(n)

import (
	"fmt"
	"math/rand"
)

func main() {
	sol := Constructor([][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}})
	for i := 0; i < 3; i++ {
		p := sol.Pick()
		fmt.Println(p)
	}
}

type Solution struct {
	rects      [][]int
	prefixSum  []int
	totalPts   int
}

func Constructor(rects [][]int) Solution {
	prefixSum := make([]int, len(rects))
	total := 0
	for i, r := range rects {
		pts := (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		total += pts
		prefixSum[i] = total
	}
	return Solution{rects: rects, prefixSum: prefixSum, totalPts: total}
}

func (s *Solution) Pick() []int {
	// Pick a random point index
	r := rand.Intn(s.totalPts) + 1
	// Binary search to find which rectangle
	idx := search(s.prefixSum, r)

	rect := s.rects[idx]
	prev := 0
	if idx > 0 {
		prev = s.prefixSum[idx-1]
	}
	offset := r - prev - 1
	width := rect[2] - rect[0] + 1
	x := rect[0] + offset%width
	y := rect[1] + offset/width
	return []int{x, y}
}

func search(prefixSum []int, target int) int {
	lo, hi := 0, len(prefixSum)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if prefixSum[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
