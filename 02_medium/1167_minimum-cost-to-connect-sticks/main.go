package main

import (
	"container/heap"
	"fmt"
)

// LeetCode #1167: Minimum Cost to Connect Sticks
// https://leetcode.com/problems/minimum-cost-to-connect-sticks/
// Difficulty: Medium [Paid]

// Always combine the two smallest sticks (greedy + min-heap).
// Cost of each combination = sum of the two sticks.

// Time: O(n log n)
// Space: O(n)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
	if len(sticks) <= 1 {
		return 0
	}

	h := &MinHeap{}
	heap.Init(h)
	for _, s := range sticks {
		heap.Push(h, s)
	}

	total := 0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		cost := a + b
		total += cost
		heap.Push(h, cost)
	}

	return total
}

func main() {
	fmt.Printf("%d (expected: 14)\n", connectSticks([]int{2, 4, 3}))
	fmt.Printf("%d (expected: 30)\n", connectSticks([]int{1, 8, 3, 5}))
	fmt.Printf("%d (expected: 0)\n", connectSticks([]int{5}))
}
