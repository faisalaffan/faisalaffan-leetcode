package main

// LeetCode #2233: Maximum Product After K Increments
// https://leetcode.com/problems/maximum-product-after-k-increments/
// Difficulty: Medium
// Time: O(n + k log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	const mod = 1_000_000_007
	h := &minHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	for i := 0; i < k; i++ {
		v := heap.Pop(h).(int)
		heap.Push(h, v+1)
	}
	prod := 1
	for h.Len() > 0 {
		prod = (prod * heap.Pop(h).(int)) % mod
	}
	return prod
}

func main() {
	// Test case 1
	fmt.Println(maximumProduct([]int{0, 4}, 5))
	// Expected: 20

	// Test case 2
	fmt.Println(maximumProduct([]int{6, 3, 3, 2}, 2))
	// Expected: 216
}
