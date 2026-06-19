package main

// LeetCode #857: Minimum Cost to Hire K Workers
// https://leetcode.com/problems/minimum-cost-to-hire-k-workers/
// Difficulty: Hard
// Approach: Sort workers by wage/quality ratio. For any group hired at a given ratio,
// each worker gets paid at least their wage expectation. Use a max-heap to track the
// K smallest qualities among workers with ratio <= current ratio.

import (
	"container/heap"
	"fmt"
	"sort"
)

type worker struct {
	quality int
	ratio   float64
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] } // max-heap
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	workers := make([]worker, n)
	for i := 0; i < n; i++ {
		workers[i] = worker{quality: quality[i], ratio: float64(wage[i]) / float64(quality[i])}
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].ratio < workers[j].ratio
	})

	h := &maxHeap{}
	heap.Init(h)
	sumQ := 0

	for i := 0; i < k; i++ {
		sumQ += workers[i].quality
		heap.Push(h, workers[i].quality)
	}

	ans := workers[k-1].ratio * float64(sumQ)

	for i := k; i < n; i++ {
		// Remove worker with largest quality, add current
		largest := heap.Pop(h).(int)
		sumQ -= largest

		sumQ += workers[i].quality
		heap.Push(h, workers[i].quality)

		cost := workers[i].ratio * float64(sumQ)
		if cost < ans {
			ans = cost
		}
	}

	return ans
}

func main() {
	fmt.Printf("%.1f\n", mincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2))
	// Expected: 105.0

	fmt.Printf("%.1f\n", mincostToHireWorkers([]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3))
	// Additional test
}
