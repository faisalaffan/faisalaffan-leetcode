package main

// LeetCode #703: Kth Largest Element in a Stream
// https://leetcode.com/problems/kth-largest-element-in-a-stream/
// Difficulty: Easy

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// KthLargest maintains the kth largest element in a stream.
type KthLargest struct {
	k    int
	heap *MinHeap
}

// Constructor creates a KthLargest instance.
func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kl := KthLargest{k: k, heap: h}
	for _, v := range nums {
		kl.Add(v)
	}
	return kl
}

// Add adds a new value and returns the kth largest.
// Time: O(log k). Space: O(k).
func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.heap, val)
	if kl.heap.Len() > kl.k {
		heap.Pop(kl.heap)
	}
	return (*kl.heap)[0]
}

func main() {
	kl := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kl.Add(3))  // 4
	fmt.Println(kl.Add(5))  // 5
	fmt.Println(kl.Add(10)) // 5
	fmt.Println(kl.Add(9))  // 8
	fmt.Println(kl.Add(4))  // 8
}
