package main

// LeetCode #1642: Furthest Building You Can Reach
// https://leetcode.com/problems/furthest-building-you-can-reach/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(FurthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Println(FurthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(FurthestBuilding([]int{14, 3, 19, 3}, 17, 0))
}

// MinHeap for int
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	// Time: O(N log K), Space: O(K) where K = ladders
	// Use a min-heap to store the largest climbs where ladders are used
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff <= 0 {
			continue
		}

		heap.Push(h, diff)

		// If we have more climbs than ladders, use bricks for the smallest climb
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
			if bricks < 0 {
				return i
			}
		}
	}

	return len(heights) - 1
}
