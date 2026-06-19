package main

// LeetCode #1172: Dinner Plate Stacks
// https://leetcode.com/problems/dinner-plate-stacks/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
)

// MinHeap of available (non-full) stack indices
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type DinnerPlates struct {
	cap      int
	stacks   [][]int
	avail    MinHeap // min-heap of non-full stack indices
	nonEmpty int     // rightmost non-empty stack index (for pop)
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		cap:      capacity,
		stacks:   make([][]int, 0),
		avail:    MinHeap{},
		nonEmpty: -1,
	}
}

func (dp *DinnerPlates) Push(val int) {
	// Clean up stale available indices
	for dp.avail.Len() > 0 && dp.avail[0] < len(dp.stacks) && len(dp.stacks[dp.avail[0]]) == dp.cap {
		heap.Pop(&dp.avail)
	}

	if dp.avail.Len() > 0 {
		idx := dp.avail[0]
		dp.stacks[idx] = append(dp.stacks[idx], val)
		if len(dp.stacks[idx]) == dp.cap {
			heap.Pop(&dp.avail)
		}
		if idx > dp.nonEmpty {
			dp.nonEmpty = idx
		}
		return
	}

	// No available stack, create new one
	dp.stacks = append(dp.stacks, []int{val})
	dp.nonEmpty = len(dp.stacks) - 1
	if dp.cap > 1 {
		heap.Push(&dp.avail, len(dp.stacks)-1)
	}
}

func (dp *DinnerPlates) Pop() int {
	if dp.nonEmpty < 0 {
		return -1
	}
	idx := dp.nonEmpty
	val := dp.stacks[idx][len(dp.stacks[idx])-1]
	dp.stacks[idx] = dp.stacks[idx][:len(dp.stacks[idx])-1]

	// If stack becomes non-full, add to avail
	if len(dp.stacks[idx]) == dp.cap-1 {
		heap.Push(&dp.avail, idx)
	}

	// Update nonEmpty (find rightmost non-empty, skip empty)
	for dp.nonEmpty >= 0 && len(dp.stacks[dp.nonEmpty]) == 0 {
		dp.nonEmpty--
	}
	return val
}

func (dp *DinnerPlates) PopAtStack(index int) int {
	if index >= len(dp.stacks) || len(dp.stacks[index]) == 0 {
		return -1
	}
	val := dp.stacks[index][len(dp.stacks[index])-1]
	dp.stacks[index] = dp.stacks[index][:len(dp.stacks[index])-1]

	// If it became non-full and index < nonEmpty, add to avail
	if index < dp.nonEmpty && len(dp.stacks[index]) == dp.cap-1 {
		heap.Push(&dp.avail, index)
	}

	// Update nonEmpty if needed
	if len(dp.stacks[index]) == 0 && index == dp.nonEmpty {
		for dp.nonEmpty >= 0 && len(dp.stacks[dp.nonEmpty]) == 0 {
			dp.nonEmpty--
		}
	}
	return val
}

func main() {
	// Test case
	dp := Constructor(2)
	dp.Push(1)
	dp.Push(2)
	dp.Push(3)
	dp.Push(4)
	dp.Push(5)
	fmt.Println(dp.PopAtStack(0)) // 2
	dp.Push(6)
	dp.Push(7)
	fmt.Println(dp.Pop())  // 7
	fmt.Println(dp.Pop())  // 6
	fmt.Println(dp.Pop())  // 5
	fmt.Println(dp.PopAtStack(0)) // 1
	fmt.Println(dp.Pop())  // 4
	fmt.Println(dp.Pop())  // 3
	fmt.Println(dp.Pop())  // -1

	fmt.Println("---")

	// Test case 2: capacity 1
	dp2 := Constructor(1)
	dp2.Push(10)
	dp2.Push(20)
	fmt.Println(dp2.PopAtStack(0)) // 10
	fmt.Println(dp2.Pop())  // 20
	fmt.Println(dp2.Pop())  // -1
}
