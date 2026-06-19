package main

// LeetCode #2530: Maximal Score After Applying K Operations
// https://leetcode.com/problems/maximal-score-after-applying-k-operations/
// Difficulty: Medium
// Time: O((n + k) log n) | Space: O(n)
// Max heap: each operation take max, add ceil(v/3), put back.

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x
}

func main() {
	fmt.Println(maxKelements([]int{10, 10, 10, 10, 10}, 5)) // 50
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))     // 17
}

func maxKelements(nums []int, k int) int64 {
	h := &MaxHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	var score int64
	for i := 0; i < k; i++ {
		v := heap.Pop(h).(int)
		score += int64(v)
		heap.Push(h, (v+2)/3) // ceil(v/3)
	}
	return score
}
