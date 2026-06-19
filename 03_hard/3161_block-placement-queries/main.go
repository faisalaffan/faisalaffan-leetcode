package main

// LeetCode #3161: Block Placement Queries
// https://leetcode.com/problems/block-placement-queries/
// Difficulty: Hard
//
// Queries of two types:
//   1 x -> answer whether a block of length x can be placed (any gap >= x)
//   2 pos -> place an obstacle at coordinate pos
//
// Approach: maintain a sorted set of obstacle positions and a max-heap of gaps.

import (
	"container/heap"
	"fmt"
	"sort"
)

// max-heap of integers.
type maxIntHeap []int

func (h maxIntHeap) Len() int           { return len(h) }
func (h maxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxIntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func blockPlacementQueries(queries [][]int) []bool {
	obstacles := []int{} // sorted positions
	gapHeap := &maxIntHeap{}
	gapCount := make(map[int]int)

	// Insert initial gap from 0 to a large coordinate (1e9).
	const maxCoord = 1_000_000_000
	initialGap := maxCoord
	heap.Push(gapHeap, initialGap)
	gapCount[initialGap] = 1

	addGap := func(sz int) {
		if sz <= 0 {
			return
		}
		gapCount[sz]++
		heap.Push(gapHeap, sz)
	}

	removeGap := func(sz int) {
		if sz <= 0 {
			return
		}
		gapCount[sz]--
	}

	peekMaxGap := func() int {
		for gapHeap.Len() > 0 {
			top := (*gapHeap)[0]
			if cnt := gapCount[top]; cnt > 0 {
				return top
			}
			heap.Pop(gapHeap)
		}
		return 0
	}

	ans := make([]bool, 0)
	for _, q := range queries {
		if q[0] == 1 {
			x := q[1]
			ans = append(ans, peekMaxGap() >= x)
		} else {
			pos := q[1]
			idx := sort.SearchInts(obstacles, pos)
			if idx < len(obstacles) && obstacles[idx] == pos {
				// Already an obstacle here.
				continue
			}

			// Neighbors before insertion.
			left := 0
			if idx > 0 {
				left = obstacles[idx-1]
			}
			right := maxCoord
			if idx < len(obstacles) {
				right = obstacles[idx]
			}

			// Remove old gap between left and right.
			removeGap(right - left)

			// Add new gaps (left-pos) and (pos-right).
			addGap(pos - left)
			addGap(right - pos)

			// Insert obstacle.
			obstacles = append(obstacles, 0)
			copy(obstacles[idx+1:], obstacles[idx:])
			obstacles[idx] = pos
		}
	}
	return ans
}

func main() {
	queries := [][]int{{1, 3}, {2, 2}, {1, 3}, {2, 5}, {1, 3}}
	fmt.Println(blockPlacementQueries(queries))
}
