package main

// LeetCode #502: IPO
// https://leetcode.com/problems/ipo/
// Difficulty: Hard
// Approach: Two heaps. Sort projects by capital, use a max-heap for profits
// of affordable projects. Iterate k times, each time add all projects whose
// capital <= current w to the max-heap, then pick the most profitable one.

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("502 - IPO")

	// Test cases
	k1 := 2
	w1 := 0
	profits1 := []int{1, 2, 3}
	capital1 := []int{0, 1, 1}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 4)\n",
		k1, w1, profits1, capital1, findMaximizedCapital(k1, w1, profits1, capital1))

	k2 := 3
	w2 := 0
	profits2 := []int{1, 2, 3}
	capital2 := []int{0, 1, 2}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 6)\n",
		k2, w2, profits2, capital2, findMaximizedCapital(k2, w2, profits2, capital2))

	k3 := 1
	w3 := 2
	profits3 := []int{1, 2, 3}
	capital3 := []int{1, 1, 2}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 5)\n",
		k3, w3, profits3, capital3, findMaximizedCapital(k3, w3, profits3, capital3))

	k4 := 0
	w4 := 0
	profits4 := []int{1, 2, 3}
	capital4 := []int{0, 1, 1}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 0)\n",
		k4, w4, profits4, capital4, findMaximizedCapital(k4, w4, profits4, capital4))

	// With w=0, no project is affordable since min capital=1 > 0
	k5 := 10
	w5 := 0
	profits5 := []int{1, 2, 3, 4, 5}
	capital5 := []int{5, 4, 3, 2, 1}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 0, none affordable)\n",
		k5, w5, profits5, capital5, findMaximizedCapital(k5, w5, profits5, capital5))

	// LeetCode example
	k6 := 2
	w6 := 0
	profits6 := []int{1, 2, 3}
	capital6 := []int{0, 9, 10}
	fmt.Printf("k=%d, w=%d, profits=%v, capital=%v -> %d (expected: 1)\n",
		k6, w6, profits6, capital6, findMaximizedCapital(k6, w6, profits6, capital6))
}

// Max-heap for profits
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)

	// Create project pairs (capital, profit) and sort by capital
	projects := make([][2]int, n)
	for i := 0; i < n; i++ {
		projects[i] = [2]int{capital[i], profits[i]}
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i][0] < projects[j][0]
	})

	// Max-heap for profits of available projects
	profitHeap := &MaxHeap{}
	heap.Init(profitHeap)

	idx := 0
	for i := 0; i < k; i++ {
		// Add all projects we can afford now
		for idx < n && projects[idx][0] <= w {
			heap.Push(profitHeap, projects[idx][1])
			idx++
		}

		// If no projects are available, we're done
		if profitHeap.Len() == 0 {
			break
		}

		// Pick the most profitable project
		w += heap.Pop(profitHeap).(int)
	}

	return w
}
